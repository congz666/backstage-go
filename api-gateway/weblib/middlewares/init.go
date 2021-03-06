// Package middlewares ..
/*
 * @Descripttion:gin中间件
 * @Author: congz
 * @Date: 2020-09-20 11:33:37
 * @LastEditors: congz
 * @LastEditTime: 2020-10-31 14:59:03
 */
package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//InitMiddleware 接收实例并存到gin.Keys里
func InitMiddleware(service []interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {
		//将实例存到gin.Keys里
		context.Keys = make(map[string]interface{})
		context.Keys["productService"] = service[0]
		context.Keys["otherService"] = service[1]
		context.Keys["userService"] = service[2]
		context.Next()
	}
}

//ErrorMiddleware 错误处理中间件
func ErrorMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				context.JSON(200, gin.H{"code": 404, "msg": fmt.Sprintf("%s", r)})
				context.Abort()
			}
		}()
		context.Next()
	}
}
