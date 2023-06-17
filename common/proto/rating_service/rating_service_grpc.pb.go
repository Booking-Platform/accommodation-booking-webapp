// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.6
// source: rating_service.proto

package rating

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
	RatingService_CreateRating_FullMethodName                             = "/rating.rating_service/CreateRating"
	RatingService_CreateAccommodationRating_FullMethodName                = "/rating.rating_service/CreateAccommodationRating"
	RatingService_GetAccommodationRatingsByAccommodationID_FullMethodName = "/rating.rating_service/GetAccommodationRatingsByAccommodationID"
)

// RatingServiceClient is the client API for RatingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RatingServiceClient interface {
	CreateRating(ctx context.Context, in *CreateRatingRequest, opts ...grpc.CallOption) (*BlankResponse, error)
	CreateAccommodationRating(ctx context.Context, in *CreateAccommodationRatingRequest, opts ...grpc.CallOption) (*BlankResponse, error)
	GetAccommodationRatingsByAccommodationID(ctx context.Context, in *IdMessageRequest, opts ...grpc.CallOption) (*RatingsResponse, error)
}

type ratingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRatingServiceClient(cc grpc.ClientConnInterface) RatingServiceClient {
	return &ratingServiceClient{cc}
}

func (c *ratingServiceClient) CreateRating(ctx context.Context, in *CreateRatingRequest, opts ...grpc.CallOption) (*BlankResponse, error) {
	out := new(BlankResponse)
	err := c.cc.Invoke(ctx, RatingService_CreateRating_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) CreateAccommodationRating(ctx context.Context, in *CreateAccommodationRatingRequest, opts ...grpc.CallOption) (*BlankResponse, error) {
	out := new(BlankResponse)
	err := c.cc.Invoke(ctx, RatingService_CreateAccommodationRating_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetAccommodationRatingsByAccommodationID(ctx context.Context, in *IdMessageRequest, opts ...grpc.CallOption) (*RatingsResponse, error) {
	out := new(RatingsResponse)
	err := c.cc.Invoke(ctx, RatingService_GetAccommodationRatingsByAccommodationID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RatingServiceServer is the server API for RatingService service.
// All implementations must embed UnimplementedRatingServiceServer
// for forward compatibility
type RatingServiceServer interface {
	CreateRating(context.Context, *CreateRatingRequest) (*BlankResponse, error)
	CreateAccommodationRating(context.Context, *CreateAccommodationRatingRequest) (*BlankResponse, error)
	GetAccommodationRatingsByAccommodationID(context.Context, *IdMessageRequest) (*RatingsResponse, error)
	mustEmbedUnimplementedRatingServiceServer()
}

// UnimplementedRatingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRatingServiceServer struct {
}

func (UnimplementedRatingServiceServer) CreateRating(context.Context, *CreateRatingRequest) (*BlankResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRating not implemented")
}
func (UnimplementedRatingServiceServer) CreateAccommodationRating(context.Context, *CreateAccommodationRatingRequest) (*BlankResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccommodationRating not implemented")
}
func (UnimplementedRatingServiceServer) GetAccommodationRatingsByAccommodationID(context.Context, *IdMessageRequest) (*RatingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccommodationRatingsByAccommodationID not implemented")
}
func (UnimplementedRatingServiceServer) mustEmbedUnimplementedRatingServiceServer() {}

// UnsafeRatingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RatingServiceServer will
// result in compilation errors.
type UnsafeRatingServiceServer interface {
	mustEmbedUnimplementedRatingServiceServer()
}

func RegisterRatingServiceServer(s grpc.ServiceRegistrar, srv RatingServiceServer) {
	s.RegisterService(&RatingService_ServiceDesc, srv)
}

func _RatingService_CreateRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).CreateRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_CreateRating_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).CreateRating(ctx, req.(*CreateRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_CreateAccommodationRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccommodationRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).CreateAccommodationRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_CreateAccommodationRating_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).CreateAccommodationRating(ctx, req.(*CreateAccommodationRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetAccommodationRatingsByAccommodationID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetAccommodationRatingsByAccommodationID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_GetAccommodationRatingsByAccommodationID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetAccommodationRatingsByAccommodationID(ctx, req.(*IdMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RatingService_ServiceDesc is the grpc.ServiceDesc for RatingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RatingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rating.rating_service",
	HandlerType: (*RatingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRating",
			Handler:    _RatingService_CreateRating_Handler,
		},
		{
			MethodName: "CreateAccommodationRating",
			Handler:    _RatingService_CreateAccommodationRating_Handler,
		},
		{
			MethodName: "GetAccommodationRatingsByAccommodationID",
			Handler:    _RatingService_GetAccommodationRatingsByAccommodationID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rating_service.proto",
}
