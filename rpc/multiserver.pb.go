// Code generated by protoc-gen-go. DO NOT EDIT.
// source: multiserver.proto

package rpc

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

type FwdReq struct {
	Dest                 string    `protobuf:"bytes,1,opt,name=dest,proto3" json:"dest,omitempty"`
	Message              *CelerMsg `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *FwdReq) Reset()         { *m = FwdReq{} }
func (m *FwdReq) String() string { return proto.CompactTextString(m) }
func (*FwdReq) ProtoMessage()    {}
func (*FwdReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9732a59a30a1ae34, []int{0}
}

func (m *FwdReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FwdReq.Unmarshal(m, b)
}
func (m *FwdReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FwdReq.Marshal(b, m, deterministic)
}
func (m *FwdReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FwdReq.Merge(m, src)
}
func (m *FwdReq) XXX_Size() int {
	return xxx_messageInfo_FwdReq.Size(m)
}
func (m *FwdReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FwdReq.DiscardUnknown(m)
}

var xxx_messageInfo_FwdReq proto.InternalMessageInfo

func (m *FwdReq) GetDest() string {
	if m != nil {
		return m.Dest
	}
	return ""
}

func (m *FwdReq) GetMessage() *CelerMsg {
	if m != nil {
		return m.Message
	}
	return nil
}

type FwdReply struct {
	Accepted             bool     `protobuf:"varint,1,opt,name=accepted,proto3" json:"accepted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FwdReply) Reset()         { *m = FwdReply{} }
func (m *FwdReply) String() string { return proto.CompactTextString(m) }
func (*FwdReply) ProtoMessage()    {}
func (*FwdReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_9732a59a30a1ae34, []int{1}
}

func (m *FwdReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FwdReply.Unmarshal(m, b)
}
func (m *FwdReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FwdReply.Marshal(b, m, deterministic)
}
func (m *FwdReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FwdReply.Merge(m, src)
}
func (m *FwdReply) XXX_Size() int {
	return xxx_messageInfo_FwdReply.Size(m)
}
func (m *FwdReply) XXX_DiscardUnknown() {
	xxx_messageInfo_FwdReply.DiscardUnknown(m)
}

var xxx_messageInfo_FwdReply proto.InternalMessageInfo

func (m *FwdReply) GetAccepted() bool {
	if m != nil {
		return m.Accepted
	}
	return false
}

type PingReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingReq) Reset()         { *m = PingReq{} }
func (m *PingReq) String() string { return proto.CompactTextString(m) }
func (*PingReq) ProtoMessage()    {}
func (*PingReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9732a59a30a1ae34, []int{2}
}

func (m *PingReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingReq.Unmarshal(m, b)
}
func (m *PingReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingReq.Marshal(b, m, deterministic)
}
func (m *PingReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingReq.Merge(m, src)
}
func (m *PingReq) XXX_Size() int {
	return xxx_messageInfo_PingReq.Size(m)
}
func (m *PingReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PingReq.DiscardUnknown(m)
}

var xxx_messageInfo_PingReq proto.InternalMessageInfo

type PingReply struct {
	Numclients           uint32   `protobuf:"varint,1,opt,name=numclients,proto3" json:"numclients,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingReply) Reset()         { *m = PingReply{} }
func (m *PingReply) String() string { return proto.CompactTextString(m) }
func (*PingReply) ProtoMessage()    {}
func (*PingReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_9732a59a30a1ae34, []int{3}
}

func (m *PingReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingReply.Unmarshal(m, b)
}
func (m *PingReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingReply.Marshal(b, m, deterministic)
}
func (m *PingReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingReply.Merge(m, src)
}
func (m *PingReply) XXX_Size() int {
	return xxx_messageInfo_PingReply.Size(m)
}
func (m *PingReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PingReply.DiscardUnknown(m)
}

var xxx_messageInfo_PingReply proto.InternalMessageInfo

func (m *PingReply) GetNumclients() uint32 {
	if m != nil {
		return m.Numclients
	}
	return 0
}

type PickReq struct {
	Client               string   `protobuf:"bytes,1,opt,name=client,proto3" json:"client,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PickReq) Reset()         { *m = PickReq{} }
func (m *PickReq) String() string { return proto.CompactTextString(m) }
func (*PickReq) ProtoMessage()    {}
func (*PickReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9732a59a30a1ae34, []int{4}
}

func (m *PickReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PickReq.Unmarshal(m, b)
}
func (m *PickReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PickReq.Marshal(b, m, deterministic)
}
func (m *PickReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PickReq.Merge(m, src)
}
func (m *PickReq) XXX_Size() int {
	return xxx_messageInfo_PickReq.Size(m)
}
func (m *PickReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PickReq.DiscardUnknown(m)
}

var xxx_messageInfo_PickReq proto.InternalMessageInfo

func (m *PickReq) GetClient() string {
	if m != nil {
		return m.Client
	}
	return ""
}

type PickReply struct {
	Server               string   `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PickReply) Reset()         { *m = PickReply{} }
func (m *PickReply) String() string { return proto.CompactTextString(m) }
func (*PickReply) ProtoMessage()    {}
func (*PickReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_9732a59a30a1ae34, []int{5}
}

func (m *PickReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PickReply.Unmarshal(m, b)
}
func (m *PickReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PickReply.Marshal(b, m, deterministic)
}
func (m *PickReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PickReply.Merge(m, src)
}
func (m *PickReply) XXX_Size() int {
	return xxx_messageInfo_PickReply.Size(m)
}
func (m *PickReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PickReply.DiscardUnknown(m)
}

var xxx_messageInfo_PickReply proto.InternalMessageInfo

func (m *PickReply) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func init() {
	proto.RegisterType((*FwdReq)(nil), "rpc.FwdReq")
	proto.RegisterType((*FwdReply)(nil), "rpc.FwdReply")
	proto.RegisterType((*PingReq)(nil), "rpc.PingReq")
	proto.RegisterType((*PingReply)(nil), "rpc.PingReply")
	proto.RegisterType((*PickReq)(nil), "rpc.PickReq")
	proto.RegisterType((*PickReply)(nil), "rpc.PickReply")
}

func init() { proto.RegisterFile("multiserver.proto", fileDescriptor_9732a59a30a1ae34) }

var fileDescriptor_9732a59a30a1ae34 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x51, 0x4d, 0x4f, 0x02, 0x31,
	0x14, 0x64, 0x23, 0x2e, 0xf0, 0x70, 0x4d, 0xec, 0x81, 0x90, 0x3d, 0x18, 0xac, 0x66, 0x25, 0x51,
	0x97, 0x04, 0xff, 0x81, 0x46, 0x6f, 0x24, 0xa6, 0xde, 0xbc, 0x41, 0x79, 0xa9, 0x1b, 0xf6, 0xa3,
	0xb4, 0x45, 0xc2, 0x6f, 0xf0, 0x4f, 0x9b, 0x7e, 0x2c, 0x72, 0xeb, 0xbc, 0x37, 0x6f, 0x32, 0x33,
	0x85, 0xab, 0x6a, 0x57, 0x9a, 0x42, 0xa3, 0xfa, 0x41, 0x95, 0x4b, 0xd5, 0x98, 0x86, 0x9c, 0x29,
	0xc9, 0xd3, 0xa4, 0x42, 0xad, 0x97, 0x02, 0xfd, 0x8c, 0xbe, 0x41, 0xfc, 0xbe, 0x5f, 0x33, 0xdc,
	0x12, 0x02, 0xdd, 0x35, 0x6a, 0x33, 0x8e, 0x26, 0xd1, 0x74, 0xc0, 0xdc, 0x9b, 0xdc, 0x43, 0x2f,
	0xd0, 0xc7, 0xe7, 0x93, 0x68, 0x3a, 0x9c, 0x27, 0xb9, 0x92, 0x3c, 0x7f, 0xc5, 0x12, 0xd5, 0x42,
	0x0b, 0xd6, 0x6e, 0x69, 0x06, 0x7d, 0x27, 0x23, 0xcb, 0x03, 0x49, 0xa1, 0xbf, 0xe4, 0x1c, 0xa5,
	0xc1, 0xb5, 0x13, 0xeb, 0xb3, 0x23, 0xa6, 0x03, 0xe8, 0x7d, 0x14, 0xb5, 0x60, 0xb8, 0xa5, 0x0f,
	0x30, 0xf0, 0x4f, 0x7b, 0x73, 0x0d, 0x50, 0xef, 0x2a, 0x5e, 0x16, 0x58, 0x1b, 0xed, 0xae, 0x12,
	0x76, 0x32, 0xa1, 0x37, 0xf6, 0x8e, 0x6f, 0xac, 0xcf, 0x11, 0xc4, 0x7e, 0x1a, 0x9c, 0x06, 0x44,
	0x6f, 0xad, 0x9e, 0xa5, 0x58, 0xbd, 0x11, 0xc4, 0x3e, 0x7a, 0x4b, 0xf2, 0x68, 0xfe, 0x1b, 0xc1,
	0x70, 0x61, 0x8b, 0xf9, 0x74, 0x98, 0x64, 0x2e, 0xfe, 0x42, 0x0b, 0x32, 0x74, 0xc9, 0x7c, 0x17,
	0x69, 0xf2, 0x0f, 0x64, 0x79, 0xa0, 0x1d, 0x92, 0x41, 0xd7, 0x9a, 0x25, 0x17, 0x6e, 0x11, 0x22,
	0xa4, 0x97, 0x27, 0xc8, 0xf3, 0x1e, 0x01, 0xac, 0x89, 0xa0, 0xde, 0xb2, 0x9d, 0xf1, 0x23, 0x3b,
	0x78, 0xa4, 0x9d, 0x97, 0xec, 0xeb, 0x4e, 0x14, 0xe6, 0x7b, 0xb7, 0xca, 0x79, 0x53, 0xcd, 0xb8,
	0x6d, 0xf5, 0xa9, 0x46, 0xb3, 0x6f, 0xd4, 0x66, 0x26, 0x1a, 0xd7, 0xf2, 0x4c, 0x49, 0xbe, 0x8a,
	0xdd, 0x5f, 0x3d, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x76, 0x1e, 0x9a, 0x81, 0xd4, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MultiServerClient is the client API for MultiServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MultiServerClient interface {
	FwdMsg(ctx context.Context, in *FwdReq, opts ...grpc.CallOption) (*FwdReply, error)
	Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingReply, error)
	PickServer(ctx context.Context, in *PickReq, opts ...grpc.CallOption) (*PickReply, error)
}

type multiServerClient struct {
	cc *grpc.ClientConn
}

func NewMultiServerClient(cc *grpc.ClientConn) MultiServerClient {
	return &multiServerClient{cc}
}

func (c *multiServerClient) FwdMsg(ctx context.Context, in *FwdReq, opts ...grpc.CallOption) (*FwdReply, error) {
	out := new(FwdReply)
	err := c.cc.Invoke(ctx, "/rpc.MultiServer/FwdMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *multiServerClient) Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/rpc.MultiServer/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *multiServerClient) PickServer(ctx context.Context, in *PickReq, opts ...grpc.CallOption) (*PickReply, error) {
	out := new(PickReply)
	err := c.cc.Invoke(ctx, "/rpc.MultiServer/PickServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MultiServerServer is the server API for MultiServer service.
type MultiServerServer interface {
	FwdMsg(context.Context, *FwdReq) (*FwdReply, error)
	Ping(context.Context, *PingReq) (*PingReply, error)
	PickServer(context.Context, *PickReq) (*PickReply, error)
}

// UnimplementedMultiServerServer can be embedded to have forward compatible implementations.
type UnimplementedMultiServerServer struct {
}

func (*UnimplementedMultiServerServer) FwdMsg(ctx context.Context, req *FwdReq) (*FwdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FwdMsg not implemented")
}
func (*UnimplementedMultiServerServer) Ping(ctx context.Context, req *PingReq) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedMultiServerServer) PickServer(ctx context.Context, req *PickReq) (*PickReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PickServer not implemented")
}

func RegisterMultiServerServer(s *grpc.Server, srv MultiServerServer) {
	s.RegisterService(&_MultiServer_serviceDesc, srv)
}

func _MultiServer_FwdMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FwdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultiServerServer).FwdMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.MultiServer/FwdMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultiServerServer).FwdMsg(ctx, req.(*FwdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MultiServer_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultiServerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.MultiServer/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultiServerServer).Ping(ctx, req.(*PingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MultiServer_PickServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PickReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultiServerServer).PickServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.MultiServer/PickServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultiServerServer).PickServer(ctx, req.(*PickReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _MultiServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.MultiServer",
	HandlerType: (*MultiServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FwdMsg",
			Handler:    _MultiServer_FwdMsg_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _MultiServer_Ping_Handler,
		},
		{
			MethodName: "PickServer",
			Handler:    _MultiServer_PickServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "multiserver.proto",
}
