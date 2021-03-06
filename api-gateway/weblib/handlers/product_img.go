// Package handlers ..
/*
 * @Descripttion:商品图片处理函数
 * @Author: congz
 * @Date: 2020-09-20 11:33:37
 * @LastEditors: congz
 * @LastEditTime: 2020-10-27 13:22:21
 */
package handlers

import (
	"api-gateway/services"
	"context"

	"github.com/gin-gonic/gin"
)

//CreateProductImg 使用rpc创建商品图片
func CreateProductImg(ginCtx *gin.Context) {
	var productImgReq services.ProductImgRequest
	PanicIfProductError(ginCtx.Bind(&productImgReq))
	//从gin.keys取出服务实例
	productService := ginCtx.Keys["productService"].(services.ProductService)
	//调用服务端的函数
	productImgRes, err := productService.CreateProductImg(context.Background(), &productImgReq)
	PanicIfProductError(err)
	ginCtx.JSON(200, gin.H{"data": productImgRes.ProductImgDetail})
}

//GetProductImgsList 使用rpc获取商品图片列表
func GetProductImgsList(ginCtx *gin.Context) {
	var productImgReq services.ProductImgRequest
	PanicIfProductError(ginCtx.Bind(&productImgReq))
	//从gin.keys取出服务实例
	productService := ginCtx.Keys["productService"].(services.ProductService)
	//调用服务端的函数
	productImgRes, err := productService.GetProductImgsList(context.Background(), &productImgReq)
	PanicIfProductError(err)
	ginCtx.JSON(200, gin.H{"data": gin.H{"product_img": productImgRes.ProductImgsList, "count": productImgRes.Count}})
}

//UpdateProductImg 使用rpc更新商品图片
func UpdateProductImg(ginCtx *gin.Context) {
	var productImgReq services.ProductImgRequest
	PanicIfProductError(ginCtx.Bind(&productImgReq))
	//从gin.keys取出服务实例
	productService := ginCtx.Keys["productService"].(services.ProductService)
	//调用服务端的函数
	productImgRes, err := productService.UpdateProductImg(context.Background(), &productImgReq)
	PanicIfProductError(err)
	ginCtx.JSON(200, gin.H{"data": productImgRes.ProductImgDetail})
}

//DeleteProductImg 使用rpc删除商品图片
func DeleteProductImg(ginCtx *gin.Context) {
	var productImgReq services.ProductImgRequest
	PanicIfProductError(ginCtx.Bind(&productImgReq))
	//从gin.keys取出服务实例
	productService := ginCtx.Keys["productService"].(services.ProductService)
	//调用服务端的函数
	productImgRes, err := productService.DeleteProductImg(context.Background(), &productImgReq)
	PanicIfProductError(err)
	ginCtx.JSON(200, gin.H{"data": productImgRes.ProductImgDetail})
}
