// Code generated by protoc-gen-go. DO NOT EDIT.
// source: member_data.proto

package member

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type UpdateSessionKeyByUnionIdRequest struct {
	UnionId              string   `protobuf:"bytes,1,opt,name=unionId,proto3" json:"unionId,omitempty"`
	SessionKey           string   `protobuf:"bytes,2,opt,name=sessionKey,proto3" json:"sessionKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateSessionKeyByUnionIdRequest) Reset()         { *m = UpdateSessionKeyByUnionIdRequest{} }
func (m *UpdateSessionKeyByUnionIdRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateSessionKeyByUnionIdRequest) ProtoMessage()    {}
func (*UpdateSessionKeyByUnionIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4010a41aeb367e84, []int{0}
}

func (m *UpdateSessionKeyByUnionIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateSessionKeyByUnionIdRequest.Unmarshal(m, b)
}
func (m *UpdateSessionKeyByUnionIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateSessionKeyByUnionIdRequest.Marshal(b, m, deterministic)
}
func (m *UpdateSessionKeyByUnionIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateSessionKeyByUnionIdRequest.Merge(m, src)
}
func (m *UpdateSessionKeyByUnionIdRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateSessionKeyByUnionIdRequest.Size(m)
}
func (m *UpdateSessionKeyByUnionIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateSessionKeyByUnionIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateSessionKeyByUnionIdRequest proto.InternalMessageInfo

func (m *UpdateSessionKeyByUnionIdRequest) GetUnionId() string {
	if m != nil {
		return m.UnionId
	}
	return ""
}

func (m *UpdateSessionKeyByUnionIdRequest) GetSessionKey() string {
	if m != nil {
		return m.SessionKey
	}
	return ""
}

type UpdateSessionKeyByUnionIdResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateSessionKeyByUnionIdResponse) Reset()         { *m = UpdateSessionKeyByUnionIdResponse{} }
func (m *UpdateSessionKeyByUnionIdResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateSessionKeyByUnionIdResponse) ProtoMessage()    {}
func (*UpdateSessionKeyByUnionIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4010a41aeb367e84, []int{1}
}

func (m *UpdateSessionKeyByUnionIdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateSessionKeyByUnionIdResponse.Unmarshal(m, b)
}
func (m *UpdateSessionKeyByUnionIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateSessionKeyByUnionIdResponse.Marshal(b, m, deterministic)
}
func (m *UpdateSessionKeyByUnionIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateSessionKeyByUnionIdResponse.Merge(m, src)
}
func (m *UpdateSessionKeyByUnionIdResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateSessionKeyByUnionIdResponse.Size(m)
}
func (m *UpdateSessionKeyByUnionIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateSessionKeyByUnionIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateSessionKeyByUnionIdResponse proto.InternalMessageInfo

type VerifyUserNumberWithPwdRequest struct {
	UserNumber           int64    `protobuf:"varint,1,opt,name=userNumber,proto3" json:"userNumber,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyUserNumberWithPwdRequest) Reset()         { *m = VerifyUserNumberWithPwdRequest{} }
func (m *VerifyUserNumberWithPwdRequest) String() string { return proto.CompactTextString(m) }
func (*VerifyUserNumberWithPwdRequest) ProtoMessage()    {}
func (*VerifyUserNumberWithPwdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4010a41aeb367e84, []int{2}
}

func (m *VerifyUserNumberWithPwdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyUserNumberWithPwdRequest.Unmarshal(m, b)
}
func (m *VerifyUserNumberWithPwdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyUserNumberWithPwdRequest.Marshal(b, m, deterministic)
}
func (m *VerifyUserNumberWithPwdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyUserNumberWithPwdRequest.Merge(m, src)
}
func (m *VerifyUserNumberWithPwdRequest) XXX_Size() int {
	return xxx_messageInfo_VerifyUserNumberWithPwdRequest.Size(m)
}
func (m *VerifyUserNumberWithPwdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyUserNumberWithPwdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyUserNumberWithPwdRequest proto.InternalMessageInfo

func (m *VerifyUserNumberWithPwdRequest) GetUserNumber() int64 {
	if m != nil {
		return m.UserNumber
	}
	return 0
}

func (m *VerifyUserNumberWithPwdRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type VerifyUserNumberWithPwdResponse struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyUserNumberWithPwdResponse) Reset()         { *m = VerifyUserNumberWithPwdResponse{} }
func (m *VerifyUserNumberWithPwdResponse) String() string { return proto.CompactTextString(m) }
func (*VerifyUserNumberWithPwdResponse) ProtoMessage()    {}
func (*VerifyUserNumberWithPwdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4010a41aeb367e84, []int{3}
}

func (m *VerifyUserNumberWithPwdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyUserNumberWithPwdResponse.Unmarshal(m, b)
}
func (m *VerifyUserNumberWithPwdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyUserNumberWithPwdResponse.Marshal(b, m, deterministic)
}
func (m *VerifyUserNumberWithPwdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyUserNumberWithPwdResponse.Merge(m, src)
}
func (m *VerifyUserNumberWithPwdResponse) XXX_Size() int {
	return xxx_messageInfo_VerifyUserNumberWithPwdResponse.Size(m)
}
func (m *VerifyUserNumberWithPwdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyUserNumberWithPwdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyUserNumberWithPwdResponse proto.InternalMessageInfo

func (m *VerifyUserNumberWithPwdResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type CreateUserRequest struct {
	UserNumber           int64    `protobuf:"varint,1,opt,name=userNumber,proto3" json:"userNumber,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Mobile               string   `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	UserName             string   `protobuf:"bytes,4,opt,name=userName,proto3" json:"userName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4010a41aeb367e84, []int{4}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetUserNumber() int64 {
	if m != nil {
		return m.UserNumber
	}
	return 0
}

func (m *CreateUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CreateUserRequest) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *CreateUserRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type CreateUserResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4010a41aeb367e84, []int{5}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*UpdateSessionKeyByUnionIdRequest)(nil), "member.updateSessionKeyByUnionIdRequest")
	proto.RegisterType((*UpdateSessionKeyByUnionIdResponse)(nil), "member.updateSessionKeyByUnionIdResponse")
	proto.RegisterType((*VerifyUserNumberWithPwdRequest)(nil), "member.verifyUserNumberWithPwdRequest")
	proto.RegisterType((*VerifyUserNumberWithPwdResponse)(nil), "member.verifyUserNumberWithPwdResponse")
	proto.RegisterType((*CreateUserRequest)(nil), "member.createUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "member.createUserResponse")
}

func init() { proto.RegisterFile("member_data.proto", fileDescriptor_4010a41aeb367e84) }

var fileDescriptor_4010a41aeb367e84 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0x69, 0xfb, 0xbe, 0x51, 0xe7, 0xd6, 0x45, 0x34, 0xcd, 0x21, 0xd6, 0x15, 0xb4, 0x5e,
	0x72, 0xd0, 0x93, 0x57, 0x3d, 0x89, 0x20, 0x12, 0x29, 0x5e, 0x0a, 0xb2, 0x31, 0x23, 0x46, 0x4c,
	0x36, 0xee, 0x1f, 0x4b, 0xce, 0x7e, 0x1c, 0xbf, 0xa4, 0x64, 0x77, 0x13, 0x03, 0x9a, 0xd6, 0x83,
	0xc7, 0x67, 0x66, 0x9e, 0x67, 0x7f, 0x99, 0x09, 0x8c, 0x73, 0xcc, 0x13, 0x14, 0xf7, 0x29, 0x53,
	0x2c, 0x2a, 0x05, 0x57, 0x9c, 0x78, 0xb6, 0x44, 0x17, 0x30, 0xd5, 0x65, 0xca, 0x14, 0xde, 0xa2,
	0x94, 0x19, 0x2f, 0xae, 0xb0, 0x3a, 0xaf, 0xe6, 0x45, 0xc6, 0x8b, 0xcb, 0x34, 0xc6, 0x57, 0x8d,
	0x52, 0x11, 0x1f, 0x36, 0xb4, 0xad, 0xf8, 0x83, 0xe9, 0x60, 0xb6, 0x15, 0x37, 0x92, 0x84, 0x00,
	0xb2, 0xf5, 0xf9, 0x43, 0xd3, 0xec, 0x54, 0xe8, 0x01, 0xec, 0xaf, 0x48, 0x97, 0x25, 0x2f, 0x24,
	0xd2, 0x05, 0x84, 0x6f, 0x28, 0xb2, 0xc7, 0x6a, 0x2e, 0x51, 0x5c, 0xeb, 0x1a, 0xeb, 0x2e, 0x53,
	0x4f, 0x37, 0xcb, 0x16, 0x20, 0x04, 0xd0, 0x6d, 0xcf, 0x30, 0x8c, 0xe2, 0x4e, 0x85, 0x04, 0xb0,
	0x59, 0x32, 0x29, 0x97, 0x5c, 0xa4, 0x0e, 0xa2, 0xd5, 0xf4, 0x0c, 0xf6, 0x7a, 0xd3, 0x2d, 0x00,
	0xd9, 0x01, 0x4f, 0x2a, 0xa6, 0xb4, 0x34, 0xd1, 0xff, 0x63, 0xa7, 0xe8, 0xfb, 0x00, 0xc6, 0x0f,
	0x02, 0x99, 0xc2, 0xda, 0xfb, 0x07, 0x30, 0xf5, 0x4b, 0x39, 0x4f, 0xb2, 0x17, 0xf4, 0x47, 0xa6,
	0xe3, 0x54, 0xed, 0x31, 0x09, 0x2c, 0x47, 0xff, 0x9f, 0xf5, 0x34, 0x9a, 0x6e, 0x03, 0xe9, 0x42,
	0x58, 0xe6, 0x93, 0x8f, 0x21, 0xb8, 0x13, 0x92, 0x12, 0x26, 0xbd, 0x4b, 0x26, 0xb3, 0xc8, 0x4e,
	0x45, 0xeb, 0xae, 0x1c, 0x1c, 0xff, 0x62, 0xd2, 0x2d, 0xec, 0x19, 0x76, 0x7b, 0x76, 0x4a, 0x0e,
	0x9b, 0x94, 0xd5, 0x27, 0x0d, 0x8e, 0xd6, 0xce, 0xb9, 0xb7, 0x2e, 0x00, 0xbe, 0x3e, 0x9f, 0x4c,
	0x1a, 0xdb, 0xb7, 0xbb, 0x04, 0xc1, 0x4f, 0x2d, 0x1b, 0x92, 0x78, 0xe6, 0xa7, 0x3f, 0xfd, 0x0c,
	0x00, 0x00, 0xff, 0xff, 0xf9, 0x63, 0xdd, 0xde, 0x09, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MemberClient is the client API for Member service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MemberClient interface {
	UpdateSessionKeyByUnionId(ctx context.Context, in *UpdateSessionKeyByUnionIdRequest, opts ...grpc.CallOption) (*UpdateSessionKeyByUnionIdResponse, error)
	VerifyUserNumberWithPwd(ctx context.Context, in *VerifyUserNumberWithPwdRequest, opts ...grpc.CallOption) (*VerifyUserNumberWithPwdResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
}

type memberClient struct {
	cc *grpc.ClientConn
}

func NewMemberClient(cc *grpc.ClientConn) MemberClient {
	return &memberClient{cc}
}

func (c *memberClient) UpdateSessionKeyByUnionId(ctx context.Context, in *UpdateSessionKeyByUnionIdRequest, opts ...grpc.CallOption) (*UpdateSessionKeyByUnionIdResponse, error) {
	out := new(UpdateSessionKeyByUnionIdResponse)
	err := c.cc.Invoke(ctx, "/member.member/updateSessionKeyByUnionId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberClient) VerifyUserNumberWithPwd(ctx context.Context, in *VerifyUserNumberWithPwdRequest, opts ...grpc.CallOption) (*VerifyUserNumberWithPwdResponse, error) {
	out := new(VerifyUserNumberWithPwdResponse)
	err := c.cc.Invoke(ctx, "/member.member/verifyUserNumberWithPwd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/member.member/createUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MemberServer is the server API for Member service.
type MemberServer interface {
	UpdateSessionKeyByUnionId(context.Context, *UpdateSessionKeyByUnionIdRequest) (*UpdateSessionKeyByUnionIdResponse, error)
	VerifyUserNumberWithPwd(context.Context, *VerifyUserNumberWithPwdRequest) (*VerifyUserNumberWithPwdResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
}

// UnimplementedMemberServer can be embedded to have forward compatible implementations.
type UnimplementedMemberServer struct {
}

func (*UnimplementedMemberServer) UpdateSessionKeyByUnionId(ctx context.Context, req *UpdateSessionKeyByUnionIdRequest) (*UpdateSessionKeyByUnionIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSessionKeyByUnionId not implemented")
}
func (*UnimplementedMemberServer) VerifyUserNumberWithPwd(ctx context.Context, req *VerifyUserNumberWithPwdRequest) (*VerifyUserNumberWithPwdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyUserNumberWithPwd not implemented")
}
func (*UnimplementedMemberServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

func RegisterMemberServer(s *grpc.Server, srv MemberServer) {
	s.RegisterService(&_Member_serviceDesc, srv)
}

func _Member_UpdateSessionKeyByUnionId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSessionKeyByUnionIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServer).UpdateSessionKeyByUnionId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/member.member/UpdateSessionKeyByUnionId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServer).UpdateSessionKeyByUnionId(ctx, req.(*UpdateSessionKeyByUnionIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Member_VerifyUserNumberWithPwd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyUserNumberWithPwdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServer).VerifyUserNumberWithPwd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/member.member/VerifyUserNumberWithPwd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServer).VerifyUserNumberWithPwd(ctx, req.(*VerifyUserNumberWithPwdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Member_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/member.member/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Member_serviceDesc = grpc.ServiceDesc{
	ServiceName: "member.member",
	HandlerType: (*MemberServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "updateSessionKeyByUnionId",
			Handler:    _Member_UpdateSessionKeyByUnionId_Handler,
		},
		{
			MethodName: "verifyUserNumberWithPwd",
			Handler:    _Member_VerifyUserNumberWithPwd_Handler,
		},
		{
			MethodName: "createUser",
			Handler:    _Member_CreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "member_data.proto",
}
