package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "week13.myapp/api/rate/service/v1"
	"week13.myapp/app/rate/service/internal/service"
)

func NewGRPCServer(s *service.RateService) *grpc.Server {
	var opts []grpc.ServerOption
	opts = append(opts, grpc.Address("0.0.0.0:9001"))
	svr := grpc.NewServer(opts...)
	v1.RegisterRateServiceServer(svr, s)
	return svr
}
