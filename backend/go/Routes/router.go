package Routes

import "github.com/gin-gonic/gin"

func InitRoutes(r *gin.Engine, authorized *gin.RouterGroup) {
	r.POST("/register", register)

	authorized.POST("/accounts", createAccount)
	authorized.PUT("/accounts", renameAccount)
	authorized.GET("/accounts", getAccounts)
	authorized.GET("/test", hello)
}

func hello(c *gin.Context) {
	c.JSON(200, &gin.H{
		"message": "Hi!",
	})
}
