package data

import (
	"context"
	"fmt"
	"github.com/google/wire"
	"week13.myapp/app/goods/service/internal/biz"
)

type Data struct {
}

type goodsRepo struct {
	data *Data
}

func (d *goodsRepo) GetGoods(ctx context.Context, id int64) (biz.Goods, error) {
	return biz.Goods{
		id,
		fmt.Sprintf("goods_%d", id),
		float64(id * 10),
	}, nil
}

func NewGoodsRepo() biz.GoodsRepo {
	repo := &goodsRepo{}
	return repo
}

var ProviderSet = wire.NewSet(NewGoodsRepo)
