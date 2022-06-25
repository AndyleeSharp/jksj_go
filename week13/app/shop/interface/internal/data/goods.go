package data

import (
	"context"
	goodsv1 "week13.myapp/api/goods/service/v1"

	"week13.myapp/app/shop/interface/internal/biz"
)

var _ biz.GoodsRepo = (*goodsRepo)(nil)

type goodsRepo struct {
	data *Data
}

func (r *goodsRepo) GetGoods(ctx context.Context, id int64) (*biz.Goods, error) {
	reply, err := r.data.gc.GetGoodsInfo(ctx, &goodsv1.GetGoodsInfoRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return &biz.Goods{
		Id:    reply.Goods.Id,
		Name:  reply.Goods.Name,
		Price: reply.Goods.Price,
	}, err
}

func NewGoodsRepo(data *Data) biz.GoodsRepo {
	return &goodsRepo{data: data}
}
