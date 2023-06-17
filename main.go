package main

import (
	"FurballCommunity_backend/config/database"
	"FurballCommunity_backend/models"
	"FurballCommunity_backend/routers"
	"log"
)

// @title 毛球社区
func main() {
	// 连接数据库
	database.InitMySQL()

	// 启动接口文档服务
	routers.SetupSwagger()

	// 绑定表
	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.Pet{})
	database.DB.AutoMigrate(&models.Order{})
	err := database.DB.AutoMigrate(&models.OrderCmt{})
	if err != nil {
		log.Println("err:", err)
		return
	}

	//database.DB.Model(&models.Pet{}).Association("Orders")

	// 启动路由服务
	routers.SetupRouter()
}
