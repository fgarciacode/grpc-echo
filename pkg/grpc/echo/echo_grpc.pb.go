// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: echo.proto

package echo

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EchoServiceClient is the client API for EchoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EchoServiceClient interface {
	EchoUnary(ctx context.Context, in *EchoMessage, opts ...grpc.CallOption) (*EchoMessage, error)
	EchoClientStream(ctx context.Context, opts ...grpc.CallOption) (EchoService_EchoClientStreamClient, error)
	EchoServerStream(ctx context.Context, in *EchoMessage, opts ...grpc.CallOption) (EchoService_EchoServerStreamClient, error)
	EchoBidiStream(ctx context.Context, opts ...grpc.CallOption) (EchoService_EchoBidiStreamClient, error)
	EchoStatus(ctx context.Context, in *StatusCode, opts ...grpc.CallOption) (*StatusCode, error)
}

type echoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEchoServiceClient(cc grpc.ClientConnInterface) EchoServiceClient {
	return &echoServiceClient{cc}
}

func (c *echoServiceClient) EchoUnary(ctx context.Context, in *EchoMessage, opts ...grpc.CallOption) (*EchoMessage, error) {
	out := new(EchoMessage)
	err := c.cc.Invoke(ctx, "/echo.EchoService/EchoUnary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoServiceClient) EchoClientStream(ctx context.Context, opts ...grpc.CallOption) (EchoService_EchoClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &EchoService_ServiceDesc.Streams[0], "/echo.EchoService/EchoClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoServiceEchoClientStreamClient{stream}
	return x, nil
}

type EchoService_EchoClientStreamClient interface {
	Send(*EchoMessage) error
	CloseAndRecv() (*EchoMessage, error)
	grpc.ClientStream
}

type echoServiceEchoClientStreamClient struct {
	grpc.ClientStream
}

func (x *echoServiceEchoClientStreamClient) Send(m *EchoMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *echoServiceEchoClientStreamClient) CloseAndRecv() (*EchoMessage, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(EchoMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *echoServiceClient) EchoServerStream(ctx context.Context, in *EchoMessage, opts ...grpc.CallOption) (EchoService_EchoServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &EchoService_ServiceDesc.Streams[1], "/echo.EchoService/EchoServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoServiceEchoServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EchoService_EchoServerStreamClient interface {
	Recv() (*EchoMessage, error)
	grpc.ClientStream
}

type echoServiceEchoServerStreamClient struct {
	grpc.ClientStream
}

func (x *echoServiceEchoServerStreamClient) Recv() (*EchoMessage, error) {
	m := new(EchoMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *echoServiceClient) EchoBidiStream(ctx context.Context, opts ...grpc.CallOption) (EchoService_EchoBidiStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &EchoService_ServiceDesc.Streams[2], "/echo.EchoService/EchoBidiStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoServiceEchoBidiStreamClient{stream}
	return x, nil
}

type EchoService_EchoBidiStreamClient interface {
	Send(*EchoMessage) error
	Recv() (*EchoMessage, error)
	grpc.ClientStream
}

type echoServiceEchoBidiStreamClient struct {
	grpc.ClientStream
}

func (x *echoServiceEchoBidiStreamClient) Send(m *EchoMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *echoServiceEchoBidiStreamClient) Recv() (*EchoMessage, error) {
	m := new(EchoMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *echoServiceClient) EchoStatus(ctx context.Context, in *StatusCode, opts ...grpc.CallOption) (*StatusCode, error) {
	out := new(StatusCode)
	err := c.cc.Invoke(ctx, "/echo.EchoService/EchoStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoServiceServer is the server API for EchoService service.
// All implementations must embed UnimplementedEchoServiceServer
// for forward compatibility
type EchoServiceServer interface {
	EchoUnary(context.Context, *EchoMessage) (*EchoMessage, error)
	EchoClientStream(EchoService_EchoClientStreamServer) error
	EchoServerStream(*EchoMessage, EchoService_EchoServerStreamServer) error
	EchoBidiStream(EchoService_EchoBidiStreamServer) error
	EchoStatus(context.Context, *StatusCode) (*StatusCode, error)
	mustEmbedUnimplementedEchoServiceServer()
}

// UnimplementedEchoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEchoServiceServer struct {
}

func (UnimplementedEchoServiceServer) EchoUnary(context.Context, *EchoMessage) (*EchoMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EchoUnary not implemented")
}
func (UnimplementedEchoServiceServer) EchoClientStream(EchoService_EchoClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method EchoClientStream not implemented")
}
func (UnimplementedEchoServiceServer) EchoServerStream(*EchoMessage, EchoService_EchoServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method EchoServerStream not implemented")
}
func (UnimplementedEchoServiceServer) EchoBidiStream(EchoService_EchoBidiStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method EchoBidiStream not implemented")
}
func (UnimplementedEchoServiceServer) EchoStatus(context.Context, *StatusCode) (*StatusCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EchoStatus not implemented")
}
func (UnimplementedEchoServiceServer) mustEmbedUnimplementedEchoServiceServer() {}

// UnsafeEchoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EchoServiceServer will
// result in compilation errors.
type UnsafeEchoServiceServer interface {
	mustEmbedUnimplementedEchoServiceServer()
}

func RegisterEchoServiceServer(s grpc.ServiceRegistrar, srv EchoServiceServer) {
	s.RegisterService(&EchoService_ServiceDesc, srv)
}

func _EchoService_EchoUnary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).EchoUnary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.EchoService/EchoUnary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).EchoUnary(ctx, req.(*EchoMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _EchoService_EchoClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EchoServiceServer).EchoClientStream(&echoServiceEchoClientStreamServer{stream})
}

type EchoService_EchoClientStreamServer interface {
	SendAndClose(*EchoMessage) error
	Recv() (*EchoMessage, error)
	grpc.ServerStream
}

type echoServiceEchoClientStreamServer struct {
	grpc.ServerStream
}

func (x *echoServiceEchoClientStreamServer) SendAndClose(m *EchoMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *echoServiceEchoClientStreamServer) Recv() (*EchoMessage, error) {
	m := new(EchoMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EchoService_EchoServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EchoMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EchoServiceServer).EchoServerStream(m, &echoServiceEchoServerStreamServer{stream})
}

type EchoService_EchoServerStreamServer interface {
	Send(*EchoMessage) error
	grpc.ServerStream
}

type echoServiceEchoServerStreamServer struct {
	grpc.ServerStream
}

func (x *echoServiceEchoServerStreamServer) Send(m *EchoMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _EchoService_EchoBidiStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EchoServiceServer).EchoBidiStream(&echoServiceEchoBidiStreamServer{stream})
}

type EchoService_EchoBidiStreamServer interface {
	Send(*EchoMessage) error
	Recv() (*EchoMessage, error)
	grpc.ServerStream
}

type echoServiceEchoBidiStreamServer struct {
	grpc.ServerStream
}

func (x *echoServiceEchoBidiStreamServer) Send(m *EchoMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *echoServiceEchoBidiStreamServer) Recv() (*EchoMessage, error) {
	m := new(EchoMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EchoService_EchoStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusCode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).EchoStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.EchoService/EchoStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).EchoStatus(ctx, req.(*StatusCode))
	}
	return interceptor(ctx, in, info, handler)
}

// EchoService_ServiceDesc is the grpc.ServiceDesc for EchoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EchoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "echo.EchoService",
	HandlerType: (*EchoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EchoUnary",
			Handler:    _EchoService_EchoUnary_Handler,
		},
		{
			MethodName: "EchoStatus",
			Handler:    _EchoService_EchoStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EchoClientStream",
			Handler:       _EchoService_EchoClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "EchoServerStream",
			Handler:       _EchoService_EchoServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "EchoBidiStream",
			Handler:       _EchoService_EchoBidiStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "echo.proto",
}
