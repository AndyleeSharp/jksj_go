package biz

import "context"

type Goods struct {
	Id    int64
	Name  string
	Price float64
}
type GoodsRepo interface {
	GetGoods(ctx context.Context, id int64) (*Goods, error)
}
type GoodsUseCase struct {
	repo GoodsRepo
}

func (uc *GoodsUseCase) GetGoods(ctx context.Context, id int64) (*Goods, error) {
	return uc.repo.GetGoods(ctx, id)
}

func NewGoodsUseCase(repo GoodsRepo) *GoodsUseCase {
	return &GoodsUseCase{repo: repo}
}
