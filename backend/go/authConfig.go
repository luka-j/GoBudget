package main

import (
	"GoBudget/Database/Entities"
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
)

type User struct {
	Username string
	Email    string
	UserId   string
}

type UserCredentials struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func InitAuthMiddleware(r *gin.Engine) *gin.RouterGroup {
	identityKey := "sub"
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test",
		Key:         []byte(os.Getenv("JWT_KEY")),
		IdentityKey: identityKey,
		Timeout:     time.Hour,
		MaxRefresh:  24 * time.Hour,
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals UserCredentials
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			user, err := Entities.LoginUser(loginVals.Email, loginVals.Password)
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			return &User{user.Username, user.Email, user.ID.String()}, nil
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserId,
					"email":     v.Email,
					"username":  v.Username,
				}
			}
			return jwt.MapClaims{}
		},
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	r.POST("/login", authMiddleware.LoginHandler)
	authorized := r.Group("/")
	authorized.GET("/refresh_token", authMiddleware.RefreshHandler)
	authorized.Use(authMiddleware.MiddlewareFunc())

	return authorized
}
