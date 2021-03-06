# 微服务后台管理系统

## 前言

本人正在读大三，最近接触到了微服务，所以尝试使用微服务来实现这个项目，如果有错误或者实现不好的地方欢迎 issues

如果觉得这个项目不错，您可以右上角 Star 支持一下！谢谢您的支持，您的支持是我完善的动力！

## 项目简介

使用 go-micro 实现的微服务后台管理系统，负责管理[CMall http://cmall.congz.top/#/](http://cmall.congz.top/#/)的用户、商品等数据

api-gateway(网关)是使用 gin 实现，包含暴露 HTTP、负载均衡、熔断和降级处理、token 验证等基本功能，网关和各模块间使用 rpc 通信

根据功能复杂程度将微服务拆分成三个模块，分别是用户模块、商品模块、其他模块

目前使用 consul 来进行服务的注册与发现，后续根据情况可能会转用 etcd

## 项目进度

后端：完成了各模块的基本功能,在创建商品功能使用了消息队列

前端：使用 vue 实现了简单的界面

## 项目依赖

本项目采用了一系列 golang 中比较流行的组件来进行开发

- gin
- gorm
- go-micro
- mysql
- godotenv
- jwt-go
- protobuf
- rabbitMQ

## 目录结构

```
backstage/
├── api-gateway
│	├── pkg
│	├── services
│	├── weblib
│	├── wrappers
├── mq-server
│	├── conf
│	├── model
│	├── service
├── other
│	├── conf
│	├── model
│	├── serviceimpl
│	├── services
├── product
└── user

```

- pkg：工具函数以及错误码封装
- services：放置 proto 和 pb.go 文件
- weblib：放置使用 gin 处理的文件
- wrappers：go-micro 中间件
- conf：配置文件
- model：应用数据库模型
- serviceimpl：放置各服务的处理函数

## 运行

本项目使用[Go Mod](https://github.com/golang/go/wiki/Modules)管理依赖。

每个模块都需要运行

`go run main.go`

项目运行后启动在 4000，4001，4002，4003 端口
