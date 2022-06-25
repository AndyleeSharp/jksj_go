package service

import (
	"context"
	"github.com/google/wire"
	v1 "week13.myapp/api/rate/service/v1"
	"week13.myapp/app/rate/service/internal/biz"
)

type RateService struct {
	v1.UnimplementedRateServiceServer
	uc *biz.RateUseCase
}

func (rs *RateService) GetGoodsRate(ctx context.Context, req *v1.GetRateRequest) (*v1.GetRateReply, error) {
	rate, err := rs.uc.GetRate(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	reply := &v1.GetRateReply{
		Rate: int32(rate),
	}
	return reply, nil
}

func NewRateService(uc *biz.RateUseCase) *RateService {
	return &RateService{uc: uc}
}

var ProviderSet = wire.NewSet(NewRateService)
