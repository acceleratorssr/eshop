package web

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type OrderHandler struct {
}

func (o *OrderHandler) order() {

}

func (o *OrderHandler) RegisterRoutes(h *server.Hertz) {
	hg := h.Group("/order")

	hg.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})

	hg.GET("/list_order", o.ListOrder())
}

func (o *OrderHandler) ListOrder() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		c.JSON(consts.StatusOK, utils.H{"message": "pong"})
	}
}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{}
}
