package main

import (
	"FurballCommunity_backend/config/database"
	"FurballCommunity_backend/models"
	"FurballCommunity_backend/routers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// @title 毛球社区
func main() {
	// 连接数据库
	database.InitMySQL()

	// 启动接口文档服务
	routers.SetupSwagger()

	// 绑定表
	database.DB.AutoMigrate(&models.User{})

	// 启动路由服务
	routers.SetupRouter()
}
