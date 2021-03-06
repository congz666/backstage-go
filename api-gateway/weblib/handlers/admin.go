// Package handlers ..
/*
 * @Descripttion:处理函数
 * @Author: congz
 * @Date: 2020-09-20 11:33:37
 * @LastEditors: congz
 * @LastEditTime: 2020-10-31 14:21:06
 */
package handlers

import (
	"api-gateway/pkg/e"
	"api-gateway/pkg/util"
	"api-gateway/pkg/util/sdk"
	"api-gateway/services"
	"context"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//AdminRegister 使用rpc进行管理员注册
func AdminRegister(ginCtx *gin.Context) {
	var adminReq services.AdminRequest
	PanicIfUserError(ginCtx.Bind(&adminReq))
	//从gin.keys取出服务实例
	userService := ginCtx.Keys["userService"].(services.UserService)
	//调用服务端的函数
	adminRes, err := userService.AdminRegister(context.Background(), &adminReq)
	PanicIfUserError(err)
	ginCtx.JSON(200, gin.H{"data": adminRes.AdminDetail})
}

//AdminLogin 使用rpc进行管理员登录
func AdminLogin(ginCtx *gin.Context) {
	var adminReq services.AdminRequest
	PanicIfUserError(ginCtx.Bind(&adminReq))
	//极验SDK验证
	session := sessions.Default(ginCtx)
	status := session.Get(sdk.GEETEST_SERVER_STATUS_SESSION_KEY)
	userID := session.Get("userId")
	gtLib := sdk.NewGeetestLib(os.Getenv("GEETEST_ID"), os.Getenv("GEETEST_KEY"))
	var result *sdk.GeetestLibResult
	if status.(int) == 1 {
		/*
			 自定义参数,可选择添加
					 user_id 客户端用户的唯一标识，确定用户的唯一性；作用于提供进阶数据分析服务，可在register和validate接口传入，不传入也不影响验证服务的使用；若担心用户信息风险，可作预处理(如哈希处理)再提供到极验
				 client_type 客户端类型，web：电脑上的浏览器；h5：手机上的浏览器，包括移动应用内完全内置的web_view；native：通过原生sdk植入app应用的方式；unknown：未知
				 ip_address 客户端请求sdk服务器的ip地址
		*/
		params := map[string]string{
			"user_id":     userID.(string),
			"client_type": "web",
			"ip_address":  "127.0.0.1",
		}
		result = gtLib.SuccessValidate(adminReq.Challenge, adminReq.Validate, adminReq.Seccode, params)
	} else {
		result = gtLib.FailValidate(adminReq.Challenge, adminReq.Validate, adminReq.Seccode)
	}
	// 注意，不要更改返回的结构和值类型
	if result.Status != 1 {
		ginCtx.JSON(200, gin.H{"code": 404, "msg": result.Msg})
		return
	}
	//从gin.keys取出服务实例
	userService := ginCtx.Keys["userService"].(services.UserService)
	//调用服务端的函数
	adminRes, err := userService.AdminLogin(context.Background(), &adminReq)
	PanicIfUserError(err)
	token, err := util.GenerateToken()
	if err != nil {
		adminRes.Code = e.ERROR_AUTH_TOKEN
	}
	ginCtx.JSON(200, gin.H{"code": adminRes.Code, "msg": e.GetMsg(adminRes.Code), "data": gin.H{"admin": adminRes.AdminDetail, "token": token}})
}

// InitGeetest 极验初始化
func InitGeetest(c *gin.Context) {
	gtLib := sdk.NewGeetestLib(os.Getenv("GEETEST_ID"), os.Getenv("GEETEST_KEY"))
	digestmod := "md5"
	userID := "test"
	params := map[string]string{
		"digestmod":   digestmod,
		"user_id":     userID,
		"client_type": "web",
		"ip_address":  "127.0.0.1",
	}
	result := gtLib.Register(digestmod, params)
	// 将结果状态写到session中，此处register接口存入session，后续validate接口会取出使用
	// 注意，此demo应用的session是单机模式，格外注意分布式环境下session的应用
	session := sessions.Default(c)
	session.Set(sdk.GEETEST_SERVER_STATUS_SESSION_KEY, result.Status)
	session.Set("userId", userID)
	session.Save()
	// 注意，不要更改返回的结构和值类型
	c.Header("Content-Type", "application/json;charset=UTF-8")
	c.String(http.StatusOK, result.Data)
}
