package service

import (
	"context"
	"github.com/google/wire"
	v1 "week13.myapp/api/goods/service/v1"
	"week13.myapp/app/goods/service/internal/biz"
)

type GoodsService struct {
	v1.UnimplementedGoodsServiceServer
	uc *biz.GoodsUseCase
}

func (gs *GoodsService) GetGoodsInfo(ctx context.Context, req *v1.GetGoodsInfoRequest) (*v1.GetGoodsInfoReply, error) {
	goods, err := gs.uc.GetGoods(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	g := &v1.Goods{
		Id:    goods.Id,
		Name:  goods.Name,
		Price: goods.Price,
	}
	reply := &v1.GetGoodsInfoReply{Goods: g}
	return reply, nil

}

func NewGoodsService(uc *biz.GoodsUseCase) *GoodsService {
	return &GoodsService{uc: uc}
}

var ProviderSet = wire.NewSet(NewGoodsService)
