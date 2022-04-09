package service

import (
	"github.com/google/wire"
	v1 "week04.myapp/api/shop/interface/v1"
	"week04.myapp/app/shop/interface/internal/biz"
)

var ProviderSet = wire.NewSet(NewShopInterface)

type ShopInterface struct {
	v1.UnimplementedShopInterfaceServer

	gc *biz.GoodsUseCase
}

func NewShopInterface(gc *biz.GoodsUseCase) *ShopInterface {
	return &ShopInterface{
		gc: gc,
	}
}
