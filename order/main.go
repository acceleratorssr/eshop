package main

import (
	order "eshop/api/kitex_gen/order/orderservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"net"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:9091") // 指定监听地址
	if err != nil {
		klog.Fatalf("Failed to resolve address: %v", err)
	}

	svr := order.NewServer(
		new(OrderServiceImpl),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "order",
		}),
	)

	err = svr.Run()
	if err != nil {
		klog.Fatalf("Server failed to start: %v", err)
	}
}
