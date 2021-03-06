// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"week13.myapp/app/rate/service/internal/biz"
	"week13.myapp/app/rate/service/internal/data"
	"week13.myapp/app/rate/service/internal/server"
	"week13.myapp/app/rate/service/internal/service"
)

// Injectors from wire.go:

func initApp() (*kratos.App, error) {
	rateRepo := data.NewRateRepo()
	rateUseCase := biz.NewRateUseCase(rateRepo)
	rateService := service.NewRateService(rateUseCase)
	grpcServer := server.NewGRPCServer(rateService)
	app := newApp(grpcServer)
	return app, nil
}
