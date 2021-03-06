// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

/*
Package hello is a generated protocol buffer package.

It is generated from these files:
	hello.proto

It has these top-level messages:
	Knock
	Reply
*/
package hello

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// *
// Knock the door or not
type Knock struct {
	KnockDoor bool `protobuf:"varint,1,opt,name=knockDoor" json:"knockDoor,omitempty"`
}

func (m *Knock) Reset()                    { *m = Knock{} }
func (m *Knock) String() string            { return proto.CompactTextString(m) }
func (*Knock) ProtoMessage()               {}
func (*Knock) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Knock) GetKnockDoor() bool {
	if m != nil {
		return m.KnockDoor
	}
	return false
}

// *
// Reply based on whether a door was knocked or not.
type Reply struct {
	Reply        bool   `protobuf:"varint,1,opt,name=reply" json:"reply,omitempty"`
	ReplyMessage string `protobuf:"bytes,2,opt,name=replyMessage" json:"replyMessage,omitempty"`
}

func (m *Reply) Reset()                    { *m = Reply{} }
func (m *Reply) String() string            { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()               {}
func (*Reply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Reply) GetReply() bool {
	if m != nil {
		return m.Reply
	}
	return false
}

func (m *Reply) GetReplyMessage() string {
	if m != nil {
		return m.ReplyMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*Knock)(nil), "hello.Knock")
	proto.RegisterType((*Reply)(nil), "hello.Reply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for HelloService service

type HelloServiceClient interface {
	// *
	// Determines if you knocked the door and the appropriate
	// reply if you didnt.
	GetHello(ctx context.Context, in *Knock, opts ...grpc.CallOption) (*Reply, error)
}

type helloServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloServiceClient(cc *grpc.ClientConn) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) GetHello(ctx context.Context, in *Knock, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/hello.HelloService/GetHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HelloService service

type HelloServiceServer interface {
	// *
	// Determines if you knocked the door and the appropriate
	// reply if you didnt.
	GetHello(context.Context, *Knock) (*Reply, error)
}

func RegisterHelloServiceServer(s *grpc.Server, srv HelloServiceServer) {
	s.RegisterService(&_HelloService_serviceDesc, srv)
}

func _HelloService_GetHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Knock)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).GetHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.HelloService/GetHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).GetHello(ctx, req.(*Knock))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHello",
			Handler:    _HelloService_GetHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0x54, 0xb9, 0x58, 0xbd,
	0xf3, 0xf2, 0x93, 0xb3, 0x85, 0x64, 0xb8, 0x38, 0xb3, 0x41, 0x0c, 0x97, 0xfc, 0xfc, 0x22, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0x8e, 0x20, 0x84, 0x80, 0x92, 0x23, 0x17, 0x6b, 0x50, 0x6a, 0x41, 0x4e,
	0xa5, 0x90, 0x08, 0x17, 0x6b, 0x11, 0x88, 0x01, 0x55, 0x02, 0xe1, 0x08, 0x29, 0x71, 0xf1, 0x80,
	0x19, 0xbe, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x28,
	0x62, 0x46, 0x16, 0x5c, 0x3c, 0x1e, 0x20, 0x2b, 0x83, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85,
	0x34, 0xb8, 0x38, 0xdc, 0x53, 0x4b, 0xc0, 0x42, 0x42, 0x3c, 0x7a, 0x10, 0xa7, 0x81, 0x9d, 0x22,
	0x05, 0xe3, 0x81, 0x6d, 0x54, 0x62, 0x48, 0x62, 0x03, 0xbb, 0xd8, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x3a, 0xff, 0x1a, 0x47, 0xc0, 0x00, 0x00, 0x00,
}
