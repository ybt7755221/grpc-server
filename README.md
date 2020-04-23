# GRPC-Sever

#### 简介

Grpc-server是一个给予gprc封装的rpc框架，包含了基本的CURD操作，同时使用[rpctool](https://github.com/ybt7755221/m2p)生成对应的proto文件

#### 结构
- config 配置文件
- entity 实体
- library 工具库
- model 数据库操作
- protos proto文件
- router 注册rpc服务文件
- service rpc服务文件

#### 使用

    //先执行自动生成脚本
    rpctool --mysql user:password@tcp\(host:port\)/database\?charset=utf8 --dbName database --table tableName --out-file ./
    //执行proto生成脚本
    protoc --go_out=plugins=grpc:. ./protos/tableName/tableName.proto
    //测试启动
    fresh
    go run server.go
    //编译
    go build

#### 工具
【