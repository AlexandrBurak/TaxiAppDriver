// Code generated by protoc-gen-go. DO NOT EDIT.
// source: driver.proto

package grpc

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type Driver struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Phone                string   `protobuf:"bytes,2,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
	TaxiType             string   `protobuf:"bytes,4,opt,name=TaxiType,proto3" json:"TaxiType,omitempty"`
	Status               string   `protobuf:"bytes,5,opt,name=Status,proto3" json:"Status,omitempty"`
	Rating               int32    `protobuf:"varint,6,opt,name=Rating,proto3" json:"Rating,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Driver) Reset()         { *m = Driver{} }
func (m *Driver) String() string { return proto.CompactTextString(m) }
func (*Driver) ProtoMessage()    {}
func (*Driver) Descriptor() ([]byte, []int) {
	return fileDescriptor_521003751d596b5e, []int{0}
}

func (m *Driver) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Driver.Unmarshal(m, b)
}
func (m *Driver) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Driver.Marshal(b, m, deterministic)
}
func (m *Driver) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Driver.Merge(m, src)
}
func (m *Driver) XXX_Size() int {
	return xxx_messageInfo_Driver.Size(m)
}
func (m *Driver) XXX_DiscardUnknown() {
	xxx_messageInfo_Driver.DiscardUnknown(m)
}

var xxx_messageInfo_Driver proto.InternalMessageInfo

func (m *Driver) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Driver) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Driver) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Driver) GetTaxiType() string {
	if m != nil {
		return m.TaxiType
	}
	return ""
}

func (m *Driver) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Driver) GetRating() int32 {
	if m != nil {
		return m.Rating
	}
	return 0
}

type AllDrivers struct {
	Drivers              []*Driver `protobuf:"bytes,1,rep,name=drivers,proto3" json:"drivers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AllDrivers) Reset()         { *m = AllDrivers{} }
func (m *AllDrivers) String() string { return proto.CompactTextString(m) }
func (*AllDrivers) ProtoMessage()    {}
func (*AllDrivers) Descriptor() ([]byte, []int) {
	return fileDescriptor_521003751d596b5e, []int{1}
}

func (m *AllDrivers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllDrivers.Unmarshal(m, b)
}
func (m *AllDrivers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllDrivers.Marshal(b, m, deterministic)
}
func (m *AllDrivers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllDrivers.Merge(m, src)
}
func (m *AllDrivers) XXX_Size() int {
	return xxx_messageInfo_AllDrivers.Size(m)
}
func (m *AllDrivers) XXX_DiscardUnknown() {
	xxx_messageInfo_AllDrivers.DiscardUnknown(m)
}

var xxx_messageInfo_AllDrivers proto.InternalMessageInfo

func (m *AllDrivers) GetDrivers() []*Driver {
	if m != nil {
		return m.Drivers
	}
	return nil
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_521003751d596b5e, []int{2}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type AllFreeDrivers struct {
	Drivers              []*Driver `protobuf:"bytes,1,rep,name=drivers,proto3" json:"drivers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AllFreeDrivers) Reset()         { *m = AllFreeDrivers{} }
func (m *AllFreeDrivers) String() string { return proto.CompactTextString(m) }
func (*AllFreeDrivers) ProtoMessage()    {}
func (*AllFreeDrivers) Descriptor() ([]byte, []int) {
	return fileDescriptor_521003751d596b5e, []int{3}
}

func (m *AllFreeDrivers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllFreeDrivers.Unmarshal(m, b)
}
func (m *AllFreeDrivers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllFreeDrivers.Marshal(b, m, deterministic)
}
func (m *AllFreeDrivers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllFreeDrivers.Merge(m, src)
}
func (m *AllFreeDrivers) XXX_Size() int {
	return xxx_messageInfo_AllFreeDrivers.Size(m)
}
func (m *AllFreeDrivers) XXX_DiscardUnknown() {
	xxx_messageInfo_AllFreeDrivers.DiscardUnknown(m)
}

var xxx_messageInfo_AllFreeDrivers proto.InternalMessageInfo

func (m *AllFreeDrivers) GetDrivers() []*Driver {
	if m != nil {
		return m.Drivers
	}
	return nil
}

func init() {
	proto.RegisterType((*Driver)(nil), "api.Driver")
	proto.RegisterType((*AllDrivers)(nil), "api.AllDrivers")
	proto.RegisterType((*Request)(nil), "api.Request")
	proto.RegisterType((*AllFreeDrivers)(nil), "api.AllFreeDrivers")
}

func init() { proto.RegisterFile("driver.proto", fileDescriptor_521003751d596b5e) }

var fileDescriptor_521003751d596b5e = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x6d, 0x6b, 0x3b, 0xf7, 0x76, 0x9d, 0x18, 0x45, 0xc2, 0x4e, 0x25, 0x20, 0xf4, 0x54,
	0x65, 0x3b, 0xe8, 0x75, 0xe2, 0x1f, 0x76, 0x93, 0x6c, 0x27, 0x6f, 0x71, 0x7d, 0x99, 0x81, 0xb8,
	0xc6, 0x34, 0x1b, 0x8a, 0x5f, 0xc2, 0x8f, 0x2c, 0x4b, 0xba, 0x81, 0xe2, 0x65, 0xc7, 0xdf, 0xef,
	0xe9, 0xf3, 0xf6, 0x81, 0x40, 0xaf, 0x32, 0x72, 0x8d, 0xa6, 0xd4, 0xa6, 0xb6, 0x35, 0x89, 0x84,
	0x96, 0xec, 0x3b, 0x80, 0xe4, 0xce, 0x59, 0xd2, 0x87, 0x70, 0x52, 0xd1, 0x20, 0x0f, 0x8a, 0x2e,
	0x0f, 0x27, 0x15, 0x39, 0x83, 0xf8, 0xe9, 0xb5, 0x5e, 0x22, 0x0d, 0x9d, 0xf2, 0xb0, 0xb1, 0xf7,
	0x6f, 0x42, 0x2a, 0x1a, 0x79, 0xeb, 0x80, 0x0c, 0xe0, 0x68, 0x26, 0x3e, 0xe4, 0xec, 0x53, 0x23,
	0x3d, 0x74, 0xc1, 0x8e, 0xc9, 0x39, 0x24, 0x53, 0x2b, 0xec, 0xaa, 0xa1, 0xb1, 0x4b, 0x5a, 0xda,
	0x78, 0x2e, 0xac, 0x5c, 0x2e, 0x68, 0x92, 0x07, 0x45, 0xcc, 0x5b, 0x62, 0x23, 0x80, 0xb1, 0x52,
	0x7e, 0x54, 0x43, 0x2e, 0xa0, 0xe3, 0x57, 0x37, 0x34, 0xc8, 0xa3, 0x22, 0x1d, 0xa6, 0xa5, 0xd0,
	0xb2, 0xf4, 0x31, 0xdf, 0x66, 0xac, 0x0b, 0x1d, 0x8e, 0xef, 0x2b, 0x6c, 0x2c, 0xbb, 0x86, 0xfe,
	0x58, 0xa9, 0x07, 0x83, 0xb8, 0xdf, 0x8d, 0xe1, 0x17, 0x64, 0x5e, 0x4d, 0xd1, 0xac, 0xe5, 0x1c,
	0xc9, 0x0d, 0x9c, 0x3c, 0xa2, 0xfd, 0x73, 0xac, 0xe7, 0xba, 0xed, 0xcf, 0x06, 0xa7, 0x8e, 0x7e,
	0x7f, 0xc2, 0x0e, 0xc8, 0x15, 0x64, 0xbe, 0xf9, 0x7f, 0xeb, 0x78, 0xdb, 0xda, 0x35, 0x6e, 0xb3,
	0xe7, 0x74, 0x61, 0xf4, 0xfc, 0xd2, 0x8f, 0x79, 0x49, 0xdc, 0x1b, 0x8d, 0x7e, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x68, 0x04, 0xa7, 0x66, 0xb3, 0x01, 0x00, 0x00,
}
