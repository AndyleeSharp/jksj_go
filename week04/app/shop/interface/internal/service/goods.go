package service

import (
	"context"

	v1 "week04.myapp/api/shop/interface/v1"
)

func (si *ShopInterface) GetGoods(ctx context.Context, req *v1.GetGoodsRequest) (*v1.GetGoodsReply, error) {
	goods, err := si.gc.GetGoods(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	g := &v1.Goods{
		Id:    goods.Id,
		Name:  goods.Name,
		Price: goods.Price,
	}
	reply := &v1.GetGoodsReply{
		Goods: g,
	}
	return reply, nil
}
