package data

import (
	"context"
	"github.com/google/wire"
	"week13.myapp/app/rate/service/internal/biz"
)

type Data struct {
}

type rateRepo struct {
	data *Data
}

func (d *rateRepo) GetRate(ctx context.Context, id int64) (int, error) {
	return (int(id) * 2) % 100, nil
}

func NewRateRepo() biz.RateRepo {
	repo := &rateRepo{}
	return repo
}

var ProviderSet = wire.NewSet(NewRateRepo)
