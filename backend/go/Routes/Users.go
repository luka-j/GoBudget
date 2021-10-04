package Routes

import (
	"GoBudget/Database/Entities"
	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	Status string `json:"status"`
}

func register(c *gin.Context) {
	var request RegisterRequest
	err := c.BindJSON(&request)
	if err != nil {
		return
	}
	Entities.CreateUser(request.Username, request.Email, request.Password)
	c.JSON(200, UserResponse{
		Status: "SUCCESS",
	})
}
