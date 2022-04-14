//go:generate wire
//go:build wireinject
// +build wireinject

package main

import (
	"usm/internal/biz"
	"usm/internal/conf"
	"usm/internal/data"
	"usm/internal/server"
	"usm/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
