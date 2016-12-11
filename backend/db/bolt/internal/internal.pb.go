// Code generated by protoc-gen-go.
// source: internal.proto
// DO NOT EDIT!

/*
Package internal is a generated protocol buffer package.

It is generated from these files:
	internal.proto

It has these top-level messages:
	Booking
*/
package internal

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Booking struct {
	CreateTime       int64  `protobuf:"varint,1,opt,name=CreateTime" json:"CreateTime,omitempty"`
	ModTime          int64  `protobuf:"varint,2,opt,name=ModTime" json:"ModTime,omitempty"`
	ID               string `protobuf:"bytes,3,opt,name=ID" json:"ID,omitempty"`
	BookingDate      string `protobuf:"bytes,4,opt,name=BookingDate" json:"BookingDate,omitempty"`
	RespContCustomer string `protobuf:"bytes,5,opt,name=RespContCustomer" json:"RespContCustomer,omitempty"`
	RespContSeller   string `protobuf:"bytes,6,opt,name=RespContSeller" json:"RespContSeller,omitempty"`
	ProjectCode      string `protobuf:"bytes,7,opt,name=ProjectCode" json:"ProjectCode,omitempty"`
}

func (m *Booking) Reset()                    { *m = Booking{} }
func (m *Booking) String() string            { return proto.CompactTextString(m) }
func (*Booking) ProtoMessage()               {}
func (*Booking) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*Booking)(nil), "internal.Booking")
}

func init() { proto.RegisterFile("internal.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x8f, 0xc1, 0x0a, 0x82, 0x40,
	0x10, 0x86, 0x59, 0x2d, 0xad, 0x09, 0x24, 0xf6, 0xb4, 0xa7, 0x90, 0x0e, 0x21, 0x1d, 0xba, 0xf4,
	0x06, 0xad, 0x17, 0x0f, 0x41, 0x58, 0x2f, 0x60, 0x39, 0x84, 0xa5, 0x3b, 0xb2, 0x4e, 0x4f, 0xdd,
	0x4b, 0x44, 0x5b, 0x82, 0xd4, 0xf1, 0xff, 0xbe, 0x9f, 0xf9, 0x19, 0x88, 0x2a, 0xc3, 0x68, 0x4d,
	0x51, 0x6f, 0x5a, 0x4b, 0x4c, 0x72, 0xd2, 0xe7, 0xe5, 0x53, 0x40, 0xb8, 0x23, 0xba, 0x57, 0xe6,
	0x2a, 0x17, 0x00, 0xda, 0x62, 0xc1, 0x78, 0xaa, 0x1a, 0x54, 0x22, 0x16, 0x89, 0x9f, 0x0f, 0x88,
	0x54, 0x10, 0xee, 0xa9, 0x74, 0xd2, 0x73, 0xb2, 0x8f, 0x32, 0x02, 0x2f, 0x4b, 0x95, 0x1f, 0x8b,
	0x64, 0x9a, 0x7b, 0x59, 0x2a, 0x63, 0x98, 0x7d, 0x8f, 0xa6, 0x05, 0xa3, 0x1a, 0x39, 0x31, 0x44,
	0x72, 0x0d, 0xf3, 0x1c, 0xbb, 0x56, 0x93, 0x61, 0xfd, 0xe8, 0x98, 0x1a, 0xb4, 0x6a, 0xec, 0x6a,
	0x7f, 0x5c, 0xae, 0x20, 0xea, 0xd9, 0x11, 0xeb, 0x1a, 0xad, 0x0a, 0x5c, 0xf3, 0x87, 0xbe, 0x57,
	0x0f, 0x96, 0x6e, 0x78, 0x61, 0x4d, 0x25, 0xaa, 0xf0, 0xb3, 0x3a, 0x40, 0xe7, 0xc0, 0xbd, 0xbf,
	0x7d, 0x05, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x00, 0xc2, 0x22, 0x10, 0x01, 0x00, 0x00,
}
