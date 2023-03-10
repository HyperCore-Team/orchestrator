// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wrap_request_znn.proto

package events

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

type WrapRequestZnnProto struct {
	NetworkClass         uint32   `protobuf:"varint,1,opt,name=networkClass,proto3" json:"networkClass,omitempty"`
	ChainId              uint32   `protobuf:"varint,2,opt,name=chainId,proto3" json:"chainId,omitempty"`
	Id                   []byte   `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	ToAddress            string   `protobuf:"bytes,4,opt,name=toAddress,proto3" json:"toAddress,omitempty"`
	TokenAddress         string   `protobuf:"bytes,5,opt,name=tokenAddress,proto3" json:"tokenAddress,omitempty"`
	Amount               []byte   `protobuf:"bytes,6,opt,name=amount,proto3" json:"amount,omitempty"`
	Fee                  []byte   `protobuf:"bytes,7,opt,name=fee,proto3" json:"fee,omitempty"`
	Signature            string   `protobuf:"bytes,8,opt,name=signature,proto3" json:"signature,omitempty"`
	Status               uint32   `protobuf:"varint,9,opt,name=status,proto3" json:"status,omitempty"`
	SentSignature        bool     `protobuf:"varint,10,opt,name=sentSignature,proto3" json:"sentSignature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WrapRequestZnnProto) Reset()         { *m = WrapRequestZnnProto{} }
func (m *WrapRequestZnnProto) String() string { return proto.CompactTextString(m) }
func (*WrapRequestZnnProto) ProtoMessage()    {}
func (*WrapRequestZnnProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_e589c5a77ce52b4d, []int{0}
}

func (m *WrapRequestZnnProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WrapRequestZnnProto.Unmarshal(m, b)
}
func (m *WrapRequestZnnProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WrapRequestZnnProto.Marshal(b, m, deterministic)
}
func (m *WrapRequestZnnProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WrapRequestZnnProto.Merge(m, src)
}
func (m *WrapRequestZnnProto) XXX_Size() int {
	return xxx_messageInfo_WrapRequestZnnProto.Size(m)
}
func (m *WrapRequestZnnProto) XXX_DiscardUnknown() {
	xxx_messageInfo_WrapRequestZnnProto.DiscardUnknown(m)
}

var xxx_messageInfo_WrapRequestZnnProto proto.InternalMessageInfo

func (m *WrapRequestZnnProto) GetNetworkClass() uint32 {
	if m != nil {
		return m.NetworkClass
	}
	return 0
}

func (m *WrapRequestZnnProto) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *WrapRequestZnnProto) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *WrapRequestZnnProto) GetToAddress() string {
	if m != nil {
		return m.ToAddress
	}
	return ""
}

func (m *WrapRequestZnnProto) GetTokenAddress() string {
	if m != nil {
		return m.TokenAddress
	}
	return ""
}

func (m *WrapRequestZnnProto) GetAmount() []byte {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *WrapRequestZnnProto) GetFee() []byte {
	if m != nil {
		return m.Fee
	}
	return nil
}

func (m *WrapRequestZnnProto) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *WrapRequestZnnProto) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *WrapRequestZnnProto) GetSentSignature() bool {
	if m != nil {
		return m.SentSignature
	}
	return false
}

func init() {
	proto.RegisterType((*WrapRequestZnnProto)(nil), "events.WrapRequestZnnProto")
}

func init() { proto.RegisterFile("wrap_request_znn.proto", fileDescriptor_e589c5a77ce52b4d) }

var fileDescriptor_e589c5a77ce52b4d = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0xe5, 0x14, 0xd2, 0xe4, 0x68, 0x11, 0x32, 0x52, 0xe5, 0x81, 0xc1, 0xaa, 0x18, 0x32,
	0x95, 0x81, 0x5f, 0x00, 0x1b, 0x1b, 0x32, 0x48, 0x48, 0x5d, 0x2a, 0x83, 0x0f, 0x88, 0x2a, 0xce,
	0xc1, 0xbe, 0x50, 0xc1, 0x3f, 0xe2, 0x5f, 0xa2, 0x38, 0x6d, 0x69, 0x36, 0xbf, 0xef, 0x9e, 0xef,
	0x9d, 0x1e, 0xcc, 0x36, 0xc1, 0x36, 0xab, 0x80, 0x9f, 0x2d, 0x46, 0x5e, 0xfd, 0x10, 0x2d, 0x9a,
	0xe0, 0xd9, 0xcb, 0x1c, 0xbf, 0x90, 0x38, 0xce, 0x7f, 0x33, 0x38, 0x7f, 0x0a, 0xb6, 0x31, 0xbd,
	0x63, 0x49, 0x74, 0x9f, 0xe6, 0x1a, 0x4e, 0x08, 0x79, 0xe3, 0xc3, 0xfa, 0xf1, 0xbb, 0x41, 0x25,
	0xb4, 0xa8, 0xa6, 0xe6, 0x10, 0x49, 0x05, 0xe3, 0x97, 0x77, 0x5b, 0xd3, 0x9d, 0x53, 0x59, 0x9a,
	0xee, 0xa4, 0x3c, 0x85, 0xac, 0x76, 0x6a, 0xa4, 0x45, 0x35, 0x31, 0x59, 0xed, 0xe4, 0x05, 0x94,
	0xec, 0x6f, 0x9c, 0x0b, 0x18, 0xa3, 0x3a, 0xd2, 0xa2, 0x2a, 0xcd, 0x3f, 0x90, 0x73, 0x98, 0xb0,
	0x5f, 0x23, 0xed, 0x0c, 0xc7, 0xc9, 0x30, 0x60, 0x72, 0x06, 0xb9, 0xfd, 0xf0, 0x2d, 0xb1, 0xca,
	0xd3, 0xd6, 0xad, 0x92, 0x67, 0x30, 0x7a, 0x45, 0x54, 0xe3, 0x04, 0xbb, 0x67, 0x97, 0x15, 0xeb,
	0x37, 0xb2, 0xdc, 0x06, 0x54, 0x45, 0x9f, 0xb5, 0x07, 0xdd, 0x9e, 0xc8, 0x96, 0xdb, 0xa8, 0xca,
	0x74, 0xf2, 0x56, 0xc9, 0x4b, 0x98, 0x46, 0x24, 0x7e, 0xd8, 0xff, 0x04, 0x2d, 0xaa, 0xc2, 0x0c,
	0xe1, 0x2d, 0x2c, 0x8b, 0xc5, 0x55, 0xdf, 0xdb, 0x73, 0x9e, 0x6a, 0xbc, 0xfe, 0x0b, 0x00, 0x00,
	0xff, 0xff, 0xb1, 0x27, 0xcf, 0xb1, 0x60, 0x01, 0x00, 0x00,
}
