// Package handlers ..
/*
 * @Descripttion:处理函数
 * @Author: congz
 * @Date: 2020-09-20 11:33:37
 * @LastEditors: congz
 * @LastEditTime: 2020-10-27 14:09:06
 */
package handlers

import (
	"api-gateway/services"
	"context"

	"github.com/gin-gonic/gin"
)

//CreateNotice 使用rpc创建公告
func CreateNotice(ginCtx *gin.Context) {
	var noticeReq services.NoticeRequest
	PanicIfOtherError(ginCtx.Bind(&noticeReq))
	//从gin.keys取出服务实例
	otherService := ginCtx.Keys["otherService"].(services.OtherService)
	//调用服务端的函数
	noticeRes, err := otherService.CreateNotice(context.Background(), &noticeReq)
	PanicIfOtherError(err)
	ginCtx.JSON(200, gin.H{"data": noticeRes.NoticeDetail})
}

//GetNotice 使用rpc获取公告
func GetNotice(ginCtx *gin.Context) {
	var noticeReq services.NoticeRequest
	PanicIfOtherError(ginCtx.BindUri(&noticeReq))
	//从gin.keys取出服务实例
	otherService := ginCtx.Keys["otherService"].(services.OtherService)
	//调用服务端的函数
	noticeRes, err := otherService.GetNotice(context.Background(), &noticeReq)
	PanicIfOtherError(err)
	ginCtx.JSON(200, gin.H{"data": noticeRes.NoticeDetail})
}

//UpdateNotice 使用rpc修改公告
func UpdateNotice(ginCtx *gin.Context) {
	var noticeReq services.NoticeRequest
	PanicIfOtherError(ginCtx.Bind(&noticeReq))
	//从gin.keys取出服务实例
	otherService := ginCtx.Keys["otherService"].(services.OtherService)
	//调用服务端的函数
	noticeRes, err := otherService.UpdateNotice(context.Background(), &noticeReq)
	PanicIfOtherError(err)
	ginCtx.JSON(200, gin.H{"data": noticeRes.NoticeDetail})
}

//DeleteNotice 使用rpc删除公告
func DeleteNotice(ginCtx *gin.Context) {
	var noticeReq services.NoticeRequest
	PanicIfOtherError(ginCtx.Bind(&noticeReq))
	//从gin.keys取出服务实例
	otherService := ginCtx.Keys["otherService"].(services.OtherService)
	//调用服务端的函数
	noticeRes, err := otherService.DeleteNotice(context.Background(), &noticeReq)
	PanicIfOtherError(err)
	ginCtx.JSON(200, gin.H{"data": noticeRes.NoticeDetail})
}
