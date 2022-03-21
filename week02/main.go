package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

//1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
//答：应该Wrap,因为这是调用第三方库产生的error，应该保留原始错误的根因以及堆栈信息供最外层打印，同时附加一些当前的业务错误信息

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
