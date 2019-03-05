package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

type forward struct {
	hubAddr          *net.UDPAddr

	// TODO FIXME This should be per connection!
	connectionId     string
	localTcpListener net.Listener
	localUdpConn     net.PacketConn
	peerUdpAddr      *net.UDPAddr

	// TODO FIXME This should be per TCP connection!
	dataChan         chan DataMessage
}

func (f *forward) start(hubAddr string, source string, sourcePort int, target string, targetForwardAddr string) {
	var err error

	// Resolve hub address
	f.hubAddr, err = net.ResolveUDPAddr("udp4", hubAddr)
	if err != nil {
		panic(err)
	}

	// TODO This only supports one forward and one connection!!!

	f.dataChan = make(chan DataMessage)

	// Listen to local UDP address
	rand.Seed(time.Now().Unix())
	localUdpPort := fmt.Sprintf(":%d", 10000+rand.Intn(10000))
	f.localUdpConn, err = net.ListenPacket("udp", localUdpPort)
	if err != nil {
		panic(err)
	}

	// Listen to local TCP address
	f.localTcpListener, err = net.Listen("tcp", fmt.Sprintf(":%d", sourcePort))
	if err != nil {
		panic(err)
	}

	go f.listenUdp()
	go f.listenTcp()
	go f.sendRequest(source, sourcePort, target, targetForwardAddr)

	for {
		time.Sleep(30 * time.Second)
	}
}

func (f *forward) sendRequest(source string, sourcePort int, target string, targetForwardAddr string) {
	log.Printf("Requesting connection to %s:%d\n", target, targetForwardAddr)

	request := &ForwardRequest{
		Source: source,
		Target: target,
		TargetForwardAddr: targetForwardAddr,
	}

	sendmsg(f.localUdpConn, f.hubAddr, messageTypeForwardRequest, request)

	time.Sleep(5 * time.Second)
}

func (f *forward) listenUdp() {
	for {
		addr, messageType, message := recvmsg(f.localUdpConn)

		switch messageType {
		case messageTypeKeepaliveResponse:
			response, _ := message.(*KeepaliveResponse)
			log.Println("> keepalive", response.Id)
		case messageTypeKeepaliveRequest:
			request, _ := message.(*KeepaliveRequest)
			sendmsg(f.localUdpConn, addr, messageTypeKeepaliveResponse, &KeepaliveResponse{
				Id: request.Id,
				Rand: request.Rand,
			})
		case messageTypeDataMessage:
			msg, _ := message.(*DataMessage)
			f.dataChan <- *msg
		case messageTypeForwardResponse:
			response, _ := message.(*ForwardResponse)

			if response.Success {
				var err error
				log.Print("Peer address: ", response.TargetAddr)

				f.peerUdpAddr, err = net.ResolveUDPAddr("udp4", response.TargetAddr)
				if err != nil {
					panic(err)
				}

				f.connectionId = response.Id

				go f.keepalive()
			} else {
				log.Println("Failed forward response")
			}

		}
	}
}

func (f *forward) listenTcp() {
	for {
		conn, err := f.localTcpListener.Accept()
		if err != nil {
			panic(err)
		}

		go f.handleTcpOutgoing(conn)
		go f.handleTcpIncoming(conn)
	}
}

func (f *forward) handleTcpOutgoing(conn net.Conn) {
	for f.peerUdpAddr == nil { // TODO racy
		log.Println("Cannot forward yet. UDP connection not active yet.")
		time.Sleep(1 * time.Second)
	}

	buf := make([]byte, messageSendBufferBytes)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			break
		}

		sendmsg(f.localUdpConn, f.peerUdpAddr, messageTypeDataMessage, &DataMessage{
			Id: f.connectionId,
			Data: buf[:n],
		})
	}

	conn.Close()
}

func (f *forward) handleTcpIncoming(conn net.Conn) {
	for {
		msg := <-f.dataChan
		conn.Write(msg.Data)
	}
}

func (f *forward) keepalive() {
	for {
		sendmsg(f.localUdpConn, f.peerUdpAddr, messageTypeKeepaliveRequest, &KeepaliveRequest{
			Id: f.connectionId,
			Rand: rand.Int31(),
		})

		time.Sleep(15 * time.Second)
	}
}