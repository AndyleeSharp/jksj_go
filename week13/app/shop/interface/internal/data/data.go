package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	"github.com/google/wire"
	goodsv1 "week13.myapp/api/goods/service/v1"
	ratev1 "week13.myapp/api/rate/service/v1"
)

var ProviderSet = wire.NewSet(NewData, NewGoodsRepo, NewRateRepo)

type Data struct {
	gc goodsv1.GoodsServiceClient
	rc ratev1.RateServiceClient
}

func NewData() *Data {
	connGoods, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint("localhost:9002"))
	if err != nil {
		panic(err)
	}
	gc := goodsv1.NewGoodsServiceClient(connGoods)

	connRate, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint("localhost:9001"))
	if err != nil {
		panic(err)
	}
	rc := ratev1.NewRateServiceClient(connRate)
	return &Data{gc: gc, rc: rc}
}
