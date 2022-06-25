//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"week13.myapp/app/rate/service/internal/biz"
	"week13.myapp/app/rate/service/internal/data"
	"week13.myapp/app/rate/service/internal/server"
	"week13.myapp/app/rate/service/internal/service"
)

func initApp() (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
