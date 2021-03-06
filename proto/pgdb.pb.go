// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pgdb.proto

package pgdb

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

type ServiceApp struct {
	Svc                  string   `protobuf:"bytes,1,opt,name=svc,proto3" json:"svc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceApp) Reset()         { *m = ServiceApp{} }
func (m *ServiceApp) String() string { return proto.CompactTextString(m) }
func (*ServiceApp) ProtoMessage()    {}
func (*ServiceApp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4560b3dcc472fe5, []int{0}
}

func (m *ServiceApp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceApp.Unmarshal(m, b)
}
func (m *ServiceApp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceApp.Marshal(b, m, deterministic)
}
func (m *ServiceApp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceApp.Merge(m, src)
}
func (m *ServiceApp) XXX_Size() int {
	return xxx_messageInfo_ServiceApp.Size(m)
}
func (m *ServiceApp) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceApp.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceApp proto.InternalMessageInfo

func (m *ServiceApp) GetSvc() string {
	if m != nil {
		return m.Svc
	}
	return ""
}

type ServiceDb struct {
	Host                 string      `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port                 string      `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	User                 string      `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Password             string      `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Dbname               string      `protobuf:"bytes,5,opt,name=dbname,proto3" json:"dbname,omitempty"`
	Svc                  *ServiceApp `protobuf:"bytes,6,opt,name=svc,proto3" json:"svc,omitempty"`
	ConnectionDb         string      `protobuf:"bytes,7,opt,name=connectionDb,proto3" json:"connectionDb,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ServiceDb) Reset()         { *m = ServiceDb{} }
func (m *ServiceDb) String() string { return proto.CompactTextString(m) }
func (*ServiceDb) ProtoMessage()    {}
func (*ServiceDb) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4560b3dcc472fe5, []int{1}
}

func (m *ServiceDb) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceDb.Unmarshal(m, b)
}
func (m *ServiceDb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceDb.Marshal(b, m, deterministic)
}
func (m *ServiceDb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceDb.Merge(m, src)
}
func (m *ServiceDb) XXX_Size() int {
	return xxx_messageInfo_ServiceDb.Size(m)
}
func (m *ServiceDb) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceDb.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceDb proto.InternalMessageInfo

func (m *ServiceDb) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *ServiceDb) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *ServiceDb) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *ServiceDb) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *ServiceDb) GetDbname() string {
	if m != nil {
		return m.Dbname
	}
	return ""
}

func (m *ServiceDb) GetSvc() *ServiceApp {
	if m != nil {
		return m.Svc
	}
	return nil
}

func (m *ServiceDb) GetConnectionDb() string {
	if m != nil {
		return m.ConnectionDb
	}
	return ""
}

func init() {
	proto.RegisterType((*ServiceApp)(nil), "pgdb.ServiceApp")
	proto.RegisterType((*ServiceDb)(nil), "pgdb.ServiceDb")
}

func init() { proto.RegisterFile("pgdb.proto", fileDescriptor_e4560b3dcc472fe5) }

var fileDescriptor_e4560b3dcc472fe5 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x41, 0x4e, 0xc4, 0x20,
	0x14, 0x86, 0xc5, 0xa9, 0x1d, 0xe7, 0x69, 0xe2, 0xf8, 0x16, 0x86, 0xcc, 0xc2, 0x34, 0xac, 0xba,
	0xea, 0xa2, 0x9e, 0xc0, 0x94, 0x13, 0xd8, 0x13, 0x14, 0x4a, 0xb4, 0x0b, 0x81, 0x00, 0xad, 0x07,
	0xf4, 0x62, 0x06, 0x4a, 0x6a, 0x9a, 0xd9, 0x7d, 0xef, 0xe3, 0xf1, 0xe7, 0xcf, 0x03, 0xb0, 0x9f,
	0xa3, 0x68, 0xac, 0x33, 0xc1, 0x60, 0x11, 0x99, 0xbd, 0x02, 0xf4, 0xca, 0x2d, 0x93, 0x54, 0xef,
	0xd6, 0xe2, 0x19, 0x0e, 0x7e, 0x91, 0x94, 0x54, 0xa4, 0x3e, 0x7d, 0x44, 0x64, 0xbf, 0x04, 0x4e,
	0x79, 0x81, 0x0b, 0x44, 0x28, 0xbe, 0x8c, 0x0f, 0x79, 0x21, 0x71, 0x74, 0xd6, 0xb8, 0x40, 0x6f,
	0x57, 0x17, 0x39, 0xba, 0xd9, 0x2b, 0x47, 0x0f, 0xab, 0x8b, 0x8c, 0x17, 0xb8, 0xb7, 0x83, 0xf7,
	0x3f, 0xc6, 0x8d, 0xb4, 0x48, 0x7e, 0x9b, 0xf1, 0x05, 0xca, 0x51, 0xe8, 0xe1, 0x5b, 0xd1, 0xbb,
	0xf4, 0x92, 0x27, 0x64, 0x6b, 0x9f, 0xb2, 0x22, 0xf5, 0x43, 0x7b, 0x6e, 0x52, 0xfb, 0xff, 0xba,
	0xa9, 0x21, 0x32, 0x78, 0x94, 0x46, 0x6b, 0x25, 0xc3, 0x64, 0x34, 0x17, 0xf4, 0x98, 0x12, 0x76,
	0xae, 0xed, 0xe0, 0x39, 0x7f, 0xeb, 0x36, 0x8d, 0x0d, 0x1c, 0x7b, 0x15, 0x66, 0xcb, 0x05, 0x5e,
	0x45, 0x5f, 0x9e, 0x76, 0x86, 0x0b, 0x76, 0x23, 0xca, 0x74, 0xb7, 0xb7, 0xbf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x29, 0xbb, 0xe0, 0x3f, 0x45, 0x01, 0x00, 0x00,
}
