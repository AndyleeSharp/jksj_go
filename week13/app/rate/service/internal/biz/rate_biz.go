package biz

import (
	"context"
	"github.com/google/wire"
)

type RateRepo interface {
	GetRate(ctx context.Context, id int64) (int, error)
}

type RateUseCase struct {
	repo RateRepo
}

func (uc *RateUseCase) GetRate(ctx context.Context, id int64) (int, error) {
	return uc.repo.GetRate(ctx, id)
}

func NewRateUseCase(repo RateRepo) *RateUseCase {
	return &RateUseCase{repo: repo}
}

var ProviderSet = wire.NewSet(NewRateUseCase)
