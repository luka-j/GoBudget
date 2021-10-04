package Routes

import "github.com/gin-gonic/gin"

func InitRoutes(r *gin.Engine, authorized *gin.RouterGroup) {
	r.POST("/register", register)
	authorized.GET("/test", hello)
}

func hello(c *gin.Context) {
	c.JSON(200, &gin.H{
		"message": "Hi!",
	})
}
