//go:build wireinject

package main

import (
	"eshop/bff/ioc"
	"eshop/bff/web"
	"github.com/google/wire"
)

func InitApp() *App {
	wire.Build(
		ioc.InitGinServer,

		web.NewOrderHandler,

		wire.Struct(new(App), "*"),
	)
	return new(App)
}
