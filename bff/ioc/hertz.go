package ioc

import (
	"eshop/bff/web"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func InitGinServer(order *web.OrderHandler) *server.Hertz {
	h := server.New(server.WithHostPorts("127.0.0.1:9190")) // 注意 localhost

	order.RegisterRoutes(h)

	return h
}

// 跨域
func corsHdl() {
}
