// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0--rc2
// source: v1/goods.proto

package v1

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

// GoodsServiceClient is the client API for GoodsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoodsServiceClient interface {
	GetGoodsInfo(ctx context.Context, in *GetGoodsInfoRequest, opts ...grpc.CallOption) (*GetGoodsInfoReply, error)
}

type goodsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGoodsServiceClient(cc grpc.ClientConnInterface) GoodsServiceClient {
	return &goodsServiceClient{cc}
}

func (c *goodsServiceClient) GetGoodsInfo(ctx context.Context, in *GetGoodsInfoRequest, opts ...grpc.CallOption) (*GetGoodsInfoReply, error) {
	out := new(GetGoodsInfoReply)
	err := c.cc.Invoke(ctx, "/goods.service.v1.GoodsService/GetGoodsInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoodsServiceServer is the server API for GoodsService service.
// All implementations must embed UnimplementedGoodsServiceServer
// for forward compatibility
type GoodsServiceServer interface {
	GetGoodsInfo(context.Context, *GetGoodsInfoRequest) (*GetGoodsInfoReply, error)
	mustEmbedUnimplementedGoodsServiceServer()
}

// UnimplementedGoodsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGoodsServiceServer struct {
}

func (UnimplementedGoodsServiceServer) GetGoodsInfo(context.Context, *GetGoodsInfoRequest) (*GetGoodsInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodsInfo not implemented")
}
func (UnimplementedGoodsServiceServer) mustEmbedUnimplementedGoodsServiceServer() {}

// UnsafeGoodsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoodsServiceServer will
// result in compilation errors.
type UnsafeGoodsServiceServer interface {
	mustEmbedUnimplementedGoodsServiceServer()
}

func RegisterGoodsServiceServer(s grpc.ServiceRegistrar, srv GoodsServiceServer) {
	s.RegisterService(&GoodsService_ServiceDesc, srv)
}

func _GoodsService_GetGoodsInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodsInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServiceServer).GetGoodsInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.service.v1.GoodsService/GetGoodsInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServiceServer).GetGoodsInfo(ctx, req.(*GetGoodsInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GoodsService_ServiceDesc is the grpc.ServiceDesc for GoodsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoodsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "goods.service.v1.GoodsService",
	HandlerType: (*GoodsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGoodsInfo",
			Handler:    _GoodsService_GetGoodsInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/goods.proto",
}
