package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "week13.myapp/api/shop/interface/v1"
	"week13.myapp/app/shop/interface/internal/service"
)

func NewGRPCServer(s *service.ShopInterface) *grpc.Server {
	var opts = []grpc.ServerOption{}
	srv := grpc.NewServer(opts...)
	v1.RegisterShopInterfaceServer(srv, s)
	return srv
}
