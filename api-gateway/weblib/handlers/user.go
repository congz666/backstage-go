// Package handlers ..
/*
 * @Descripttion:处理函数
 * @Author: congz
 * @Date: 2020-09-20 11:33:37
 * @LastEditors: congz
 * @LastEditTime: 2020-10-29 14:29:24
 */
package handlers

import (
	"api-gateway/services"
	"context"

	"github.com/gin-gonic/gin"
)

//GetUsersList 使用rpc获取用户列表
func GetUsersList(ginCtx *gin.Context) {
	var userReq services.UserRequest
	PanicIfUserError(ginCtx.Bind(&userReq))
	//从gin.keys取出服务实例
	userService := ginCtx.Keys["userService"].(services.UserService)
	//调用服务端的函数
	userRes, err := userService.GetUsersList(context.Background(), &userReq)
	PanicIfUserError(err)
	ginCtx.JSON(200, gin.H{"data": gin.H{"user": userRes.UsersList, "count": userRes.Count}})
}

//GetUser 使用rpc获取用户详情
func GetUser(ginCtx *gin.Context) {
	var userReq services.UserRequest
	PanicIfUserError(ginCtx.BindUri(&userReq))
	//从gin.keys取出服务实例
	userService := ginCtx.Keys["userService"].(services.UserService)
	//调用服务端的函数
	userRes, err := userService.GetUser(context.Background(), &userReq)
	PanicIfUserError(err)
	ginCtx.JSON(200, gin.H{"data": userRes.UserDetail})
}
