// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.2
// source: job/Job.proto

package job

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
	Job_FindAllJobs_FullMethodName       = "/job.Job/FindAllJobs"
	Job_FindJobByName_FullMethodName     = "/job.Job/FindJobByName"
	Job_FindJobById_FullMethodName       = "/job.Job/FindJobById"
	Job_FindJobs_FullMethodName          = "/job.Job/FindJobs"
	Job_SaveJobs_FullMethodName          = "/job.Job/SaveJobs"
	Job_UpdateJobs_FullMethodName        = "/job.Job/UpdateJobs"
	Job_RegisteredCompany_FullMethodName = "/job.Job/RegisteredCompany"
	Job_UpdateCompany_FullMethodName     = "/job.Job/UpdateCompany"
	Job_BoundCompany_FullMethodName      = "/job.Job/BoundCompany"
	Job_VerifyCompany_FullMethodName     = "/job.Job/VerifyCompany"
)

// JobClient is the client API for Job service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JobClient interface {
	FindAllJobs(ctx context.Context, in *FindAllJobRequest, opts ...grpc.CallOption) (*FindJobResponse, error)
	FindJobByName(ctx context.Context, in *FindJobByNameRequest, opts ...grpc.CallOption) (*FindJobResponse, error)
	FindJobById(ctx context.Context, in *FindJobByIdRequest, opts ...grpc.CallOption) (*FindJobResponse, error)
	FindJobs(ctx context.Context, in *FindJobRequest, opts ...grpc.CallOption) (*FindJobResponse, error)
	SaveJobs(ctx context.Context, in *JobMessage, opts ...grpc.CallOption) (*SaveJobResponse, error)
	UpdateJobs(ctx context.Context, in *JobMessage, opts ...grpc.CallOption) (*SaveJobResponse, error)
	RegisteredCompany(ctx context.Context, in *RegisteredCompanyReq, opts ...grpc.CallOption) (*CommonResponse, error)
	UpdateCompany(ctx context.Context, in *RegisteredCompanyReq, opts ...grpc.CallOption) (*CommonResponse, error)
	BoundCompany(ctx context.Context, in *BoundCompanyReq, opts ...grpc.CallOption) (*CommonResponse, error)
	VerifyCompany(ctx context.Context, in *VerifyCompanyReq, opts ...grpc.CallOption) (*CommonResponse, error)
}

type jobClient struct {
	cc grpc.ClientConnInterface
}

func NewJobClient(cc grpc.ClientConnInterface) JobClient {
	return &jobClient{cc}
}

func (c *jobClient) FindAllJobs(ctx context.Context, in *FindAllJobRequest, opts ...grpc.CallOption) (*FindJobResponse, error) {
	out := new(FindJobResponse)
	err := c.cc.Invoke(ctx, Job_FindAllJobs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) FindJobByName(ctx context.Context, in *FindJobByNameRequest, opts ...grpc.CallOption) (*FindJobResponse, error) {
	out := new(FindJobResponse)
	err := c.cc.Invoke(ctx, Job_FindJobByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) FindJobById(ctx context.Context, in *FindJobByIdRequest, opts ...grpc.CallOption) (*FindJobResponse, error) {
	out := new(FindJobResponse)
	err := c.cc.Invoke(ctx, Job_FindJobById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) FindJobs(ctx context.Context, in *FindJobRequest, opts ...grpc.CallOption) (*FindJobResponse, error) {
	out := new(FindJobResponse)
	err := c.cc.Invoke(ctx, Job_FindJobs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) SaveJobs(ctx context.Context, in *JobMessage, opts ...grpc.CallOption) (*SaveJobResponse, error) {
	out := new(SaveJobResponse)
	err := c.cc.Invoke(ctx, Job_SaveJobs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) UpdateJobs(ctx context.Context, in *JobMessage, opts ...grpc.CallOption) (*SaveJobResponse, error) {
	out := new(SaveJobResponse)
	err := c.cc.Invoke(ctx, Job_UpdateJobs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) RegisteredCompany(ctx context.Context, in *RegisteredCompanyReq, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, Job_RegisteredCompany_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) UpdateCompany(ctx context.Context, in *RegisteredCompanyReq, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, Job_UpdateCompany_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) BoundCompany(ctx context.Context, in *BoundCompanyReq, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, Job_BoundCompany_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) VerifyCompany(ctx context.Context, in *VerifyCompanyReq, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, Job_VerifyCompany_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobServer is the server API for Job service.
// All implementations must embed UnimplementedJobServer
// for forward compatibility
type JobServer interface {
	FindAllJobs(context.Context, *FindAllJobRequest) (*FindJobResponse, error)
	FindJobByName(context.Context, *FindJobByNameRequest) (*FindJobResponse, error)
	FindJobById(context.Context, *FindJobByIdRequest) (*FindJobResponse, error)
	FindJobs(context.Context, *FindJobRequest) (*FindJobResponse, error)
	SaveJobs(context.Context, *JobMessage) (*SaveJobResponse, error)
	UpdateJobs(context.Context, *JobMessage) (*SaveJobResponse, error)
	RegisteredCompany(context.Context, *RegisteredCompanyReq) (*CommonResponse, error)
	UpdateCompany(context.Context, *RegisteredCompanyReq) (*CommonResponse, error)
	BoundCompany(context.Context, *BoundCompanyReq) (*CommonResponse, error)
	VerifyCompany(context.Context, *VerifyCompanyReq) (*CommonResponse, error)
	mustEmbedUnimplementedJobServer()
}

// UnimplementedJobServer must be embedded to have forward compatible implementations.
type UnimplementedJobServer struct {
}

func (UnimplementedJobServer) FindAllJobs(context.Context, *FindAllJobRequest) (*FindJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllJobs not implemented")
}
func (UnimplementedJobServer) FindJobByName(context.Context, *FindJobByNameRequest) (*FindJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindJobByName not implemented")
}
func (UnimplementedJobServer) FindJobById(context.Context, *FindJobByIdRequest) (*FindJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindJobById not implemented")
}
func (UnimplementedJobServer) FindJobs(context.Context, *FindJobRequest) (*FindJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindJobs not implemented")
}
func (UnimplementedJobServer) SaveJobs(context.Context, *JobMessage) (*SaveJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveJobs not implemented")
}
func (UnimplementedJobServer) UpdateJobs(context.Context, *JobMessage) (*SaveJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateJobs not implemented")
}
func (UnimplementedJobServer) RegisteredCompany(context.Context, *RegisteredCompanyReq) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisteredCompany not implemented")
}
func (UnimplementedJobServer) UpdateCompany(context.Context, *RegisteredCompanyReq) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCompany not implemented")
}
func (UnimplementedJobServer) BoundCompany(context.Context, *BoundCompanyReq) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BoundCompany not implemented")
}
func (UnimplementedJobServer) VerifyCompany(context.Context, *VerifyCompanyReq) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyCompany not implemented")
}
func (UnimplementedJobServer) mustEmbedUnimplementedJobServer() {}

// UnsafeJobServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JobServer will
// result in compilation errors.
type UnsafeJobServer interface {
	mustEmbedUnimplementedJobServer()
}

func RegisterJobServer(s grpc.ServiceRegistrar, srv JobServer) {
	s.RegisterService(&Job_ServiceDesc, srv)
}

func _Job_FindAllJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).FindAllJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_FindAllJobs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).FindAllJobs(ctx, req.(*FindAllJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_FindJobByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindJobByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).FindJobByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_FindJobByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).FindJobByName(ctx, req.(*FindJobByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_FindJobById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindJobByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).FindJobById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_FindJobById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).FindJobById(ctx, req.(*FindJobByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_FindJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).FindJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_FindJobs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).FindJobs(ctx, req.(*FindJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_SaveJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).SaveJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_SaveJobs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).SaveJobs(ctx, req.(*JobMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_UpdateJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).UpdateJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_UpdateJobs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).UpdateJobs(ctx, req.(*JobMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_RegisteredCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisteredCompanyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).RegisteredCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_RegisteredCompany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).RegisteredCompany(ctx, req.(*RegisteredCompanyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_UpdateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisteredCompanyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).UpdateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_UpdateCompany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).UpdateCompany(ctx, req.(*RegisteredCompanyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_BoundCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BoundCompanyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).BoundCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_BoundCompany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).BoundCompany(ctx, req.(*BoundCompanyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_VerifyCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyCompanyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).VerifyCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_VerifyCompany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).VerifyCompany(ctx, req.(*VerifyCompanyReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Job_ServiceDesc is the grpc.ServiceDesc for Job service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Job_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "job.Job",
	HandlerType: (*JobServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindAllJobs",
			Handler:    _Job_FindAllJobs_Handler,
		},
		{
			MethodName: "FindJobByName",
			Handler:    _Job_FindJobByName_Handler,
		},
		{
			MethodName: "FindJobById",
			Handler:    _Job_FindJobById_Handler,
		},
		{
			MethodName: "FindJobs",
			Handler:    _Job_FindJobs_Handler,
		},
		{
			MethodName: "SaveJobs",
			Handler:    _Job_SaveJobs_Handler,
		},
		{
			MethodName: "UpdateJobs",
			Handler:    _Job_UpdateJobs_Handler,
		},
		{
			MethodName: "RegisteredCompany",
			Handler:    _Job_RegisteredCompany_Handler,
		},
		{
			MethodName: "UpdateCompany",
			Handler:    _Job_UpdateCompany_Handler,
		},
		{
			MethodName: "BoundCompany",
			Handler:    _Job_BoundCompany_Handler,
		},
		{
			MethodName: "VerifyCompany",
			Handler:    _Job_VerifyCompany_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "job/Job.proto",
}
