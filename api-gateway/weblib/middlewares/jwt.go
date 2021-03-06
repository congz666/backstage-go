//Package middleware ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-07-15 14:48:46
 * @LastEditors: congz
 * @LastEditTime: 2020-10-31 15:45:08
 */
package middlewares

import (
	"api-gateway/pkg/e"
	"api-gateway/pkg/util"
	"os"

	"github.com/gin-gonic/gin"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code uint32

		code = 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if claims.Identification != os.Getenv("Identification") {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			}
		}
		if code != e.SUCCESS {
			c.JSON(200, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
