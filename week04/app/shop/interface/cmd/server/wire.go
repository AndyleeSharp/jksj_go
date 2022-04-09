//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"week04.myapp/app/shop/interface/internal/biz"
	"week04.myapp/app/shop/interface/internal/data"
	"week04.myapp/app/shop/interface/internal/server"
	"week04.myapp/app/shop/interface/internal/service"
)

func initApp() (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
