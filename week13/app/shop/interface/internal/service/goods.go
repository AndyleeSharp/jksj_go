package service

import (
	"context"
	goodsv1 "week13.myapp/api/goods/service/v1"
	shopv1 "week13.myapp/api/shop/interface/v1"
)

func (si *ShopInterface) GetGoods(ctx context.Context, req *shopv1.GetGoodsRequest) (*shopv1.GetGoodsReply, error) {
	goods, rate, err := si.gc.GetGoods(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	g := &goodsv1.Goods{
		Id:    goods.Id,
		Name:  goods.Name,
		Price: goods.Price,
	}
	reply := &shopv1.GetGoodsReply{
		Goods: g,
		Rate:  int32(rate),
	}
	return reply, nil
}
