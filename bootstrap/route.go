// Package bootstrap 处理程序初始化逻辑
package bootstrap

import (
	"gohub/routes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// SetupRoute 路由初始化
func SetupRoute(router *gin.Engine) {

	// 注册全局中间件
	registerGlobalMiddleware(router)

	// 注册 API 路由
	routes.RegisterAPIRoutes(router)

	// 配置 404 路由
	setup404Handler(router)
}

func registerGlobalMiddleware(router *gin.Engine) {
	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)
}

func setup404Handler(router *gin.Engine) {
	// 处理 404 请求
	router.NoRoute(func(c *gin.Context) {
		// 获取标头信息的accept信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是html的话
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			// 默认返回json
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由为定义，请确认url和请求方法是否正确",
			})
		}
	})
}
