package main

import (
	"fmt"
	"syscall"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func newApp(hs *http.Server, gs *grpc.Server) *kratos.App {
	return kratos.New(
		kratos.Name("week04"),
		kratos.Server(
			hs,
			gs,
		),
		kratos.Signal(syscall.SIGINT, syscall.SIGTERM),
	)
}
func main() {
	var a []int
	fmt.Println(a == nil)
	a = append(a, 1)
	fmt.Println(a)

	app, err := initApp()
	if err != nil {
		panic(err)
	}
	if err := app.Run(); err != nil {
		panic(err)
	}
}
