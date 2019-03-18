// Code generated by protoc-gen-go. DO NOT EDIT.
// source: natter.proto

package internal

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// 0x01
type CheckinRequest struct {
	Source               string   `protobuf:"bytes,1,opt,name=Source,proto3" json:"Source,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckinRequest) Reset()         { *m = CheckinRequest{} }
func (m *CheckinRequest) String() string { return proto.CompactTextString(m) }
func (*CheckinRequest) ProtoMessage()    {}
func (*CheckinRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_39129eec69c18084, []int{0}
}

func (m *CheckinRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckinRequest.Unmarshal(m, b)
}
func (m *CheckinRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckinRequest.Marshal(b, m, deterministic)
}
func (m *CheckinRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckinRequest.Merge(m, src)
}
func (m *CheckinRequest) XXX_Size() int {
	return xxx_messageInfo_CheckinRequest.Size(m)
}
func (m *CheckinRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckinRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckinRequest proto.InternalMessageInfo

func (m *CheckinRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

// 0x02
type CheckinResponse struct {
	Addr                 string   `protobuf:"bytes,1,opt,name=Addr,proto3" json:"Addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckinResponse) Reset()         { *m = CheckinResponse{} }
func (m *CheckinResponse) String() string { return proto.CompactTextString(m) }
func (*CheckinResponse) ProtoMessage()    {}
func (*CheckinResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_39129eec69c18084, []int{1}
}

func (m *CheckinResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckinResponse.Unmarshal(m, b)
}
func (m *CheckinResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckinResponse.Marshal(b, m, deterministic)
}
func (m *CheckinResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckinResponse.Merge(m, src)
}
func (m *CheckinResponse) XXX_Size() int {
	return xxx_messageInfo_CheckinResponse.Size(m)
}
func (m *CheckinResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckinResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckinResponse proto.InternalMessageInfo

func (m *CheckinResponse) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

// 0x03
type ForwardRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Source               string   `protobuf:"bytes,2,opt,name=Source,proto3" json:"Source,omitempty"`
	SourceAddr           string   `protobuf:"bytes,3,opt,name=SourceAddr,proto3" json:"SourceAddr,omitempty"`
	Target               string   `protobuf:"bytes,4,opt,name=Target,proto3" json:"Target,omitempty"`
	TargetAddr           string   `protobuf:"bytes,5,opt,name=TargetAddr,proto3" json:"TargetAddr,omitempty"`
	TargetForwardAddr    string   `protobuf:"bytes,6,opt,name=TargetForwardAddr,proto3" json:"TargetForwardAddr,omitempty"`
	TargetCommand        []string `protobuf:"bytes,7,rep,name=TargetCommand,proto3" json:"TargetCommand,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ForwardRequest) Reset()         { *m = ForwardRequest{} }
func (m *ForwardRequest) String() string { return proto.CompactTextString(m) }
func (*ForwardRequest) ProtoMessage()    {}
func (*ForwardRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_39129eec69c18084, []int{2}
}

func (m *ForwardRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForwardRequest.Unmarshal(m, b)
}
func (m *ForwardRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForwardRequest.Marshal(b, m, deterministic)
}
func (m *ForwardRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForwardRequest.Merge(m, src)
}
func (m *ForwardRequest) XXX_Size() int {
	return xxx_messageInfo_ForwardRequest.Size(m)
}
func (m *ForwardRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ForwardRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ForwardRequest proto.InternalMessageInfo

func (m *ForwardRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ForwardRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *ForwardRequest) GetSourceAddr() string {
	if m != nil {
		return m.SourceAddr
	}
	return ""
}

func (m *ForwardRequest) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *ForwardRequest) GetTargetAddr() string {
	if m != nil {
		return m.TargetAddr
	}
	return ""
}

func (m *ForwardRequest) GetTargetForwardAddr() string {
	if m != nil {
		return m.TargetForwardAddr
	}
	return ""
}

func (m *ForwardRequest) GetTargetCommand() []string {
	if m != nil {
		return m.TargetCommand
	}
	return nil
}

// 0x04
type ForwardResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Success              bool     `protobuf:"varint,2,opt,name=Success,proto3" json:"Success,omitempty"`
	Source               string   `protobuf:"bytes,3,opt,name=Source,proto3" json:"Source,omitempty"`
	SourceAddr           string   `protobuf:"bytes,4,opt,name=SourceAddr,proto3" json:"SourceAddr,omitempty"`
	Target               string   `protobuf:"bytes,5,opt,name=Target,proto3" json:"Target,omitempty"`
	TargetAddr           string   `protobuf:"bytes,6,opt,name=TargetAddr,proto3" json:"TargetAddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ForwardResponse) Reset()         { *m = ForwardResponse{} }
func (m *ForwardResponse) String() string { return proto.CompactTextString(m) }
func (*ForwardResponse) ProtoMessage()    {}
func (*ForwardResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_39129eec69c18084, []int{3}
}

func (m *ForwardResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForwardResponse.Unmarshal(m, b)
}
func (m *ForwardResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForwardResponse.Marshal(b, m, deterministic)
}
func (m *ForwardResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForwardResponse.Merge(m, src)
}
func (m *ForwardResponse) XXX_Size() int {
	return xxx_messageInfo_ForwardResponse.Size(m)
}
func (m *ForwardResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ForwardResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ForwardResponse proto.InternalMessageInfo

func (m *ForwardResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ForwardResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *ForwardResponse) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *ForwardResponse) GetSourceAddr() string {
	if m != nil {
		return m.SourceAddr
	}
	return ""
}

func (m *ForwardResponse) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *ForwardResponse) GetTargetAddr() string {
	if m != nil {
		return m.TargetAddr
	}
	return ""
}

func init() {
	proto.RegisterType((*CheckinRequest)(nil), "internal.CheckinRequest")
	proto.RegisterType((*CheckinResponse)(nil), "internal.CheckinResponse")
	proto.RegisterType((*ForwardRequest)(nil), "internal.ForwardRequest")
	proto.RegisterType((*ForwardResponse)(nil), "internal.ForwardResponse")
}

func init() { proto.RegisterFile("natter.proto", fileDescriptor_39129eec69c18084) }

var fileDescriptor_39129eec69c18084 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x4b, 0x4e, 0xc3, 0x30,
	0x10, 0x55, 0x3e, 0x4d, 0xdb, 0x11, 0xa4, 0xc2, 0x0b, 0x94, 0x15, 0xaa, 0x22, 0x90, 0xb2, 0x40,
	0x6c, 0x38, 0x01, 0xaa, 0x84, 0xd4, 0x6d, 0xca, 0x05, 0x4c, 0x3c, 0x82, 0x0a, 0x6a, 0x97, 0xb1,
	0x23, 0x0e, 0xc5, 0xb5, 0x38, 0x08, 0xea, 0xd8, 0x2d, 0x26, 0x48, 0xd9, 0xcd, 0x7b, 0xf3, 0xde,
	0x8c, 0xe7, 0x19, 0xce, 0xb4, 0x74, 0x0e, 0xe9, 0x6e, 0x4f, 0xc6, 0x19, 0x31, 0xdb, 0x6a, 0x87,
	0xa4, 0xe5, 0x7b, 0xdd, 0x40, 0xb9, 0x7a, 0xc5, 0xee, 0x6d, 0xab, 0x5b, 0xfc, 0xe8, 0xd1, 0x3a,
	0x71, 0x09, 0xc5, 0xc6, 0xf4, 0xd4, 0x61, 0x95, 0x2c, 0x93, 0x66, 0xde, 0x06, 0x54, 0xdf, 0xc0,
	0xe2, 0xa4, 0xb4, 0x7b, 0xa3, 0x2d, 0x0a, 0x01, 0xf9, 0x83, 0x52, 0x14, 0x84, 0x5c, 0xd7, 0xdf,
	0x09, 0x94, 0x8f, 0x86, 0x3e, 0x25, 0xa9, 0xe3, 0xc4, 0x12, 0xd2, 0xb5, 0x0a, 0xa2, 0x74, 0xad,
	0xa2, 0x0d, 0x69, 0xbc, 0x41, 0x5c, 0x01, 0xf8, 0x8a, 0x87, 0x66, 0xdc, 0x8b, 0x98, 0x83, 0xef,
	0x49, 0xd2, 0x0b, 0xba, 0x2a, 0xf7, 0x3e, 0x8f, 0x0e, 0x3e, 0x5f, 0xb1, 0x6f, 0xe2, 0x7d, 0xbf,
	0x8c, 0xb8, 0x85, 0x0b, 0x8f, 0xc2, 0xbb, 0x58, 0x56, 0xb0, 0xec, 0x7f, 0x43, 0x5c, 0xc3, 0xb9,
	0x27, 0x57, 0x66, 0xb7, 0x93, 0x5a, 0x55, 0xd3, 0x65, 0xd6, 0xcc, 0xdb, 0xbf, 0x64, 0xfd, 0x95,
	0xc0, 0xe2, 0x74, 0x66, 0x88, 0x63, 0x78, 0x67, 0x05, 0xd3, 0x4d, 0xdf, 0x75, 0x68, 0x2d, 0x1f,
	0x3a, 0x6b, 0x8f, 0x30, 0x4a, 0x20, 0x1b, 0x49, 0x20, 0x1f, 0x49, 0x60, 0x32, 0x92, 0x40, 0x31,
	0x4c, 0xe0, 0xb9, 0xe0, 0x6f, 0xbf, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x70, 0x2f, 0xcc,
	0x06, 0x02, 0x00, 0x00,
}