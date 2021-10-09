package Routes

import (
	"GoBudget/Database/Entities"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type CreateAccountRequest struct {
	Name string `json:"name" binding:"required"`
}

type RenameAccountRequest struct {
	Id   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type AccountResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

//todo proper error handling where appropriate (no need to 500 when user sends wrong UUID)

func createAccount(c *gin.Context) {
	var request CreateAccountRequest
	if err := c.BindJSON(&request); err != nil {
		return
	}
	tokenClaims := jwt.ExtractClaims(c)
	account, err := Entities.CreateAccount(request.Name, tokenClaims["sub"].(string))
	if err != nil {
		panic(err)
	}
	c.JSON(200, AccountResponse{Id: account.ID.String(), Name: account.Name})
}

func getAccounts(c *gin.Context) {
	userId := jwt.ExtractClaims(c)["sub"].(string)
	accounts, err := Entities.GetAccountsByUser(userId)
	if err != nil {
		panic(err)
	}
	accountsResponse := make([]AccountResponse, len(accounts))
	for i, account := range accounts {
		accountsResponse[i] = AccountResponse{Id: account.ID.String(), Name: account.Name}
	}
	c.JSON(200, accountsResponse)
}

func renameAccount(c *gin.Context) {
	var request RenameAccountRequest
	if err := c.BindJSON(&request); err != nil {
		return
	}
	userId := jwt.ExtractClaims(c)["sub"].(string)
	account, err := Entities.RenameAccount(request.Id, userId, request.Name)
	if err != nil {
		panic(err)
	}
	c.JSON(200, AccountResponse{Id: account.ID.String(), Name: account.Name})
}

func deleteAccount(c *gin.Context) {
	id := c.Param("id")
	userId := jwt.ExtractClaims(c)["sub"].(string)
	if err := Entities.DeleteAccount(id, userId); err != nil {
		panic(err)
	}
	c.Status(200)
}
