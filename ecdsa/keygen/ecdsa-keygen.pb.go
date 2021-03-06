// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protob/ecdsa-keygen.proto

package keygen

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

//
// Represents a BROADCAST message sent during Round 1 of the ECDSA TSS keygen protocol.
type KGRound1Message struct {
	Commitment           []byte   `protobuf:"bytes,1,opt,name=commitment,proto3" json:"commitment,omitempty"`
	PaillierN            []byte   `protobuf:"bytes,2,opt,name=paillier_n,json=paillierN,proto3" json:"paillier_n,omitempty"`
	NTilde               []byte   `protobuf:"bytes,3,opt,name=n_tilde,json=nTilde,proto3" json:"n_tilde,omitempty"`
	H1                   []byte   `protobuf:"bytes,4,opt,name=h1,proto3" json:"h1,omitempty"`
	H2                   []byte   `protobuf:"bytes,5,opt,name=h2,proto3" json:"h2,omitempty"`
	Dlnproof_1           [][]byte `protobuf:"bytes,6,rep,name=dlnproof_1,json=dlnproof1,proto3" json:"dlnproof_1,omitempty"`
	Dlnproof_2           [][]byte `protobuf:"bytes,7,rep,name=dlnproof_2,json=dlnproof2,proto3" json:"dlnproof_2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KGRound1Message) Reset()         { *m = KGRound1Message{} }
func (m *KGRound1Message) String() string { return proto.CompactTextString(m) }
func (*KGRound1Message) ProtoMessage()    {}
func (*KGRound1Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a2e19e981cdbb01, []int{0}
}

func (m *KGRound1Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KGRound1Message.Unmarshal(m, b)
}
func (m *KGRound1Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KGRound1Message.Marshal(b, m, deterministic)
}
func (m *KGRound1Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KGRound1Message.Merge(m, src)
}
func (m *KGRound1Message) XXX_Size() int {
	return xxx_messageInfo_KGRound1Message.Size(m)
}
func (m *KGRound1Message) XXX_DiscardUnknown() {
	xxx_messageInfo_KGRound1Message.DiscardUnknown(m)
}

var xxx_messageInfo_KGRound1Message proto.InternalMessageInfo

func (m *KGRound1Message) GetCommitment() []byte {
	if m != nil {
		return m.Commitment
	}
	return nil
}

func (m *KGRound1Message) GetPaillierN() []byte {
	if m != nil {
		return m.PaillierN
	}
	return nil
}

func (m *KGRound1Message) GetNTilde() []byte {
	if m != nil {
		return m.NTilde
	}
	return nil
}

func (m *KGRound1Message) GetH1() []byte {
	if m != nil {
		return m.H1
	}
	return nil
}

func (m *KGRound1Message) GetH2() []byte {
	if m != nil {
		return m.H2
	}
	return nil
}

func (m *KGRound1Message) GetDlnproof_1() [][]byte {
	if m != nil {
		return m.Dlnproof_1
	}
	return nil
}

func (m *KGRound1Message) GetDlnproof_2() [][]byte {
	if m != nil {
		return m.Dlnproof_2
	}
	return nil
}

//
// Represents a P2P message sent to each party during Round 2 of the ECDSA TSS keygen protocol.
type KGRound2Message1 struct {
	Share                []byte   `protobuf:"bytes,1,opt,name=share,proto3" json:"share,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KGRound2Message1) Reset()         { *m = KGRound2Message1{} }
func (m *KGRound2Message1) String() string { return proto.CompactTextString(m) }
func (*KGRound2Message1) ProtoMessage()    {}
func (*KGRound2Message1) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a2e19e981cdbb01, []int{1}
}

func (m *KGRound2Message1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KGRound2Message1.Unmarshal(m, b)
}
func (m *KGRound2Message1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KGRound2Message1.Marshal(b, m, deterministic)
}
func (m *KGRound2Message1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KGRound2Message1.Merge(m, src)
}
func (m *KGRound2Message1) XXX_Size() int {
	return xxx_messageInfo_KGRound2Message1.Size(m)
}
func (m *KGRound2Message1) XXX_DiscardUnknown() {
	xxx_messageInfo_KGRound2Message1.DiscardUnknown(m)
}

var xxx_messageInfo_KGRound2Message1 proto.InternalMessageInfo

func (m *KGRound2Message1) GetShare() []byte {
	if m != nil {
		return m.Share
	}
	return nil
}

//
// Represents a BROADCAST message sent to each party during Round 2 of the ECDSA TSS keygen protocol.
type KGRound2Message2 struct {
	DeCommitment         [][]byte `protobuf:"bytes,1,rep,name=de_commitment,json=deCommitment,proto3" json:"de_commitment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KGRound2Message2) Reset()         { *m = KGRound2Message2{} }
func (m *KGRound2Message2) String() string { return proto.CompactTextString(m) }
func (*KGRound2Message2) ProtoMessage()    {}
func (*KGRound2Message2) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a2e19e981cdbb01, []int{2}
}

func (m *KGRound2Message2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KGRound2Message2.Unmarshal(m, b)
}
func (m *KGRound2Message2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KGRound2Message2.Marshal(b, m, deterministic)
}
func (m *KGRound2Message2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KGRound2Message2.Merge(m, src)
}
func (m *KGRound2Message2) XXX_Size() int {
	return xxx_messageInfo_KGRound2Message2.Size(m)
}
func (m *KGRound2Message2) XXX_DiscardUnknown() {
	xxx_messageInfo_KGRound2Message2.DiscardUnknown(m)
}

var xxx_messageInfo_KGRound2Message2 proto.InternalMessageInfo

func (m *KGRound2Message2) GetDeCommitment() [][]byte {
	if m != nil {
		return m.DeCommitment
	}
	return nil
}

//
// Represents a BROADCAST message sent to each party during Round 3 of the ECDSA TSS keygen protocol.
type KGRound3Message struct {
	PaillierProof        [][]byte `protobuf:"bytes,1,rep,name=paillier_proof,json=paillierProof,proto3" json:"paillier_proof,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KGRound3Message) Reset()         { *m = KGRound3Message{} }
func (m *KGRound3Message) String() string { return proto.CompactTextString(m) }
func (*KGRound3Message) ProtoMessage()    {}
func (*KGRound3Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a2e19e981cdbb01, []int{3}
}

func (m *KGRound3Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KGRound3Message.Unmarshal(m, b)
}
func (m *KGRound3Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KGRound3Message.Marshal(b, m, deterministic)
}
func (m *KGRound3Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KGRound3Message.Merge(m, src)
}
func (m *KGRound3Message) XXX_Size() int {
	return xxx_messageInfo_KGRound3Message.Size(m)
}
func (m *KGRound3Message) XXX_DiscardUnknown() {
	xxx_messageInfo_KGRound3Message.DiscardUnknown(m)
}

var xxx_messageInfo_KGRound3Message proto.InternalMessageInfo

func (m *KGRound3Message) GetPaillierProof() [][]byte {
	if m != nil {
		return m.PaillierProof
	}
	return nil
}

func init() {
	proto.RegisterType((*KGRound1Message)(nil), "KGRound1Message")
	proto.RegisterType((*KGRound2Message1)(nil), "KGRound2Message1")
	proto.RegisterType((*KGRound2Message2)(nil), "KGRound2Message2")
	proto.RegisterType((*KGRound3Message)(nil), "KGRound3Message")
}

func init() { proto.RegisterFile("protob/ecdsa-keygen.proto", fileDescriptor_1a2e19e981cdbb01) }

var fileDescriptor_1a2e19e981cdbb01 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x49, 0x6a, 0x53, 0x1c, 0xd2, 0x28, 0x8b, 0xe0, 0x7a, 0x50, 0x4a, 0x44, 0xe8, 0x45,
	0xcb, 0x6e, 0x0f, 0x7a, 0xd6, 0x83, 0x07, 0x51, 0xa4, 0x78, 0xf2, 0x12, 0xd2, 0xee, 0xd8, 0x04,
	0x93, 0xdd, 0x90, 0xc4, 0x83, 0xbf, 0xd0, 0xbf, 0x25, 0x3b, 0xd9, 0x04, 0x83, 0xc7, 0xf7, 0xbd,
	0x64, 0xd8, 0xf7, 0xc1, 0x59, 0x55, 0x9b, 0xd6, 0x6c, 0x57, 0xb8, 0x53, 0x4d, 0x7a, 0xfd, 0x89,
	0xdf, 0x7b, 0xd4, 0x37, 0xc4, 0xe2, 0x1f, 0x0f, 0x8e, 0x9e, 0x1e, 0x37, 0xe6, 0x4b, 0x2b, 0xf1,
	0x8c, 0x4d, 0x93, 0xee, 0x91, 0x5d, 0x00, 0xec, 0x4c, 0x59, 0xe6, 0x6d, 0x89, 0xba, 0xe5, 0xde,
	0xc2, 0x5b, 0x86, 0x9b, 0x3f, 0x84, 0x9d, 0x03, 0x54, 0x69, 0x5e, 0x14, 0x39, 0xd6, 0x89, 0xe6,
	0x3e, 0xf5, 0x87, 0x3d, 0x79, 0x61, 0xa7, 0x30, 0xd3, 0x49, 0x9b, 0x17, 0x0a, 0xf9, 0x84, 0xba,
	0x40, 0xbf, 0xd9, 0xc4, 0x22, 0xf0, 0x33, 0xc1, 0x0f, 0x88, 0xf9, 0x99, 0xa0, 0x2c, 0xf9, 0xd4,
	0x65, 0x69, 0xef, 0xaa, 0x42, 0x57, 0xb5, 0x31, 0x1f, 0x89, 0xe0, 0xc1, 0x62, 0x62, 0xef, 0xf6,
	0x44, 0x8c, 0x6a, 0xc9, 0x67, 0xe3, 0x5a, 0xc6, 0x4b, 0x38, 0x76, 0x43, 0xa4, 0x1b, 0x22, 0xd8,
	0x09, 0x4c, 0x9b, 0x2c, 0xad, 0xd1, 0x8d, 0xe8, 0x42, 0x7c, 0xfb, 0xef, 0x4b, 0xc9, 0x2e, 0x61,
	0xae, 0x30, 0x19, 0xcd, 0xb6, 0xf7, 0x43, 0x85, 0x0f, 0x03, 0x8b, 0xef, 0x06, 0x57, 0xeb, 0xde,
	0xd5, 0x15, 0x44, 0x83, 0x0b, 0x7a, 0x88, 0xfb, 0x71, 0xde, 0xd3, 0x57, 0x0b, 0xef, 0xa3, 0xf7,
	0x90, 0xe4, 0xaf, 0x3a, 0xf9, 0xdb, 0x80, 0xec, 0xaf, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe9,
	0x13, 0x63, 0x79, 0x9a, 0x01, 0x00, 0x00,
}
