// Package weblib ..
/*
 * @Descripttion:新建gin路由
 * @Author: congz
 * @Date: 2020-09-20 11:33:37
 * @LastEditors: congz
 * @LastEditTime: 2020-10-31 14:18:25
 */
package weblib

import (
	"api-gateway/pkg/util/sdk"
	"api-gateway/weblib/handlers"
	"api-gateway/weblib/middlewares"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

//NewRouter 新建gin路由
func NewRouter(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default()
	//使用中间件，接收服务调用实例
	ginRouter.Use(middlewares.Cors(), middlewares.InitMiddleware(service), middlewares.ErrorMiddleware())
	//使用session中间件来调用极验
	store := cookie.NewStore([]byte(sdk.VERSION))
	ginRouter.Use(sessions.Sessions("mysession", store))

	v1 := ginRouter.Group("/api/v1")
	{
		//用户服务
		v1.POST("/admins/register", handlers.AdminRegister)
		v1.POST("/admins", handlers.AdminLogin)

		v1.GET("/geetest", handlers.InitGeetest)
		// 需要登录保护的
		authed := v1.Group("/")
		authed.Use(middlewares.JWT())
		{
			//用户服务
			//验证token
			authed.GET("ping", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"code": 200,
					"msg":  "success",
				})
			})
			authed.GET("/users", handlers.GetUsersList)
			authed.GET("/users/:user_id", handlers.GetUser)
			//商品服务
			authed.POST("/products", handlers.CreateProduct)
			authed.GET("/products", handlers.GetProductsList)
			authed.GET("/products/:product_id", handlers.GetProductDetail)
			authed.PUT("/products", handlers.UpdateProduct)
			authed.DELETE("/products", handlers.DeleteProduct)
			//商品图片
			authed.POST("/product-imgs", handlers.CreateProductImg)
			authed.GET("/product-imgs", handlers.GetProductImgsList)
			authed.PUT("/product-imgs", handlers.UpdateProductImg)
			authed.DELETE("/product-imgs", handlers.DeleteProductImg)
			//轮播图服务
			authed.GET("/carousels", handlers.GetCarouselsList)
			authed.GET("/carousels/:carousel_id", handlers.GetCarousel)
			authed.PUT("/carousels", handlers.UpdateCarousel)
			//公告服务
			authed.POST("/notices", handlers.CreateNotice)
			authed.GET("/notices/:notice_id", handlers.GetNotice)
			authed.PUT("/notices", handlers.UpdateNotice)
			authed.DELETE("/notices", handlers.DeleteNotice)
			//分类服务
			authed.POST("/categories", handlers.CreateCategory)
			authed.GET("/categories", handlers.GetCategoriesList)
			authed.PUT("/categories", handlers.UpdateCategory)
			authed.DELETE("/categories", handlers.DeleteCategory)
		}
	}
	return ginRouter
}
