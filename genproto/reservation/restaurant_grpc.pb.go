// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: restaurant.proto

package reservation

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
	RestaurantService_Create_FullMethodName = "/reservation.RestaurantService/Create"
	RestaurantService_Get_FullMethodName    = "/reservation.RestaurantService/Get"
	RestaurantService_GetAll_FullMethodName = "/reservation.RestaurantService/GetAll"
	RestaurantService_Update_FullMethodName = "/reservation.RestaurantService/Update"
	RestaurantService_Delete_FullMethodName = "/reservation.RestaurantService/Delete"
)

// RestaurantServiceClient is the client API for RestaurantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RestaurantServiceClient interface {
	Create(ctx context.Context, in *RestaurantReq, opts ...grpc.CallOption) (*Restaurant, error)
	Get(ctx context.Context, in *GetByIdReq, opts ...grpc.CallOption) (*Restaurant, error)
	GetAll(ctx context.Context, in *GetAllRestaurantReq, opts ...grpc.CallOption) (*GetAllRestaurantRes, error)
	Update(ctx context.Context, in *RestaurantUpdate, opts ...grpc.CallOption) (*Restaurant, error)
	Delete(ctx context.Context, in *GetByIdReq, opts ...grpc.CallOption) (*Void, error)
}

type restaurantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRestaurantServiceClient(cc grpc.ClientConnInterface) RestaurantServiceClient {
	return &restaurantServiceClient{cc}
}

func (c *restaurantServiceClient) Create(ctx context.Context, in *RestaurantReq, opts ...grpc.CallOption) (*Restaurant, error) {
	out := new(Restaurant)
	err := c.cc.Invoke(ctx, RestaurantService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantServiceClient) Get(ctx context.Context, in *GetByIdReq, opts ...grpc.CallOption) (*Restaurant, error) {
	out := new(Restaurant)
	err := c.cc.Invoke(ctx, RestaurantService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantServiceClient) GetAll(ctx context.Context, in *GetAllRestaurantReq, opts ...grpc.CallOption) (*GetAllRestaurantRes, error) {
	out := new(GetAllRestaurantRes)
	err := c.cc.Invoke(ctx, RestaurantService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantServiceClient) Update(ctx context.Context, in *RestaurantUpdate, opts ...grpc.CallOption) (*Restaurant, error) {
	out := new(Restaurant)
	err := c.cc.Invoke(ctx, RestaurantService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantServiceClient) Delete(ctx context.Context, in *GetByIdReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, RestaurantService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RestaurantServiceServer is the server API for RestaurantService service.
// All implementations must embed UnimplementedRestaurantServiceServer
// for forward compatibility
type RestaurantServiceServer interface {
	Create(context.Context, *RestaurantReq) (*Restaurant, error)
	Get(context.Context, *GetByIdReq) (*Restaurant, error)
	GetAll(context.Context, *GetAllRestaurantReq) (*GetAllRestaurantRes, error)
	Update(context.Context, *RestaurantUpdate) (*Restaurant, error)
	Delete(context.Context, *GetByIdReq) (*Void, error)
	mustEmbedUnimplementedRestaurantServiceServer()
}

// UnimplementedRestaurantServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRestaurantServiceServer struct {
}

func (UnimplementedRestaurantServiceServer) Create(context.Context, *RestaurantReq) (*Restaurant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRestaurantServiceServer) Get(context.Context, *GetByIdReq) (*Restaurant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedRestaurantServiceServer) GetAll(context.Context, *GetAllRestaurantReq) (*GetAllRestaurantRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedRestaurantServiceServer) Update(context.Context, *RestaurantUpdate) (*Restaurant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRestaurantServiceServer) Delete(context.Context, *GetByIdReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRestaurantServiceServer) mustEmbedUnimplementedRestaurantServiceServer() {}

// UnsafeRestaurantServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RestaurantServiceServer will
// result in compilation errors.
type UnsafeRestaurantServiceServer interface {
	mustEmbedUnimplementedRestaurantServiceServer()
}

func RegisterRestaurantServiceServer(s grpc.ServiceRegistrar, srv RestaurantServiceServer) {
	s.RegisterService(&RestaurantService_ServiceDesc, srv)
}

func _RestaurantService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestaurantReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RestaurantService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServiceServer).Create(ctx, req.(*RestaurantReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestaurantService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RestaurantService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServiceServer).Get(ctx, req.(*GetByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestaurantService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRestaurantReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RestaurantService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServiceServer).GetAll(ctx, req.(*GetAllRestaurantReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestaurantService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestaurantUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RestaurantService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServiceServer).Update(ctx, req.(*RestaurantUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestaurantService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RestaurantService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServiceServer).Delete(ctx, req.(*GetByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// RestaurantService_ServiceDesc is the grpc.ServiceDesc for RestaurantService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RestaurantService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reservation.RestaurantService",
	HandlerType: (*RestaurantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _RestaurantService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _RestaurantService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _RestaurantService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _RestaurantService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RestaurantService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "restaurant.proto",
}
