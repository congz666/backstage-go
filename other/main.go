// package main
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-09-15 11:09:52
 * @LastEditors: congz
 * @LastEditTime: 2020-09-26 19:06:32
 */
package main

import (
	"other/conf"
	"other/serviceimpl"
	"other/services"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	conf.Init()
	//consul注册件
	consulReg := consul.NewRegistry(
		registry.Addrs("localhost:8500"),
	)
	//创建微服务实例，并注册到consul
	microService := micro.NewService(
		micro.Name("rpcOtherService"),
		micro.Address(":4002"),
		micro.Registry(consulReg),
	)
	//接收命令行参数 如--server_address
	microService.Init()

	//注册商品服务
	services.RegisterOtherServiceHandler(microService.Server(), new(serviceimpl.OtherService))
	microService.Run()
}
