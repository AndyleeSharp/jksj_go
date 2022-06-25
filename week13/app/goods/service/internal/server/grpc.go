package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "week13.myapp/api/goods/service/v1"
	"week13.myapp/app/goods/service/internal/service"
)

func NewGRPCServer(s *service.GoodsService) *grpc.Server {
	var opts []grpc.ServerOption
	opts = append(opts, grpc.Address("0.0.0.0:9002"))
	svr := grpc.NewServer(opts...)
	v1.RegisterGoodsServiceServer(svr, s)

	return svr
}
