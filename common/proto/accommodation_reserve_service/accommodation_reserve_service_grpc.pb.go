// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.6
// source: accommodation_reserve_service.proto

package accommodation_reserve

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
	AccommodationReserveService_CreateReservation_FullMethodName       = "/accommodation_reserve.accommodation_reserve_service/CreateReservation"
	AccommodationReserveService_GetAllForConfirmation_FullMethodName   = "/accommodation_reserve.accommodation_reserve_service/GetAllForConfirmation"
	AccommodationReserveService_GetReservationsByUserID_FullMethodName = "/accommodation_reserve.accommodation_reserve_service/GetReservationsByUserID"
	AccommodationReserveService_CancelReservation_FullMethodName       = "/accommodation_reserve.accommodation_reserve_service/CancelReservation"
)

// AccommodationReserveServiceClient is the client API for AccommodationReserveService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccommodationReserveServiceClient interface {
	CreateReservation(ctx context.Context, in *CreateReservationRequest, opts ...grpc.CallOption) (*CreateReservationResponse, error)
	GetAllForConfirmation(ctx context.Context, in *GetAllForConfirmationRequest, opts ...grpc.CallOption) (*GetAllForConfirmationResponse, error)
	GetReservationsByUserID(ctx context.Context, in *IdMessageRequest, opts ...grpc.CallOption) (*GetReservationsByUserIDResponse, error)
	CancelReservation(ctx context.Context, in *CancelReservationRequest, opts ...grpc.CallOption) (*CancelReservationResponse, error)
}

type accommodationReserveServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccommodationReserveServiceClient(cc grpc.ClientConnInterface) AccommodationReserveServiceClient {
	return &accommodationReserveServiceClient{cc}
}

func (c *accommodationReserveServiceClient) CreateReservation(ctx context.Context, in *CreateReservationRequest, opts ...grpc.CallOption) (*CreateReservationResponse, error) {
	out := new(CreateReservationResponse)
	err := c.cc.Invoke(ctx, AccommodationReserveService_CreateReservation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationReserveServiceClient) GetAllForConfirmation(ctx context.Context, in *GetAllForConfirmationRequest, opts ...grpc.CallOption) (*GetAllForConfirmationResponse, error) {
	out := new(GetAllForConfirmationResponse)
	err := c.cc.Invoke(ctx, AccommodationReserveService_GetAllForConfirmation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationReserveServiceClient) GetReservationsByUserID(ctx context.Context, in *IdMessageRequest, opts ...grpc.CallOption) (*GetReservationsByUserIDResponse, error) {
	out := new(GetReservationsByUserIDResponse)
	err := c.cc.Invoke(ctx, AccommodationReserveService_GetReservationsByUserID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationReserveServiceClient) CancelReservation(ctx context.Context, in *CancelReservationRequest, opts ...grpc.CallOption) (*CancelReservationResponse, error) {
	out := new(CancelReservationResponse)
	err := c.cc.Invoke(ctx, AccommodationReserveService_CancelReservation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccommodationReserveServiceServer is the server API for AccommodationReserveService service.
// All implementations must embed UnimplementedAccommodationReserveServiceServer
// for forward compatibility
type AccommodationReserveServiceServer interface {
	CreateReservation(context.Context, *CreateReservationRequest) (*CreateReservationResponse, error)
	GetAllForConfirmation(context.Context, *GetAllForConfirmationRequest) (*GetAllForConfirmationResponse, error)
	GetReservationsByUserID(context.Context, *IdMessageRequest) (*GetReservationsByUserIDResponse, error)
	CancelReservation(context.Context, *CancelReservationRequest) (*CancelReservationResponse, error)
	mustEmbedUnimplementedAccommodationReserveServiceServer()
}

// UnimplementedAccommodationReserveServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccommodationReserveServiceServer struct {
}

func (UnimplementedAccommodationReserveServiceServer) CreateReservation(context.Context, *CreateReservationRequest) (*CreateReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReservation not implemented")
}
func (UnimplementedAccommodationReserveServiceServer) GetAllForConfirmation(context.Context, *GetAllForConfirmationRequest) (*GetAllForConfirmationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllForConfirmation not implemented")
}
func (UnimplementedAccommodationReserveServiceServer) GetReservationsByUserID(context.Context, *IdMessageRequest) (*GetReservationsByUserIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReservationsByUserID not implemented")
}
func (UnimplementedAccommodationReserveServiceServer) CancelReservation(context.Context, *CancelReservationRequest) (*CancelReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelReservation not implemented")
}
func (UnimplementedAccommodationReserveServiceServer) mustEmbedUnimplementedAccommodationReserveServiceServer() {
}

// UnsafeAccommodationReserveServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccommodationReserveServiceServer will
// result in compilation errors.
type UnsafeAccommodationReserveServiceServer interface {
	mustEmbedUnimplementedAccommodationReserveServiceServer()
}

func RegisterAccommodationReserveServiceServer(s grpc.ServiceRegistrar, srv AccommodationReserveServiceServer) {
	s.RegisterService(&AccommodationReserveService_ServiceDesc, srv)
}

func _AccommodationReserveService_CreateReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationReserveServiceServer).CreateReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationReserveService_CreateReservation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationReserveServiceServer).CreateReservation(ctx, req.(*CreateReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationReserveService_GetAllForConfirmation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllForConfirmationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationReserveServiceServer).GetAllForConfirmation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationReserveService_GetAllForConfirmation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationReserveServiceServer).GetAllForConfirmation(ctx, req.(*GetAllForConfirmationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationReserveService_GetReservationsByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationReserveServiceServer).GetReservationsByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationReserveService_GetReservationsByUserID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationReserveServiceServer).GetReservationsByUserID(ctx, req.(*IdMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationReserveService_CancelReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationReserveServiceServer).CancelReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationReserveService_CancelReservation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationReserveServiceServer).CancelReservation(ctx, req.(*CancelReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccommodationReserveService_ServiceDesc is the grpc.ServiceDesc for AccommodationReserveService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccommodationReserveService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accommodation_reserve.accommodation_reserve_service",
	HandlerType: (*AccommodationReserveServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateReservation",
			Handler:    _AccommodationReserveService_CreateReservation_Handler,
		},
		{
			MethodName: "GetAllForConfirmation",
			Handler:    _AccommodationReserveService_GetAllForConfirmation_Handler,
		},
		{
			MethodName: "GetReservationsByUserID",
			Handler:    _AccommodationReserveService_GetReservationsByUserID_Handler,
		},
		{
			MethodName: "CancelReservation",
			Handler:    _AccommodationReserveService_CancelReservation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accommodation_reserve_service.proto",
}
