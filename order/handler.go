package main

import (
	"context"
	"eshop/api/kitex_gen/order"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// PlaceOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) PlaceOrder(ctx context.Context, req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	// TODO: Your code here...
	return
}

// ListOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) ListOrder(ctx context.Context, req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	// TODO: Your code here...
	return
}

// ChangeOrderStatus implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) ChangeOrderStatus(ctx context.Context, req *order.ChangeOrderStatusReq) (resp *order.ChangeOrderStatusResp, err error) {
	// TODO: Your code here...
	return
}
