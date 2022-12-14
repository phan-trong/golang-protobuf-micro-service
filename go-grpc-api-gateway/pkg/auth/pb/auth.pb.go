// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/auth/pb/auth.proto

package pb

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

// Register
type RegisterRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad076042eaab0678, []int{0}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RegisterResponse struct {
	Status               int64    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad076042eaab0678, []int{1}
}

func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *RegisterResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

// Login
type LoginRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad076042eaab0678, []int{2}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	Status               int64    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad076042eaab0678, []int{3}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *LoginResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// Validate
type ValidateRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateRequest) Reset()         { *m = ValidateRequest{} }
func (m *ValidateRequest) String() string { return proto.CompactTextString(m) }
func (*ValidateRequest) ProtoMessage()    {}
func (*ValidateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad076042eaab0678, []int{4}
}

func (m *ValidateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateRequest.Unmarshal(m, b)
}
func (m *ValidateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateRequest.Marshal(b, m, deterministic)
}
func (m *ValidateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateRequest.Merge(m, src)
}
func (m *ValidateRequest) XXX_Size() int {
	return xxx_messageInfo_ValidateRequest.Size(m)
}
func (m *ValidateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateRequest proto.InternalMessageInfo

func (m *ValidateRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ValidateResponse struct {
	Status               int64    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	UserId               int64    `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateResponse) Reset()         { *m = ValidateResponse{} }
func (m *ValidateResponse) String() string { return proto.CompactTextString(m) }
func (*ValidateResponse) ProtoMessage()    {}
func (*ValidateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad076042eaab0678, []int{5}
}

func (m *ValidateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateResponse.Unmarshal(m, b)
}
func (m *ValidateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateResponse.Marshal(b, m, deterministic)
}
func (m *ValidateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateResponse.Merge(m, src)
}
func (m *ValidateResponse) XXX_Size() int {
	return xxx_messageInfo_ValidateResponse.Size(m)
}
func (m *ValidateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateResponse proto.InternalMessageInfo

func (m *ValidateResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ValidateResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ValidateResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "auth.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "auth.RegisterResponse")
	proto.RegisterType((*LoginRequest)(nil), "auth.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "auth.LoginResponse")
	proto.RegisterType((*ValidateRequest)(nil), "auth.ValidateRequest")
	proto.RegisterType((*ValidateResponse)(nil), "auth.ValidateResponse")
}

func init() {
	proto.RegisterFile("pkg/auth/pb/auth.proto", fileDescriptor_ad076042eaab0678)
}

var fileDescriptor_ad076042eaab0678 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x4d, 0x4f, 0x83, 0x40,
	0x14, 0x2c, 0xd6, 0x12, 0x7c, 0xda, 0xd0, 0xac, 0x4a, 0x08, 0x27, 0xb3, 0x17, 0x3d, 0x41, 0x52,
	0x8f, 0xbd, 0x54, 0x3d, 0x99, 0x78, 0xa2, 0x89, 0x31, 0xde, 0xb6, 0xf2, 0x42, 0x49, 0x2b, 0xbb,
	0xee, 0x87, 0xfe, 0x38, 0xff, 0x9c, 0xe9, 0x2e, 0xb4, 0xc8, 0x91, 0x13, 0xcc, 0x30, 0xcc, 0xce,
	0xce, 0x7b, 0x10, 0x89, 0x6d, 0x99, 0x31, 0xa3, 0x37, 0x99, 0x58, 0xdb, 0x67, 0x2a, 0x24, 0xd7,
	0x9c, 0x9c, 0xee, 0xdf, 0xe9, 0x13, 0x84, 0x39, 0x96, 0x95, 0xd2, 0x28, 0x73, 0xfc, 0x32, 0xa8,
	0x34, 0xb9, 0x82, 0x09, 0x7e, 0xb2, 0x6a, 0x17, 0x7b, 0x37, 0xde, 0xdd, 0x59, 0xee, 0x00, 0x49,
	0x20, 0x10, 0x4c, 0xa9, 0x1f, 0x2e, 0x8b, 0xf8, 0xc4, 0x7e, 0x38, 0x60, 0xba, 0x84, 0xd9, 0xd1,
	0x44, 0x09, 0x5e, 0x2b, 0x24, 0x11, 0xf8, 0x4a, 0x33, 0x6d, 0x94, 0xb5, 0x19, 0xe7, 0x0d, 0xb2,
	0xee, 0x52, 0x72, 0xd9, 0x98, 0x38, 0x40, 0x97, 0x70, 0xf1, 0xc2, 0xcb, 0xaa, 0x1e, 0x9e, 0x61,
	0x05, 0xd3, 0xc6, 0x61, 0x48, 0x80, 0x3d, 0xab, 0xf9, 0x16, 0xeb, 0x78, 0xec, 0x58, 0x0b, 0xe8,
	0x2d, 0x84, 0xaf, 0x6c, 0x57, 0x15, 0x4c, 0x63, 0x27, 0x99, 0x13, 0x7a, 0x5d, 0xe1, 0x1b, 0xcc,
	0x8e, 0xc2, 0x41, 0x01, 0x22, 0xf0, 0x8d, 0x42, 0xf9, 0x5c, 0xd8, 0x04, 0xe3, 0xbc, 0x41, 0xf3,
	0x5f, 0x0f, 0xce, 0x1f, 0x8c, 0xde, 0xac, 0x50, 0x7e, 0x57, 0x1f, 0x48, 0x16, 0x10, 0xb4, 0x5d,
	0x93, 0xeb, 0xd4, 0xce, 0xb3, 0x37, 0xc0, 0x24, 0xea, 0xd3, 0x2e, 0x10, 0x1d, 0x91, 0x39, 0x4c,
	0x6c, 0x49, 0x84, 0x38, 0x49, 0xb7, 0xf3, 0xe4, 0xf2, 0x1f, 0x77, 0xf8, 0x67, 0x01, 0x41, 0x7b,
	0xb5, 0xf6, 0xc0, 0x5e, 0x27, 0xed, 0x81, 0xfd, 0x06, 0xe8, 0xe8, 0x31, 0x7c, 0x9f, 0xa6, 0x59,
	0x67, 0x01, 0xd7, 0xbe, 0x5d, 0xbe, 0xfb, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x55, 0x10, 0x8d,
	0xc7, 0x96, 0x02, 0x00, 0x00,
}
