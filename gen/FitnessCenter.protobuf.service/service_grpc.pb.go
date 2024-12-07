// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: service.proto

package FitnessCenter_protobuf_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Service_CreateService_FullMethodName       = "/fitness_center.service.Service/CreateService"
	Service_GetServiceById_FullMethodName      = "/fitness_center.service.Service/GetServiceById"
	Service_UpdateService_FullMethodName       = "/fitness_center.service.Service/UpdateService"
	Service_DeleteServiceById_FullMethodName   = "/fitness_center.service.Service/DeleteServiceById"
	Service_GetServices_FullMethodName         = "/fitness_center.service.Service/GetServices"
	Service_CreateCoachServices_FullMethodName = "/fitness_center.service.Service/CreateCoachServices"
)

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	CreateService(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[CreateServiceRequest, CreateServiceResponse], error)
	GetServiceById(ctx context.Context, in *GetServiceByIdRequest, opts ...grpc.CallOption) (*GetServiceByIdResponse, error)
	UpdateService(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[UpdateServiceRequest, UpdateServiceResponse], error)
	DeleteServiceById(ctx context.Context, in *DeleteServiceByIdRequest, opts ...grpc.CallOption) (*DeleteServiceByIdResponse, error)
	GetServices(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetServicesResponse, error)
	CreateCoachServices(ctx context.Context, in *CreateCoachServicesRequest, opts ...grpc.CallOption) (*CreateCoachServicesResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) CreateService(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[CreateServiceRequest, CreateServiceResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Service_ServiceDesc.Streams[0], Service_CreateService_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[CreateServiceRequest, CreateServiceResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Service_CreateServiceClient = grpc.ClientStreamingClient[CreateServiceRequest, CreateServiceResponse]

func (c *serviceClient) GetServiceById(ctx context.Context, in *GetServiceByIdRequest, opts ...grpc.CallOption) (*GetServiceByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetServiceByIdResponse)
	err := c.cc.Invoke(ctx, Service_GetServiceById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpdateService(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[UpdateServiceRequest, UpdateServiceResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Service_ServiceDesc.Streams[1], Service_UpdateService_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[UpdateServiceRequest, UpdateServiceResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Service_UpdateServiceClient = grpc.ClientStreamingClient[UpdateServiceRequest, UpdateServiceResponse]

func (c *serviceClient) DeleteServiceById(ctx context.Context, in *DeleteServiceByIdRequest, opts ...grpc.CallOption) (*DeleteServiceByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteServiceByIdResponse)
	err := c.cc.Invoke(ctx, Service_DeleteServiceById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetServices(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetServicesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetServicesResponse)
	err := c.cc.Invoke(ctx, Service_GetServices_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) CreateCoachServices(ctx context.Context, in *CreateCoachServicesRequest, opts ...grpc.CallOption) (*CreateCoachServicesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCoachServicesResponse)
	err := c.cc.Invoke(ctx, Service_CreateCoachServices_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility.
type ServiceServer interface {
	CreateService(grpc.ClientStreamingServer[CreateServiceRequest, CreateServiceResponse]) error
	GetServiceById(context.Context, *GetServiceByIdRequest) (*GetServiceByIdResponse, error)
	UpdateService(grpc.ClientStreamingServer[UpdateServiceRequest, UpdateServiceResponse]) error
	DeleteServiceById(context.Context, *DeleteServiceByIdRequest) (*DeleteServiceByIdResponse, error)
	GetServices(context.Context, *emptypb.Empty) (*GetServicesResponse, error)
	CreateCoachServices(context.Context, *CreateCoachServicesRequest) (*CreateCoachServicesResponse, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedServiceServer struct{}

func (UnimplementedServiceServer) CreateService(grpc.ClientStreamingServer[CreateServiceRequest, CreateServiceResponse]) error {
	return status.Errorf(codes.Unimplemented, "method CreateService not implemented")
}
func (UnimplementedServiceServer) GetServiceById(context.Context, *GetServiceByIdRequest) (*GetServiceByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServiceById not implemented")
}
func (UnimplementedServiceServer) UpdateService(grpc.ClientStreamingServer[UpdateServiceRequest, UpdateServiceResponse]) error {
	return status.Errorf(codes.Unimplemented, "method UpdateService not implemented")
}
func (UnimplementedServiceServer) DeleteServiceById(context.Context, *DeleteServiceByIdRequest) (*DeleteServiceByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteServiceById not implemented")
}
func (UnimplementedServiceServer) GetServices(context.Context, *emptypb.Empty) (*GetServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServices not implemented")
}
func (UnimplementedServiceServer) CreateCoachServices(context.Context, *CreateCoachServicesRequest) (*CreateCoachServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCoachServices not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}
func (UnimplementedServiceServer) testEmbeddedByValue()                 {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	// If the following call pancis, it indicates UnimplementedServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_CreateService_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServiceServer).CreateService(&grpc.GenericServerStream[CreateServiceRequest, CreateServiceResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Service_CreateServiceServer = grpc.ClientStreamingServer[CreateServiceRequest, CreateServiceResponse]

func _Service_GetServiceById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetServiceById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_GetServiceById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetServiceById(ctx, req.(*GetServiceByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_UpdateService_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServiceServer).UpdateService(&grpc.GenericServerStream[UpdateServiceRequest, UpdateServiceResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Service_UpdateServiceServer = grpc.ClientStreamingServer[UpdateServiceRequest, UpdateServiceResponse]

func _Service_DeleteServiceById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteServiceByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteServiceById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_DeleteServiceById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteServiceById(ctx, req.(*DeleteServiceByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_GetServices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetServices(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_CreateCoachServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCoachServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CreateCoachServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_CreateCoachServices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CreateCoachServices(ctx, req.(*CreateCoachServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fitness_center.service.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServiceById",
			Handler:    _Service_GetServiceById_Handler,
		},
		{
			MethodName: "DeleteServiceById",
			Handler:    _Service_DeleteServiceById_Handler,
		},
		{
			MethodName: "GetServices",
			Handler:    _Service_GetServices_Handler,
		},
		{
			MethodName: "CreateCoachServices",
			Handler:    _Service_CreateCoachServices_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateService",
			Handler:       _Service_CreateService_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "UpdateService",
			Handler:       _Service_UpdateService_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "service.proto",
}
