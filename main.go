package main

import (
	"FurballCommunity_backend/dao"
	"FurballCommunity_backend/models"
	"FurballCommunity_backend/routers"
)

func main() {
	// 连接数据库
	dao.InitMySQL()
	// 绑定表
	dao.DB.AutoMigrate(&models.User{})

	router := routers.SetupRouter()
	router.Run(":8081")
}
