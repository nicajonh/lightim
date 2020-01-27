// Code generated by protoc-gen-go. DO NOT EDIT.
// source: logic_server.proto

package pb

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

func init() { proto.RegisterFile("logic_server.proto", fileDescriptor_1da729ec9e401d95) }

var fileDescriptor_1da729ec9e401d95 = []byte{
	// 112 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xca, 0xc9, 0x4f, 0xcf,
	0x4c, 0x8e, 0x2f, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x2a, 0x48, 0x92, 0x82, 0x8a, 0x27, 0xe7, 0x64, 0xa6, 0xe6, 0x95, 0x40, 0xc4, 0x8d, 0xbc, 0xb9,
	0x04, 0x7d, 0x40, 0xa2, 0x6e, 0xf9, 0x45, 0xc1, 0x60, 0xf5, 0xae, 0x15, 0x25, 0x42, 0x66, 0x5c,
	0xdc, 0xc1, 0xa9, 0x79, 0x29, 0xbe, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x42, 0x42, 0x7a, 0x05,
	0x49, 0x7a, 0x48, 0x02, 0x41, 0xa9, 0x85, 0x52, 0xc2, 0x18, 0x62, 0xc5, 0x05, 0x49, 0x6c, 0x60,
	0x33, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x33, 0x1c, 0xd5, 0x81, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LogicForServerExtClient is the client API for LogicForServerExt service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogicForServerExtClient interface {
	// 发送消息
	SendMessage(ctx context.Context, in *SendMessageReq, opts ...grpc.CallOption) (*SendMessageResp, error)
}

type logicForServerExtClient struct {
	cc *grpc.ClientConn
}

func NewLogicForServerExtClient(cc *grpc.ClientConn) LogicForServerExtClient {
	return &logicForServerExtClient{cc}
}

func (c *logicForServerExtClient) SendMessage(ctx context.Context, in *SendMessageReq, opts ...grpc.CallOption) (*SendMessageResp, error) {
	out := new(SendMessageResp)
	err := c.cc.Invoke(ctx, "/pb.LogicForServerExt/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogicForServerExtServer is the server API for LogicForServerExt service.
type LogicForServerExtServer interface {
	// 发送消息
	SendMessage(context.Context, *SendMessageReq) (*SendMessageResp, error)
}

// UnimplementedLogicForServerExtServer can be embedded to have forward compatible implementations.
type UnimplementedLogicForServerExtServer struct {
}

func (*UnimplementedLogicForServerExtServer) SendMessage(ctx context.Context, req *SendMessageReq) (*SendMessageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}

func RegisterLogicForServerExtServer(s *grpc.Server, srv LogicForServerExtServer) {
	s.RegisterService(&_LogicForServerExt_serviceDesc, srv)
}

func _LogicForServerExt_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicForServerExtServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LogicForServerExt/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicForServerExtServer).SendMessage(ctx, req.(*SendMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogicForServerExt_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.LogicForServerExt",
	HandlerType: (*LogicForServerExtServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _LogicForServerExt_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "logic_server.proto",
}
