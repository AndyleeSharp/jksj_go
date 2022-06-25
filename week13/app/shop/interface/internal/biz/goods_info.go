package biz

import (
	"context"
	"golang.org/x/sync/errgroup"
)

type Goods struct {
	Id    int64
	Name  string
	Price float64
}
type GoodsRepo interface {
	GetGoods(ctx context.Context, id int64) (*Goods, error)
}
type RateRepo interface {
	GetRate(ctx context.Context, id int64) (int, error)
}
type GoodsUseCase struct {
	goodsRepo GoodsRepo
	rateRepo  RateRepo
}

func (uc *GoodsUseCase) GetGoods(ctx context.Context, id int64) (*Goods, int, error) {
	group, ctx := errgroup.WithContext(ctx)
	var goods *Goods
	var rate int
	group.Go(func() error {
		var err error
		goods, err = uc.goodsRepo.GetGoods(ctx, id)
		return err
	})

	group.Go(func() error {
		var err error
		rate, err = uc.rateRepo.GetRate(ctx, id)
		return err
	})

	err := group.Wait()
	return goods, rate, err
}

func NewGoodsUseCase(goodsRepo GoodsRepo, rateRepo RateRepo) *GoodsUseCase {
	return &GoodsUseCase{goodsRepo: goodsRepo, rateRepo: rateRepo}
}
