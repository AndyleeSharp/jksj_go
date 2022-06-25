package data

import (
	"context"
	ratev1 "week13.myapp/api/rate/service/v1"

	"week13.myapp/app/shop/interface/internal/biz"
)

var _ biz.RateRepo = (*rateRepo)(nil)

type rateRepo struct {
	data *Data
}

func (r *rateRepo) GetRate(ctx context.Context, id int64) (int, error) {
	reply, err := r.data.rc.GetGoodsRate(ctx, &ratev1.GetRateRequest{Id: id})
	if err != nil {
		return -1, err
	}
	return int(reply.Rate), err
}

func NewRateRepo(data *Data) biz.RateRepo {
	return &rateRepo{data: data}
}
