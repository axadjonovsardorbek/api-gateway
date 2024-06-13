// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: reservation_orders.proto

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

// ReservationOrderServiceClient is the client API for ReservationOrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReservationOrderServiceClient interface {
	Create(ctx context.Context, in *ReservationOrderReq, opts ...grpc.CallOption) (*ReservationOrderRes, error)
	Get(ctx context.Context, in *GetByIdReq, opts ...grpc.CallOption) (*ReservationOrderRes, error)
	GetAll(ctx context.Context, in *GetAllReservationOrderReq, opts ...grpc.CallOption) (*GetAllReservationOrderRes, error)
	Update(ctx context.Context, in *ReservationOrderUpdateReq, opts ...grpc.CallOption) (*ReservationOrderRes, error)
	Delete(ctx context.Context, in *GetByIdReq, opts ...grpc.CallOption) (*Void, error)
}

type reservationOrderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReservationOrderServiceClient(cc grpc.ClientConnInterface) ReservationOrderServiceClient {
	return &reservationOrderServiceClient{cc}
}

func (c *reservationOrderServiceClient) Create(ctx context.Context, in *ReservationOrderReq, opts ...grpc.CallOption) (*ReservationOrderRes, error) {
	out := new(ReservationOrderRes)
	err := c.cc.Invoke(ctx, "/reservation.ReservationOrderService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationOrderServiceClient) Get(ctx context.Context, in *GetByIdReq, opts ...grpc.CallOption) (*ReservationOrderRes, error) {
	out := new(ReservationOrderRes)
	err := c.cc.Invoke(ctx, "/reservation.ReservationOrderService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationOrderServiceClient) GetAll(ctx context.Context, in *GetAllReservationOrderReq, opts ...grpc.CallOption) (*GetAllReservationOrderRes, error) {
	out := new(GetAllReservationOrderRes)
	err := c.cc.Invoke(ctx, "/reservation.ReservationOrderService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationOrderServiceClient) Update(ctx context.Context, in *ReservationOrderUpdateReq, opts ...grpc.CallOption) (*ReservationOrderRes, error) {
	out := new(ReservationOrderRes)
	err := c.cc.Invoke(ctx, "/reservation.ReservationOrderService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationOrderServiceClient) Delete(ctx context.Context, in *GetByIdReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/reservation.ReservationOrderService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReservationOrderServiceServer is the server API for ReservationOrderService service.
// All implementations must embed UnimplementedReservationOrderServiceServer
// for forward compatibility
type ReservationOrderServiceServer interface {
	Create(context.Context, *ReservationOrderReq) (*ReservationOrderRes, error)
	Get(context.Context, *GetByIdReq) (*ReservationOrderRes, error)
	GetAll(context.Context, *GetAllReservationOrderReq) (*GetAllReservationOrderRes, error)
	Update(context.Context, *ReservationOrderUpdateReq) (*ReservationOrderRes, error)
	Delete(context.Context, *GetByIdReq) (*Void, error)
	mustEmbedUnimplementedReservationOrderServiceServer()
}

// UnimplementedReservationOrderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReservationOrderServiceServer struct {
}

func (UnimplementedReservationOrderServiceServer) Create(context.Context, *ReservationOrderReq) (*ReservationOrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedReservationOrderServiceServer) Get(context.Context, *GetByIdReq) (*ReservationOrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedReservationOrderServiceServer) GetAll(context.Context, *GetAllReservationOrderReq) (*GetAllReservationOrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedReservationOrderServiceServer) Update(context.Context, *ReservationOrderUpdateReq) (*ReservationOrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedReservationOrderServiceServer) Delete(context.Context, *GetByIdReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedReservationOrderServiceServer) mustEmbedUnimplementedReservationOrderServiceServer() {
}

// UnsafeReservationOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReservationOrderServiceServer will
// result in compilation errors.
type UnsafeReservationOrderServiceServer interface {
	mustEmbedUnimplementedReservationOrderServiceServer()
}

func RegisterReservationOrderServiceServer(s grpc.ServiceRegistrar, srv ReservationOrderServiceServer) {
	s.RegisterService(&ReservationOrderService_ServiceDesc, srv)
}

func _ReservationOrderService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReservationOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationOrderServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationOrderService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationOrderServiceServer).Create(ctx, req.(*ReservationOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationOrderService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationOrderServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationOrderService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationOrderServiceServer).Get(ctx, req.(*GetByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationOrderService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllReservationOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationOrderServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationOrderService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationOrderServiceServer).GetAll(ctx, req.(*GetAllReservationOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationOrderService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReservationOrderUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationOrderServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationOrderService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationOrderServiceServer).Update(ctx, req.(*ReservationOrderUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationOrderService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationOrderServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationOrderService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationOrderServiceServer).Delete(ctx, req.(*GetByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ReservationOrderService_ServiceDesc is the grpc.ServiceDesc for ReservationOrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReservationOrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reservation.ReservationOrderService",
	HandlerType: (*ReservationOrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ReservationOrderService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ReservationOrderService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _ReservationOrderService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ReservationOrderService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ReservationOrderService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reservation_orders.proto",
}
