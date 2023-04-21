package main

import (
	"FurballCommunity_backend/config/database"
	"FurballCommunity_backend/models"
	"FurballCommunity_backend/routers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// @title 毛球社区
func main() {
	// 连接数据库
	database.InitMySQL()

	routers.SetupSwagger()
	// 绑定表
	database.DB.AutoMigrate(&models.User{})

	routers.SetupRouter()
}
