// Code generated by protoc-gen-go.
// source: auth.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	auth.proto
	bgsave.proto
	chat.proto
	game.proto
	geoip.proto
	rankserver.proto
	snowflake.proto
	wordfilter.proto

It has these top-level messages:
	User
*/
package proto

import proto1 "github.com/golang/protobuf/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal

type User struct {
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto1.CompactTextString(m) }
func (*User) ProtoMessage()    {}

type User_Nil struct {
}

func (m *User_Nil) Reset()         { *m = User_Nil{} }
func (m *User_Nil) String() string { return proto1.CompactTextString(m) }
func (*User_Nil) ProtoMessage()    {}

type User_LoginInfo struct {
	Uuid     string `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	AuthType int32  `protobuf:"varint,2,opt,name=auth_type" json:"auth_type,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=username" json:"username,omitempty"`
	Passwd   string `protobuf:"bytes,4,opt,name=passwd" json:"passwd,omitempty"`
	Gsid     int32  `protobuf:"varint,5,opt,name=gsid" json:"gsid,omitempty"`
}

func (m *User_LoginInfo) Reset()         { *m = User_LoginInfo{} }
func (m *User_LoginInfo) String() string { return proto1.CompactTextString(m) }
func (*User_LoginInfo) ProtoMessage()    {}

type User_LoginResp struct {
	Uid      int32  `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	UniqueId uint64 `protobuf:"varint,2,opt,name=unique_id" json:"unique_id,omitempty"`
	Cert     string `protobuf:"bytes,3,opt,name=cert" json:"cert,omitempty"`
}

func (m *User_LoginResp) Reset()         { *m = User_LoginResp{} }
func (m *User_LoginResp) String() string { return proto1.CompactTextString(m) }
func (*User_LoginResp) ProtoMessage()    {}

func init() {
}

// Client API for AuthService service

type AuthServiceClient interface {
	Login(ctx context.Context, in *User_LoginInfo, opts ...grpc.CallOption) (*User_LoginResp, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Login(ctx context.Context, in *User_LoginInfo, opts ...grpc.CallOption) (*User_LoginResp, error) {
	out := new(User_LoginResp)
	err := grpc.Invoke(ctx, "/proto.AuthService/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthService service

type AuthServiceServer interface {
	Login(context.Context, *User_LoginInfo) (*User_LoginResp, error)
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(User_LoginInfo)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(AuthServiceServer).Login(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
