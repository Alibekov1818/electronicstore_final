// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.3
// source: computers_service.proto

package pb

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

const (
	ComputersService_GetComputer_FullMethodName    = "/auth.ComputersService/GetComputer"
	ComputersService_GetComputers_FullMethodName   = "/auth.ComputersService/GetComputers"
	ComputersService_CreateComputer_FullMethodName = "/auth.ComputersService/CreateComputer"
	ComputersService_DeleteComputer_FullMethodName = "/auth.ComputersService/DeleteComputer"
	ComputersService_UpdateComputer_FullMethodName = "/auth.ComputersService/UpdateComputer"
)

// ComputersServiceClient is the client API for ComputersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComputersServiceClient interface {
	GetComputer(ctx context.Context, in *ComputerId, opts ...grpc.CallOption) (*Computer, error)
	GetComputers(ctx context.Context, in *GetComputersRequest, opts ...grpc.CallOption) (*ComputerList, error)
	CreateComputer(ctx context.Context, in *Computer, opts ...grpc.CallOption) (*Computer, error)
	DeleteComputer(ctx context.Context, in *ComputerId, opts ...grpc.CallOption) (*Computer, error)
	UpdateComputer(ctx context.Context, in *Computer, opts ...grpc.CallOption) (*Computer, error)
}

type computersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewComputersServiceClient(cc grpc.ClientConnInterface) ComputersServiceClient {
	return &computersServiceClient{cc}
}

func (c *computersServiceClient) GetComputer(ctx context.Context, in *ComputerId, opts ...grpc.CallOption) (*Computer, error) {
	out := new(Computer)
	err := c.cc.Invoke(ctx, ComputersService_GetComputer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *computersServiceClient) GetComputers(ctx context.Context, in *GetComputersRequest, opts ...grpc.CallOption) (*ComputerList, error) {
	out := new(ComputerList)
	err := c.cc.Invoke(ctx, ComputersService_GetComputers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *computersServiceClient) CreateComputer(ctx context.Context, in *Computer, opts ...grpc.CallOption) (*Computer, error) {
	out := new(Computer)
	err := c.cc.Invoke(ctx, ComputersService_CreateComputer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *computersServiceClient) DeleteComputer(ctx context.Context, in *ComputerId, opts ...grpc.CallOption) (*Computer, error) {
	out := new(Computer)
	err := c.cc.Invoke(ctx, ComputersService_DeleteComputer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *computersServiceClient) UpdateComputer(ctx context.Context, in *Computer, opts ...grpc.CallOption) (*Computer, error) {
	out := new(Computer)
	err := c.cc.Invoke(ctx, ComputersService_UpdateComputer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComputersServiceServer is the server API for ComputersService service.
// All implementations must embed UnimplementedComputersServiceServer
// for forward compatibility
type ComputersServiceServer interface {
	GetComputer(context.Context, *ComputerId) (*Computer, error)
	GetComputers(context.Context, *GetComputersRequest) (*ComputerList, error)
	CreateComputer(context.Context, *Computer) (*Computer, error)
	DeleteComputer(context.Context, *ComputerId) (*Computer, error)
	UpdateComputer(context.Context, *Computer) (*Computer, error)
	mustEmbedUnimplementedComputersServiceServer()
}

// UnimplementedComputersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedComputersServiceServer struct {
}

func (UnimplementedComputersServiceServer) GetComputer(context.Context, *ComputerId) (*Computer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComputer not implemented")
}
func (UnimplementedComputersServiceServer) GetComputers(context.Context, *GetComputersRequest) (*ComputerList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComputers not implemented")
}
func (UnimplementedComputersServiceServer) CreateComputer(context.Context, *Computer) (*Computer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComputer not implemented")
}
func (UnimplementedComputersServiceServer) DeleteComputer(context.Context, *ComputerId) (*Computer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComputer not implemented")
}
func (UnimplementedComputersServiceServer) UpdateComputer(context.Context, *Computer) (*Computer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateComputer not implemented")
}
func (UnimplementedComputersServiceServer) mustEmbedUnimplementedComputersServiceServer() {}

// UnsafeComputersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComputersServiceServer will
// result in compilation errors.
type UnsafeComputersServiceServer interface {
	mustEmbedUnimplementedComputersServiceServer()
}

func RegisterComputersServiceServer(s grpc.ServiceRegistrar, srv ComputersServiceServer) {
	s.RegisterService(&ComputersService_ServiceDesc, srv)
}

func _ComputersService_GetComputer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComputerId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputersServiceServer).GetComputer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ComputersService_GetComputer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputersServiceServer).GetComputer(ctx, req.(*ComputerId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComputersService_GetComputers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetComputersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputersServiceServer).GetComputers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ComputersService_GetComputers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputersServiceServer).GetComputers(ctx, req.(*GetComputersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComputersService_CreateComputer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Computer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputersServiceServer).CreateComputer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ComputersService_CreateComputer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputersServiceServer).CreateComputer(ctx, req.(*Computer))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComputersService_DeleteComputer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComputerId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputersServiceServer).DeleteComputer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ComputersService_DeleteComputer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputersServiceServer).DeleteComputer(ctx, req.(*ComputerId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComputersService_UpdateComputer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Computer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputersServiceServer).UpdateComputer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ComputersService_UpdateComputer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputersServiceServer).UpdateComputer(ctx, req.(*Computer))
	}
	return interceptor(ctx, in, info, handler)
}

// ComputersService_ServiceDesc is the grpc.ServiceDesc for ComputersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ComputersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.ComputersService",
	HandlerType: (*ComputersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetComputer",
			Handler:    _ComputersService_GetComputer_Handler,
		},
		{
			MethodName: "GetComputers",
			Handler:    _ComputersService_GetComputers_Handler,
		},
		{
			MethodName: "CreateComputer",
			Handler:    _ComputersService_CreateComputer_Handler,
		},
		{
			MethodName: "DeleteComputer",
			Handler:    _ComputersService_DeleteComputer_Handler,
		},
		{
			MethodName: "UpdateComputer",
			Handler:    _ComputersService_UpdateComputer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "computers_service.proto",
}