// Code generated by protoc-gen-go. DO NOT EDIT.
// source: web_proxy.proto

package rpc

import (
	context "context"
	fmt "fmt"
	rpc "github.com/celer-network/goCeler/rpc"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
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

type SessionToken struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionToken) Reset()         { *m = SessionToken{} }
func (m *SessionToken) String() string { return proto.CompactTextString(m) }
func (*SessionToken) ProtoMessage()    {}
func (*SessionToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_988d04d3a7affae2, []int{0}
}

func (m *SessionToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionToken.Unmarshal(m, b)
}
func (m *SessionToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionToken.Marshal(b, m, deterministic)
}
func (m *SessionToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionToken.Merge(m, src)
}
func (m *SessionToken) XXX_Size() int {
	return xxx_messageInfo_SessionToken.Size(m)
}
func (m *SessionToken) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionToken.DiscardUnknown(m)
}

var xxx_messageInfo_SessionToken proto.InternalMessageInfo

func (m *SessionToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*SessionToken)(nil), "webproxyrpc.SessionToken")
}

func init() { proto.RegisterFile("web_proxy.proto", fileDescriptor_988d04d3a7affae2) }

var fileDescriptor_988d04d3a7affae2 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0x51, 0x4b, 0xfb, 0x30,
	0x14, 0xc5, 0xbb, 0x3f, 0xfc, 0x05, 0xd3, 0x0d, 0x31, 0x88, 0xce, 0xf8, 0x22, 0xc5, 0x07, 0x5f,
	0x96, 0x88, 0x8a, 0xef, 0x6e, 0xfa, 0x38, 0x94, 0x4e, 0x10, 0x7c, 0x91, 0x26, 0x5e, 0xd3, 0xb2,
	0x36, 0xc9, 0x92, 0x94, 0xea, 0x07, 0xf0, 0x7b, 0x4b, 0xd3, 0x0e, 0x2a, 0xe2, 0x5b, 0x4e, 0xee,
	0x39, 0xf7, 0x77, 0x0f, 0xda, 0x6b, 0x80, 0xbf, 0x1a, 0xab, 0x3f, 0x3e, 0xa9, 0xb1, 0xda, 0x6b,
	0x1c, 0x37, 0xc0, 0x83, 0xb6, 0x46, 0x90, 0x13, 0xa9, 0xb5, 0x2c, 0x81, 0x85, 0x11, 0xaf, 0xdf,
	0x19, 0x54, 0xc6, 0xf7, 0x4e, 0x32, 0xa9, 0xc0, 0xb9, 0x4c, 0x42, 0x27, 0x93, 0x33, 0x34, 0x5e,
	0x81, 0x73, 0x85, 0x56, 0x4f, 0x7a, 0x0d, 0x0a, 0x1f, 0xa0, 0xff, 0xbe, 0x7d, 0x4c, 0x47, 0xa7,
	0xa3, 0xf3, 0xdd, 0xb4, 0x13, 0x97, 0x5f, 0xff, 0x50, 0xfc, 0x0c, 0xfc, 0xb1, 0x25, 0xa4, 0x46,
	0xe0, 0x3b, 0x34, 0x59, 0x58, 0xc8, 0x3c, 0xf4, 0x59, 0x7c, 0x48, 0x3b, 0x26, 0xdd, 0x32, 0xe9,
	0x7d, 0xcb, 0x24, 0xc7, 0x74, 0x70, 0x18, 0x1d, 0x92, 0x92, 0x08, 0xcf, 0x51, 0xfc, 0x60, 0x40,
	0x2d, 0xf2, 0x4c, 0x29, 0x28, 0xf1, 0x11, 0x6d, 0x3d, 0x83, 0x9f, 0x14, 0x36, 0x35, 0x38, 0x4f,
	0xa6, 0xbf, 0x07, 0xce, 0x68, 0xe5, 0x20, 0x89, 0xf0, 0x35, 0xda, 0x5f, 0xd5, 0xdc, 0x09, 0x5b,
	0x70, 0x58, 0x76, 0xcd, 0x1c, 0x1e, 0x87, 0xc0, 0x6d, 0xed, 0xf3, 0x14, 0x36, 0x64, 0x12, 0xd4,
	0x02, 0x4a, 0xb0, 0x4b, 0x27, 0x93, 0xe8, 0x62, 0x84, 0x6f, 0x50, 0xbc, 0x02, 0xf5, 0xd6, 0x07,
	0xf0, 0x4f, 0x07, 0xf9, 0xa3, 0x4c, 0x12, 0xcd, 0xd9, 0xcb, 0x4c, 0x16, 0x3e, 0xaf, 0x39, 0x15,
	0xba, 0x62, 0xa2, 0x0d, 0xcc, 0x14, 0xf8, 0x46, 0xdb, 0x35, 0x93, 0x3a, 0x2c, 0x60, 0xdb, 0xc2,
	0xcc, 0x1a, 0xc1, 0x77, 0xc2, 0x8a, 0xab, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x52, 0xd9, 0x42,
	0x23, 0xb1, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WebProxyRpcClient is the client API for WebProxyRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WebProxyRpcClient interface {
	CreateSession(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*SessionToken, error)
	OpenChannel(ctx context.Context, in *rpc.OpenChannelRequest, opts ...grpc.CallOption) (*rpc.OpenChannelResponse, error)
	SubscribeMessages(ctx context.Context, in *rpc.AuthReq, opts ...grpc.CallOption) (WebProxyRpc_SubscribeMessagesClient, error)
	SendMessage(ctx context.Context, in *rpc.CelerMsg, opts ...grpc.CallOption) (*empty.Empty, error)
}

type webProxyRpcClient struct {
	cc *grpc.ClientConn
}

func NewWebProxyRpcClient(cc *grpc.ClientConn) WebProxyRpcClient {
	return &webProxyRpcClient{cc}
}

func (c *webProxyRpcClient) CreateSession(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*SessionToken, error) {
	out := new(SessionToken)
	err := c.cc.Invoke(ctx, "/webproxyrpc.WebProxyRpc/CreateSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webProxyRpcClient) OpenChannel(ctx context.Context, in *rpc.OpenChannelRequest, opts ...grpc.CallOption) (*rpc.OpenChannelResponse, error) {
	out := new(rpc.OpenChannelResponse)
	err := c.cc.Invoke(ctx, "/webproxyrpc.WebProxyRpc/OpenChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webProxyRpcClient) SubscribeMessages(ctx context.Context, in *rpc.AuthReq, opts ...grpc.CallOption) (WebProxyRpc_SubscribeMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_WebProxyRpc_serviceDesc.Streams[0], "/webproxyrpc.WebProxyRpc/SubscribeMessages", opts...)
	if err != nil {
		return nil, err
	}
	x := &webProxyRpcSubscribeMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WebProxyRpc_SubscribeMessagesClient interface {
	Recv() (*rpc.CelerMsg, error)
	grpc.ClientStream
}

type webProxyRpcSubscribeMessagesClient struct {
	grpc.ClientStream
}

func (x *webProxyRpcSubscribeMessagesClient) Recv() (*rpc.CelerMsg, error) {
	m := new(rpc.CelerMsg)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *webProxyRpcClient) SendMessage(ctx context.Context, in *rpc.CelerMsg, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/webproxyrpc.WebProxyRpc/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebProxyRpcServer is the server API for WebProxyRpc service.
type WebProxyRpcServer interface {
	CreateSession(context.Context, *empty.Empty) (*SessionToken, error)
	OpenChannel(context.Context, *rpc.OpenChannelRequest) (*rpc.OpenChannelResponse, error)
	SubscribeMessages(*rpc.AuthReq, WebProxyRpc_SubscribeMessagesServer) error
	SendMessage(context.Context, *rpc.CelerMsg) (*empty.Empty, error)
}

func RegisterWebProxyRpcServer(s *grpc.Server, srv WebProxyRpcServer) {
	s.RegisterService(&_WebProxyRpc_serviceDesc, srv)
}

func _WebProxyRpc_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebProxyRpcServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/webproxyrpc.WebProxyRpc/CreateSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebProxyRpcServer).CreateSession(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebProxyRpc_OpenChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpc.OpenChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebProxyRpcServer).OpenChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/webproxyrpc.WebProxyRpc/OpenChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebProxyRpcServer).OpenChannel(ctx, req.(*rpc.OpenChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebProxyRpc_SubscribeMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(rpc.AuthReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WebProxyRpcServer).SubscribeMessages(m, &webProxyRpcSubscribeMessagesServer{stream})
}

type WebProxyRpc_SubscribeMessagesServer interface {
	Send(*rpc.CelerMsg) error
	grpc.ServerStream
}

type webProxyRpcSubscribeMessagesServer struct {
	grpc.ServerStream
}

func (x *webProxyRpcSubscribeMessagesServer) Send(m *rpc.CelerMsg) error {
	return x.ServerStream.SendMsg(m)
}

func _WebProxyRpc_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpc.CelerMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebProxyRpcServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/webproxyrpc.WebProxyRpc/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebProxyRpcServer).SendMessage(ctx, req.(*rpc.CelerMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _WebProxyRpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "webproxyrpc.WebProxyRpc",
	HandlerType: (*WebProxyRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSession",
			Handler:    _WebProxyRpc_CreateSession_Handler,
		},
		{
			MethodName: "OpenChannel",
			Handler:    _WebProxyRpc_OpenChannel_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _WebProxyRpc_SendMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeMessages",
			Handler:       _WebProxyRpc_SubscribeMessages_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "web_proxy.proto",
}
