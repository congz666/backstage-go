//Package main ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-09-20 11:30:37
 * @LastEditors: congz
 * @LastEditTime: 2020-10-27 14:04:58
 */
package main

import (
	"api-gateway/services"
	"api-gateway/weblib"
	"api-gateway/wrappers"

	"github.com/joho/godotenv"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	//从本地读取env文件
	godotenv.Load()
	//consul注册件
	consulReq := consul.NewRegistry(
		registry.Addrs("localhost:8500"),
	)

	productMicroService := micro.NewService(
		micro.Name("productService.client"),
		micro.WrapClient(wrappers.NewProductWrapper),
	)
	//商品服务调用实例
	productService := services.NewProductService("rpcProductService", productMicroService.Client())

	otherMicroService := micro.NewService(
		micro.Name("carouselService.client"),
		micro.WrapClient(wrappers.NewOtherWrapper),
	)
	//轮播图服务调用实例
	otherService := services.NewOtherService("rpcOtherService", otherMicroService.Client())

	userMicroService := micro.NewService(
		micro.Name("userService.client"),
		micro.WrapClient(wrappers.NewUserWrapper),
	)
	//用户服务调用实例
	userService := services.NewUserService("rpcUserService", userMicroService.Client())
	//创建微服务实例，使用gin暴露http接口并注册到consul
	server := web.NewService(
		web.Name("httpService"),
		web.Address(":4000"),
		//将服务调用实例使用gin处理
		web.Handler(weblib.NewRouter(productService, otherService, userService)),
		web.Registry(consulReq),
	)
	//接收命令行参数
	server.Init()
	server.Run()
}
