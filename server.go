package main

import (
	"grpc-server/router"
	"log"
	"net"
)

const TCP_ADDR = ":8028"

func main() {
	lis, err := net.Listen("tcp", TCP_ADDR) //监听所有网卡8028端口的TCP连接
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}
	//初始化server
	s := router.InitRouter()
	//启动server
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
