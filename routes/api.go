// package routes 注册路由
package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	// 测试一个v1的路由组，我们所有的v1版本的路由都将存放到这里
	v1 := r.Group("v1")
	{
		// 注册一个路由
		v1.GET("/", func(c *gin.Context) {
			// 以json格式响应
			c.JSON(http.StatusOK, gin.H{
				"hello": "world",
			})
		})
	}
}
