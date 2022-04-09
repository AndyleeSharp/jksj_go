package data

import (
	"context"

	"week04.myapp/app/shop/interface/internal/biz"
)

var _ biz.GoodsRepo = (*goodsRepo)(nil)

type goodsRepo struct {
	data *Data
}

func (r *goodsRepo) GetGoods(ctx context.Context, id int64) (*biz.Goods, error) {
	return r.data.db.GetGoods(ctx, id)
}

func NewGoodsRepo(data *Data) biz.GoodsRepo {
	return &goodsRepo{data: data}
}
