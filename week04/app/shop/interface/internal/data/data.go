package data

import (
	"context"

	"github.com/google/wire"
	"week04.myapp/app/shop/interface/internal/biz"
)

var ProviderSet = wire.NewSet(NewData, NewGoodsRepo, NewFakeDB)

type FakeDB struct {
}

type Data struct {
	db *FakeDB
}

func NewData(db *FakeDB) *Data {
	return &Data{db: db}
}

func NewFakeDB() *FakeDB {
	return &FakeDB{}
}

func (d *FakeDB) GetGoods(ctx context.Context, id int64) (*biz.Goods, error) {
	return &biz.Goods{Id: id, Name: "goods", Price: 188}, nil
}
