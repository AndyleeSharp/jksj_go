package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

type Player struct {
	Id   int32
	Name string
}

func GetMysqlError() error {
	return fmt.Errorf("sql.ErrNoRows")
}

type IDao interface {
	FindPlayerById(id int32) (player *Player, err error)
}

type MockDao struct {
}

func (dao *MockDao) FindPlayerById(id int32) (player *Player, err error) {
	if id == 0 {
		err := GetMysqlError()
		return nil, errors.Wrap(err, "can't find player with ID 0")
	}
	var p = Player{Id: id, Name: fmt.Sprintf("name_%d", id)}
	return &p, nil
}

func main() {
	var dao IDao = &MockDao{}
	p1, err := dao.FindPlayerById(1)
	if err == nil {
		fmt.Printf("find player: %v\n", p1)
	}
	_, err = dao.FindPlayerById(0)
	if err != nil {
		fmt.Printf("original error: %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("statck trace:\n%+v\n", err)
		os.Exit(1)
	}

}
