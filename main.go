package main

import (
	"FurballCommunity_backend/config/database"
	"FurballCommunity_backend/routers"
)

func main() {
	// 连接数据库
	database.InitMySQL()
	// 绑定表
	//database.DB.AutoMigrate(&models.User{})

	router := routers.SetupRouter()
	router.Run(":8081")
}
