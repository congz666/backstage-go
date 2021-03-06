// Package handlers ..
/*
 * @Descripttion:处理函数
 * @Author: congz
 * @Date: 2020-09-20 11:33:37
 * @LastEditors: congz
 * @LastEditTime: 2020-10-27 14:09:40
 */
package handlers

import (
	"api-gateway/services"
	"context"

	"github.com/gin-gonic/gin"
)

//CreateCategory 使用rpc创建分类列表
func CreateCategory(ginCtx *gin.Context) {
	var categoryReq services.CategoryRequest
	PanicIfOtherError(ginCtx.Bind(&categoryReq))
	//从gin.keys取出服务实例
	otherService := ginCtx.Keys["otherService"].(services.OtherService)
	//调用服务端的函数
	categoryRes, err := otherService.CreateCategory(context.Background(), &categoryReq)
	PanicIfOtherError(err)
	ginCtx.JSON(200, gin.H{"data": categoryRes.CategoryDetail})
}

//GetCategoriesList 使用rpc获取分类列表
func GetCategoriesList(ginCtx *gin.Context) {
	var categoryReq services.CategoryRequest
	PanicIfOtherError(ginCtx.Bind(&categoryReq))
	//从gin.keys取出服务实例
	otherService := ginCtx.Keys["otherService"].(services.OtherService)
	//调用服务端的函数
	categoryRes, err := otherService.GetCategoriesList(context.Background(), &categoryReq)
	PanicIfOtherError(err)
	ginCtx.JSON(200, gin.H{"data": gin.H{"category": categoryRes.CategoriesList, "count": categoryRes.Count}})
}

//UpdateCategory 使用rpc修改分类列表
func UpdateCategory(ginCtx *gin.Context) {
	var categoryReq services.CategoryRequest
	PanicIfOtherError(ginCtx.Bind(&categoryReq))
	//从gin.keys取出服务实例
	otherService := ginCtx.Keys["otherService"].(services.OtherService)
	//调用服务端的函数
	categoryRes, err := otherService.UpdateCategory(context.Background(), &categoryReq)
	PanicIfOtherError(err)
	ginCtx.JSON(200, gin.H{"data": categoryRes.CategoryDetail})
}

//DeleteCategory 使用rpc删除分类列表
func DeleteCategory(ginCtx *gin.Context) {
	var categoryReq services.CategoryRequest
	PanicIfOtherError(ginCtx.Bind(&categoryReq))
	//从gin.keys取出服务实例
	otherService := ginCtx.Keys["otherService"].(services.OtherService)
	//调用服务端的函数
	categoryRes, err := otherService.DeleteCategory(context.Background(), &categoryReq)
	PanicIfOtherError(err)
	ginCtx.JSON(200, gin.H{"data": categoryRes.CategoryDetail})
}
