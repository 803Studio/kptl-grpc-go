// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.2
// source: employee/Employee.proto

package employee

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
	Employee_Deliver_FullMethodName = "/employee.Employee/Deliver"
	Employee_Browses_FullMethodName = "/employee.Employee/Browses"
)

// EmployeeClient is the client API for Employee service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmployeeClient interface {
	Deliver(ctx context.Context, in *OperateRequest, opts ...grpc.CallOption) (*OperateResponse, error)
	Browses(ctx context.Context, in *OperateRequest, opts ...grpc.CallOption) (*OperateResponse, error)
}

type employeeClient struct {
	cc grpc.ClientConnInterface
}

func NewEmployeeClient(cc grpc.ClientConnInterface) EmployeeClient {
	return &employeeClient{cc}
}

func (c *employeeClient) Deliver(ctx context.Context, in *OperateRequest, opts ...grpc.CallOption) (*OperateResponse, error) {
	out := new(OperateResponse)
	err := c.cc.Invoke(ctx, Employee_Deliver_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeClient) Browses(ctx context.Context, in *OperateRequest, opts ...grpc.CallOption) (*OperateResponse, error) {
	out := new(OperateResponse)
	err := c.cc.Invoke(ctx, Employee_Browses_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmployeeServer is the server API for Employee service.
// All implementations must embed UnimplementedEmployeeServer
// for forward compatibility
type EmployeeServer interface {
	Deliver(context.Context, *OperateRequest) (*OperateResponse, error)
	Browses(context.Context, *OperateRequest) (*OperateResponse, error)
	mustEmbedUnimplementedEmployeeServer()
}

// UnimplementedEmployeeServer must be embedded to have forward compatible implementations.
type UnimplementedEmployeeServer struct {
}

func (UnimplementedEmployeeServer) Deliver(context.Context, *OperateRequest) (*OperateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deliver not implemented")
}
func (UnimplementedEmployeeServer) Browses(context.Context, *OperateRequest) (*OperateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Browses not implemented")
}
func (UnimplementedEmployeeServer) mustEmbedUnimplementedEmployeeServer() {}

// UnsafeEmployeeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmployeeServer will
// result in compilation errors.
type UnsafeEmployeeServer interface {
	mustEmbedUnimplementedEmployeeServer()
}

func RegisterEmployeeServer(s grpc.ServiceRegistrar, srv EmployeeServer) {
	s.RegisterService(&Employee_ServiceDesc, srv)
}

func _Employee_Deliver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServer).Deliver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Employee_Deliver_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServer).Deliver(ctx, req.(*OperateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Employee_Browses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServer).Browses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Employee_Browses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServer).Browses(ctx, req.(*OperateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Employee_ServiceDesc is the grpc.ServiceDesc for Employee service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Employee_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "employee.Employee",
	HandlerType: (*EmployeeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Deliver",
			Handler:    _Employee_Deliver_Handler,
		},
		{
			MethodName: "Browses",
			Handler:    _Employee_Browses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "employee/Employee.proto",
}