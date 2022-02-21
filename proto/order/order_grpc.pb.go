// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: order.proto

package orderService

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

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	CreateOrder(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*CreateOrderRes, error)
	PayOrder(ctx context.Context, in *PayOrderReq, opts ...grpc.CallOption) (*PayOrderRes, error)
	SubmitOrder(ctx context.Context, in *SubmitOrderReq, opts ...grpc.CallOption) (*SubmitOrderRes, error)
	UpdateOrder(ctx context.Context, in *UpdateOrderReq, opts ...grpc.CallOption) (*UpdateOrderRes, error)
	CancelOrder(ctx context.Context, in *CancelOrderReq, opts ...grpc.CallOption) (*CancelOrderRes, error)
	CompleteOrder(ctx context.Context, in *CompleteOrderReq, opts ...grpc.CallOption) (*CompleteOrderRes, error)
	ChangeDeliveryAddress(ctx context.Context, in *ChangeDeliveryAddressReq, opts ...grpc.CallOption) (*ChangeDeliveryAddressRes, error)
	GetOrderByID(ctx context.Context, in *GetOrderByIDReq, opts ...grpc.CallOption) (*GetOrderByIDRes, error)
	Search(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (*SearchRes, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) CreateOrder(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*CreateOrderRes, error) {
	out := new(CreateOrderRes)
	err := c.cc.Invoke(ctx, "/orderService.orderService/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) PayOrder(ctx context.Context, in *PayOrderReq, opts ...grpc.CallOption) (*PayOrderRes, error) {
	out := new(PayOrderRes)
	err := c.cc.Invoke(ctx, "/orderService.orderService/PayOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) SubmitOrder(ctx context.Context, in *SubmitOrderReq, opts ...grpc.CallOption) (*SubmitOrderRes, error) {
	out := new(SubmitOrderRes)
	err := c.cc.Invoke(ctx, "/orderService.orderService/SubmitOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) UpdateOrder(ctx context.Context, in *UpdateOrderReq, opts ...grpc.CallOption) (*UpdateOrderRes, error) {
	out := new(UpdateOrderRes)
	err := c.cc.Invoke(ctx, "/orderService.orderService/UpdateShoppingCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CancelOrder(ctx context.Context, in *CancelOrderReq, opts ...grpc.CallOption) (*CancelOrderRes, error) {
	out := new(CancelOrderRes)
	err := c.cc.Invoke(ctx, "/orderService.orderService/CancelOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CompleteOrder(ctx context.Context, in *CompleteOrderReq, opts ...grpc.CallOption) (*CompleteOrderRes, error) {
	out := new(CompleteOrderRes)
	err := c.cc.Invoke(ctx, "/orderService.orderService/CompleteOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) ChangeDeliveryAddress(ctx context.Context, in *ChangeDeliveryAddressReq, opts ...grpc.CallOption) (*ChangeDeliveryAddressRes, error) {
	out := new(ChangeDeliveryAddressRes)
	err := c.cc.Invoke(ctx, "/orderService.orderService/ChangeDeliveryAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrderByID(ctx context.Context, in *GetOrderByIDReq, opts ...grpc.CallOption) (*GetOrderByIDRes, error) {
	out := new(GetOrderByIDRes)
	err := c.cc.Invoke(ctx, "/orderService.orderService/GetOrderByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) Search(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (*SearchRes, error) {
	out := new(SearchRes)
	err := c.cc.Invoke(ctx, "/orderService.orderService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations should embed UnimplementedOrderServiceServer
// for forward compatibility
type OrderServiceServer interface {
	CreateOrder(context.Context, *CreateOrderReq) (*CreateOrderRes, error)
	PayOrder(context.Context, *PayOrderReq) (*PayOrderRes, error)
	SubmitOrder(context.Context, *SubmitOrderReq) (*SubmitOrderRes, error)
	UpdateOrder(context.Context, *UpdateOrderReq) (*UpdateOrderRes, error)
	CancelOrder(context.Context, *CancelOrderReq) (*CancelOrderRes, error)
	CompleteOrder(context.Context, *CompleteOrderReq) (*CompleteOrderRes, error)
	ChangeDeliveryAddress(context.Context, *ChangeDeliveryAddressReq) (*ChangeDeliveryAddressRes, error)
	GetOrderByID(context.Context, *GetOrderByIDReq) (*GetOrderByIDRes, error)
	Search(context.Context, *SearchReq) (*SearchRes, error)
}

// UnimplementedOrderServiceServer should be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (UnimplementedOrderServiceServer) CreateOrder(context.Context, *CreateOrderReq) (*CreateOrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderServiceServer) PayOrder(context.Context, *PayOrderReq) (*PayOrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayOrder not implemented")
}
func (UnimplementedOrderServiceServer) SubmitOrder(context.Context, *SubmitOrderReq) (*SubmitOrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitOrder not implemented")
}
func (UnimplementedOrderServiceServer) UpdateOrder(context.Context, *UpdateOrderReq) (*UpdateOrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateShoppingCart not implemented")
}
func (UnimplementedOrderServiceServer) CancelOrder(context.Context, *CancelOrderReq) (*CancelOrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}
func (UnimplementedOrderServiceServer) CompleteOrder(context.Context, *CompleteOrderReq) (*CompleteOrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteOrder not implemented")
}
func (UnimplementedOrderServiceServer) ChangeDeliveryAddress(context.Context, *ChangeDeliveryAddressReq) (*ChangeDeliveryAddressRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeDeliveryAddress not implemented")
}
func (UnimplementedOrderServiceServer) GetOrderByID(context.Context, *GetOrderByIDReq) (*GetOrderByIDRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderByID not implemented")
}
func (UnimplementedOrderServiceServer) Search(context.Context, *SearchReq) (*SearchRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderService.orderService/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateOrder(ctx, req.(*CreateOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_PayOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).PayOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderService.orderService/PayOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).PayOrder(ctx, req.(*PayOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_SubmitOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).SubmitOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderService.orderService/SubmitOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).SubmitOrder(ctx, req.(*SubmitOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_UpdateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).UpdateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderService.orderService/UpdateShoppingCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).UpdateOrder(ctx, req.(*UpdateOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderService.orderService/CancelOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CancelOrder(ctx, req.(*CancelOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CompleteOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CompleteOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderService.orderService/CompleteOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CompleteOrder(ctx, req.(*CompleteOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_ChangeDeliveryAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeDeliveryAddressReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).ChangeDeliveryAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderService.orderService/ChangeDeliveryAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).ChangeDeliveryAddress(ctx, req.(*ChangeDeliveryAddressReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrderByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderByIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrderByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderService.orderService/GetOrderByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrderByID(ctx, req.(*GetOrderByIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderService.orderService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).Search(ctx, req.(*SearchReq))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orderService.orderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _OrderService_CreateOrder_Handler,
		},
		{
			MethodName: "PayOrder",
			Handler:    _OrderService_PayOrder_Handler,
		},
		{
			MethodName: "SubmitOrder",
			Handler:    _OrderService_SubmitOrder_Handler,
		},
		{
			MethodName: "UpdateShoppingCart",
			Handler:    _OrderService_UpdateOrder_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _OrderService_CancelOrder_Handler,
		},
		{
			MethodName: "CompleteOrder",
			Handler:    _OrderService_CompleteOrder_Handler,
		},
		{
			MethodName: "ChangeDeliveryAddress",
			Handler:    _OrderService_ChangeDeliveryAddress_Handler,
		},
		{
			MethodName: "GetOrderByID",
			Handler:    _OrderService_GetOrderByID_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _OrderService_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}
