// Package wrappers ..
/*
 * @Descripttion:go-micro中间件
 * @Author: congz
 * @Date: 2020-09-15 22:22:40
 * @LastEditors: congz
 * @LastEditTime: 2020-10-27 13:00:15
 */
package wrappers

import (
	"api-gateway/services"
	"context"
	"strconv"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/client"
)

//NewProduct 新建商品信息
func NewProduct(id uint32, name string) *services.ProductModel {
	return &services.ProductModel{
		ID:            id,
		Name:          name,
		CategoryID:    1,
		Info:          "响应超时",
		Title:         "响应超时",
		ImgPath:       "https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/0099822e42b4428cb25c4cdebc6ca53d.jpg?thumb=1&w=200&h=200&f=webp&q=90",
		Price:         "1000",
		DiscountPrice: "1000"}
}

//DefaultProducts 降级处理函数，固定生成8个商品信息
func DefaultProducts(rsp interface{}) {
	models := make([]*services.ProductModel, 0)
	var i uint32
	for i = 0; i < 8; i++ {
		models = append(models, NewProduct(i, "降级商品"+strconv.Itoa(20+int(i))))
	}
	result := rsp.(*services.ProductsListResponse)
	result.ProductsList = models
}

//DefaultData 通用降级方法
func DefaultData(rsp interface{}) {
	switch t := rsp.(type) {
	case *services.ProductsListResponse:
		DefaultProducts(rsp)
	case *services.ProductDetailResponse:
		t.ProductDetail = NewProduct(10, "降级商品")
	}
}

//ProductWrapper go-micro Wrapper中间件
type ProductWrapper struct {
	client.Client
}

//Call Wrapper中间件的执行方法
func (wrapper *ProductWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	cmdName := req.Service() + "." + req.Endpoint()
	config := hystrix.CommandConfig{
		Timeout:                3000,
		RequestVolumeThreshold: 20,   //熔断器请求阈值，默认20，意思是有20个请求才能进行错误百分比计算
		ErrorPercentThreshold:  50,   //错误百分比，当错误超过百分比时，直接进行降级处理，直至熔断器再次 开启，默认50%
		SleepWindow:            5000, //过多长时间，熔断器再次检测是否开启，单位毫秒ms（默认5秒）
	}
	hystrix.ConfigureCommand(cmdName, config)
	return hystrix.Do(cmdName, func() error {
		return wrapper.Client.Call(ctx, req, rsp)
	}, func(err error) error {
		//降级处理
		//DefaultProducts(rsp)
		//DefaultData(rsp)
		return err
	})
}

//NewProductWrapper 初始化Wrapper
func NewProductWrapper(c client.Client) client.Client {
	return &ProductWrapper{c}
}
