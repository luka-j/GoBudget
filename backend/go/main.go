package main

import (
	"GoBudget/Database"
	"GoBudget/Routes"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	r := gin.Default()
	Database.InitDB(os.Getenv("DB_HOST"), os.Getenv("DB_NAME"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"))
	authorizedRoutes := InitAuthMiddleware(r)
	Routes.InitRoutes(r, authorizedRoutes)
	r.Run()
}
