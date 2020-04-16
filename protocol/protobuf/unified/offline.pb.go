// Code generated by protoc-gen-go. DO NOT EDIT.
// source: steammessages_offline.steamclient.proto

package unified

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
// is compatible with the proto package unified is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package unified to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type COffline_GetOfflineLogonTicket_Request struct {
	Priority             *uint32  `protobuf:"varint,1,opt,name=priority" json:"priority,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *COffline_GetOfflineLogonTicket_Request) Reset() {
	*m = COffline_GetOfflineLogonTicket_Request{}
}
func (m *COffline_GetOfflineLogonTicket_Request) String() string { return proto.CompactTextString(m) }
func (*COffline_GetOfflineLogonTicket_Request) ProtoMessage()    {}
func (*COffline_GetOfflineLogonTicket_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_86478a2284c871a9, []int{0}
}

func (m *COffline_GetOfflineLogonTicket_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_COffline_GetOfflineLogonTicket_Request.Unmarshal(m, b)
}
func (m *COffline_GetOfflineLogonTicket_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_COffline_GetOfflineLogonTicket_Request.Marshal(b, m, deterministic)
}
func (m *COffline_GetOfflineLogonTicket_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_COffline_GetOfflineLogonTicket_Request.Merge(m, src)
}
func (m *COffline_GetOfflineLogonTicket_Request) XXX_Size() int {
	return xxx_messageInfo_COffline_GetOfflineLogonTicket_Request.Size(m)
}
func (m *COffline_GetOfflineLogonTicket_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_COffline_GetOfflineLogonTicket_Request.DiscardUnknown(m)
}

var xxx_messageInfo_COffline_GetOfflineLogonTicket_Request proto.InternalMessageInfo

func (m *COffline_GetOfflineLogonTicket_Request) GetPriority() uint32 {
	if m != nil && m.Priority != nil {
		return *m.Priority
	}
	return 0
}

type COffline_GetOfflineLogonTicket_Response struct {
	SerializedTicket     []byte   `protobuf:"bytes,1,opt,name=serialized_ticket,json=serializedTicket" json:"serialized_ticket,omitempty"`
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *COffline_GetOfflineLogonTicket_Response) Reset() {
	*m = COffline_GetOfflineLogonTicket_Response{}
}
func (m *COffline_GetOfflineLogonTicket_Response) String() string { return proto.CompactTextString(m) }
func (*COffline_GetOfflineLogonTicket_Response) ProtoMessage()    {}
func (*COffline_GetOfflineLogonTicket_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_86478a2284c871a9, []int{1}
}

func (m *COffline_GetOfflineLogonTicket_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_COffline_GetOfflineLogonTicket_Response.Unmarshal(m, b)
}
func (m *COffline_GetOfflineLogonTicket_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_COffline_GetOfflineLogonTicket_Response.Marshal(b, m, deterministic)
}
func (m *COffline_GetOfflineLogonTicket_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_COffline_GetOfflineLogonTicket_Response.Merge(m, src)
}
func (m *COffline_GetOfflineLogonTicket_Response) XXX_Size() int {
	return xxx_messageInfo_COffline_GetOfflineLogonTicket_Response.Size(m)
}
func (m *COffline_GetOfflineLogonTicket_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_COffline_GetOfflineLogonTicket_Response.DiscardUnknown(m)
}

var xxx_messageInfo_COffline_GetOfflineLogonTicket_Response proto.InternalMessageInfo

func (m *COffline_GetOfflineLogonTicket_Response) GetSerializedTicket() []byte {
	if m != nil {
		return m.SerializedTicket
	}
	return nil
}

func (m *COffline_GetOfflineLogonTicket_Response) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type COffline_GetUnsignedOfflineLogonTicket_Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *COffline_GetUnsignedOfflineLogonTicket_Request) Reset() {
	*m = COffline_GetUnsignedOfflineLogonTicket_Request{}
}
func (m *COffline_GetUnsignedOfflineLogonTicket_Request) String() string {
	return proto.CompactTextString(m)
}
func (*COffline_GetUnsignedOfflineLogonTicket_Request) ProtoMessage() {}
func (*COffline_GetUnsignedOfflineLogonTicket_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_86478a2284c871a9, []int{2}
}

func (m *COffline_GetUnsignedOfflineLogonTicket_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_COffline_GetUnsignedOfflineLogonTicket_Request.Unmarshal(m, b)
}
func (m *COffline_GetUnsignedOfflineLogonTicket_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_COffline_GetUnsignedOfflineLogonTicket_Request.Marshal(b, m, deterministic)
}
func (m *COffline_GetUnsignedOfflineLogonTicket_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_COffline_GetUnsignedOfflineLogonTicket_Request.Merge(m, src)
}
func (m *COffline_GetUnsignedOfflineLogonTicket_Request) XXX_Size() int {
	return xxx_messageInfo_COffline_GetUnsignedOfflineLogonTicket_Request.Size(m)
}
func (m *COffline_GetUnsignedOfflineLogonTicket_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_COffline_GetUnsignedOfflineLogonTicket_Request.DiscardUnknown(m)
}

var xxx_messageInfo_COffline_GetUnsignedOfflineLogonTicket_Request proto.InternalMessageInfo

type COffline_OfflineLogonTicket struct {
	Accountid            *uint32  `protobuf:"varint,1,opt,name=accountid" json:"accountid,omitempty"`
	Rtime32CreationTime  *uint32  `protobuf:"fixed32,2,opt,name=rtime32_creation_time,json=rtime32CreationTime" json:"rtime32_creation_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *COffline_OfflineLogonTicket) Reset()         { *m = COffline_OfflineLogonTicket{} }
func (m *COffline_OfflineLogonTicket) String() string { return proto.CompactTextString(m) }
func (*COffline_OfflineLogonTicket) ProtoMessage()    {}
func (*COffline_OfflineLogonTicket) Descriptor() ([]byte, []int) {
	return fileDescriptor_86478a2284c871a9, []int{3}
}

func (m *COffline_OfflineLogonTicket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_COffline_OfflineLogonTicket.Unmarshal(m, b)
}
func (m *COffline_OfflineLogonTicket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_COffline_OfflineLogonTicket.Marshal(b, m, deterministic)
}
func (m *COffline_OfflineLogonTicket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_COffline_OfflineLogonTicket.Merge(m, src)
}
func (m *COffline_OfflineLogonTicket) XXX_Size() int {
	return xxx_messageInfo_COffline_OfflineLogonTicket.Size(m)
}
func (m *COffline_OfflineLogonTicket) XXX_DiscardUnknown() {
	xxx_messageInfo_COffline_OfflineLogonTicket.DiscardUnknown(m)
}

var xxx_messageInfo_COffline_OfflineLogonTicket proto.InternalMessageInfo

func (m *COffline_OfflineLogonTicket) GetAccountid() uint32 {
	if m != nil && m.Accountid != nil {
		return *m.Accountid
	}
	return 0
}

func (m *COffline_OfflineLogonTicket) GetRtime32CreationTime() uint32 {
	if m != nil && m.Rtime32CreationTime != nil {
		return *m.Rtime32CreationTime
	}
	return 0
}

type COffline_GetUnsignedOfflineLogonTicket_Response struct {
	Ticket               *COffline_OfflineLogonTicket `protobuf:"bytes,1,opt,name=ticket" json:"ticket,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *COffline_GetUnsignedOfflineLogonTicket_Response) Reset() {
	*m = COffline_GetUnsignedOfflineLogonTicket_Response{}
}
func (m *COffline_GetUnsignedOfflineLogonTicket_Response) String() string {
	return proto.CompactTextString(m)
}
func (*COffline_GetUnsignedOfflineLogonTicket_Response) ProtoMessage() {}
func (*COffline_GetUnsignedOfflineLogonTicket_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_86478a2284c871a9, []int{4}
}

func (m *COffline_GetUnsignedOfflineLogonTicket_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_COffline_GetUnsignedOfflineLogonTicket_Response.Unmarshal(m, b)
}
func (m *COffline_GetUnsignedOfflineLogonTicket_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_COffline_GetUnsignedOfflineLogonTicket_Response.Marshal(b, m, deterministic)
}
func (m *COffline_GetUnsignedOfflineLogonTicket_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_COffline_GetUnsignedOfflineLogonTicket_Response.Merge(m, src)
}
func (m *COffline_GetUnsignedOfflineLogonTicket_Response) XXX_Size() int {
	return xxx_messageInfo_COffline_GetUnsignedOfflineLogonTicket_Response.Size(m)
}
func (m *COffline_GetUnsignedOfflineLogonTicket_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_COffline_GetUnsignedOfflineLogonTicket_Response.DiscardUnknown(m)
}

var xxx_messageInfo_COffline_GetUnsignedOfflineLogonTicket_Response proto.InternalMessageInfo

func (m *COffline_GetUnsignedOfflineLogonTicket_Response) GetTicket() *COffline_OfflineLogonTicket {
	if m != nil {
		return m.Ticket
	}
	return nil
}

func init() {
	proto.RegisterType((*COffline_GetOfflineLogonTicket_Request)(nil), "COffline_GetOfflineLogonTicket_Request")
	proto.RegisterType((*COffline_GetOfflineLogonTicket_Response)(nil), "COffline_GetOfflineLogonTicket_Response")
	proto.RegisterType((*COffline_GetUnsignedOfflineLogonTicket_Request)(nil), "COffline_GetUnsignedOfflineLogonTicket_Request")
	proto.RegisterType((*COffline_OfflineLogonTicket)(nil), "COffline_OfflineLogonTicket")
	proto.RegisterType((*COffline_GetUnsignedOfflineLogonTicket_Response)(nil), "COffline_GetUnsignedOfflineLogonTicket_Response")
}

func init() {
	proto.RegisterFile("steammessages_offline.steamclient.proto", fileDescriptor_86478a2284c871a9)
}

var fileDescriptor_86478a2284c871a9 = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcd, 0xca, 0xd3, 0x40,
	0x14, 0x25, 0x7e, 0xe0, 0xa7, 0xa3, 0x82, 0x8e, 0x7c, 0x10, 0x62, 0x85, 0x21, 0x0b, 0x5b, 0x50,
	0xd2, 0x12, 0xdd, 0xb8, 0x70, 0x63, 0x95, 0x22, 0x08, 0x42, 0xa8, 0xeb, 0x30, 0x26, 0x37, 0xf1,
	0x62, 0x32, 0x53, 0x67, 0x6e, 0x04, 0x5d, 0x49, 0x5f, 0xc5, 0x67, 0xc8, 0x03, 0xf8, 0x66, 0x92,
	0x64, 0x68, 0x2a, 0xd6, 0x7e, 0xed, 0x2e, 0x73, 0xcf, 0xb9, 0x3f, 0xe7, 0x9c, 0xb0, 0xa9, 0x25,
	0x90, 0x75, 0x0d, 0xd6, 0xca, 0x12, 0x6c, 0xaa, 0x8b, 0xa2, 0x42, 0x05, 0x51, 0x5f, 0xcd, 0x2a,
	0x04, 0x45, 0xd1, 0xc6, 0x68, 0xd2, 0xc1, 0xb3, 0xbf, 0x89, 0x8d, 0xc2, 0x02, 0x21, 0x4f, 0x3f,
	0x49, 0x7b, 0x80, 0x1d, 0xbe, 0x61, 0x4f, 0x96, 0x1f, 0x86, 0x59, 0xe9, 0x0a, 0xc8, 0x7d, 0xbe,
	0xd7, 0xa5, 0x56, 0x6b, 0xcc, 0xbe, 0x00, 0xa5, 0x09, 0x7c, 0x6d, 0xc0, 0x12, 0x0f, 0xd8, 0xad,
	0x8d, 0x41, 0x6d, 0x90, 0xbe, 0xfb, 0x9e, 0xf0, 0x66, 0xf7, 0x92, 0xdd, 0x3b, 0x24, 0x36, 0xbd,
	0x76, 0x8a, 0xdd, 0x68, 0x65, 0x81, 0x3f, 0x65, 0x0f, 0x2c, 0x18, 0x94, 0x15, 0xfe, 0x80, 0x3c,
	0xa5, 0x1e, 0xed, 0xe7, 0xdd, 0x4d, 0xee, 0x8f, 0xc0, 0xd0, 0xc5, 0x27, 0xec, 0xb6, 0xc5, 0x52,
	0x49, 0x6a, 0x0c, 0xf8, 0x37, 0x7a, 0xd2, 0x58, 0x08, 0x17, 0x2c, 0xda, 0xdf, 0xfa, 0x51, 0x75,
	0x10, 0xe4, 0xff, 0xd7, 0x10, 0x6a, 0xf6, 0x68, 0xd7, 0xf1, 0x2f, 0xad, 0x5b, 0x27, 0xb3, 0x4c,
	0x37, 0x8a, 0x30, 0x77, 0x1a, 0xc7, 0x02, 0x8f, 0xd9, 0x95, 0x21, 0xac, 0xe1, 0x79, 0x9c, 0x66,
	0x06, 0x24, 0xa1, 0x56, 0x69, 0xf7, 0xee, 0x0f, 0xbb, 0x4c, 0x1e, 0x3a, 0x70, 0xe9, 0xb0, 0x35,
	0xd6, 0x10, 0x96, 0x6c, 0x7e, 0xf2, 0x89, 0xce, 0xa0, 0x17, 0xec, 0xe6, 0x9e, 0x2b, 0x77, 0xe2,
	0x49, 0x74, 0xe4, 0xe4, 0xc4, 0x71, 0xe3, 0x5f, 0x17, 0xec, 0xd2, 0xc1, 0xbc, 0xf5, 0xd8, 0xd5,
	0xc1, 0x14, 0xf8, 0x34, 0x3a, 0x2d, 0xec, 0x60, 0x16, 0x9d, 0x98, 0x67, 0xf8, 0x6e, 0xdb, 0xfa,
	0x6f, 0x57, 0x40, 0x42, 0x8a, 0x31, 0x3e, 0x21, 0x55, 0x2e, 0x06, 0xa9, 0xc2, 0xfd, 0xa9, 0xa2,
	0xea, 0xba, 0xc5, 0x70, 0xae, 0x28, 0xb4, 0x11, 0xf4, 0x19, 0x44, 0xd6, 0x18, 0x03, 0x8a, 0x44,
	0x63, 0xc1, 0xf0, 0xdf, 0x1e, 0x7b, 0x7c, 0xd4, 0x24, 0x3e, 0x3f, 0x33, 0xf0, 0x60, 0x11, 0x9d,
	0x69, 0x7f, 0xf8, 0x6a, 0xdb, 0xfa, 0x2f, 0x7b, 0x3d, 0x4a, 0x34, 0xea, 0x5c, 0x0d, 0xc1, 0x64,
	0xdb, 0xfa, 0xbe, 0x9b, 0x2f, 0x2c, 0x10, 0xa1, 0x2a, 0x6d, 0xe7, 0xcc, 0x37, 0xcc, 0xe0, 0xf5,
	0xc5, 0x4f, 0xcf, 0xfb, 0x13, 0x00, 0x00, 0xff, 0xff, 0x17, 0x0d, 0x59, 0x6e, 0xca, 0x03, 0x00,
	0x00,
}
