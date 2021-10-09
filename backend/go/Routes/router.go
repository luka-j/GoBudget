package Routes

import "github.com/gin-gonic/gin"

type PatchRequest struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func InitRoutes(r *gin.Engine, authorized *gin.RouterGroup) {
	r.POST("/register", register)

	authorized.POST("/accounts", createAccount)
	authorized.PUT("/accounts", renameAccount)
	authorized.GET("/accounts", getAccounts)
	authorized.DELETE("/accounts/:id", deleteAccount)
	authorized.POST("/envelopes", createEnvelope)
	authorized.GET("/envelopes", getEnvelopes)
	authorized.PATCH("/envelopes/:id", patchEnvelope)
	authorized.DELETE("/envelopes/:id", deleteEnvelope)
	authorized.GET("/test", hello)
}

func hello(c *gin.Context) {
	c.JSON(200, &gin.H{
		"message": "Hi!",
	})
}
