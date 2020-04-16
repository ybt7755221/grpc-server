package router

import (
	"grpc-server/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func InitRouter() *grpc.Server {
	grpcServer := grpc.NewServer() //创建gRPC服务
	//注册rpc服务
	userServ := service.UserServer{}
	userServ.Register(grpcServer)
	// 在gRPC服务器上注册反射服务
	reflection.Register(grpcServer)
	return grpcServer
}
