package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	// new 一个 Gin Engine 实例
	r := gin.New()

	// 注册中间件
	r.Use(gin.Logger(), gin.Recovery())

	// 注册一个路由
	r.GET("/", func(c *gin.Context) {

		// 以json 格式响应
		c.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})

	// 处理404请求
	r.NoRoute(func(c *gin.Context) {
		// 获取标头信息的Accept信息
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

	// 运行服务，默认为8080，我们指定端口为8000
	r.Run(":8000")
}
