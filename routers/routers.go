package routers

import (
	"FurballCommunity_backend/controller"
	_ "FurballCommunity_backend/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {

	// 配置路由
	router := gin.Default()
	//设置默认路由当访问一个错误网站时返回
	router.NoRoute(controller.NotFound)
	// 提供静态文件服务 前一个路径为路由路径，后一个路径为文件目录路径
	router.Static("/pubilc/img", "../img")
	router.Use(gin.Logger()) // 设置 gin 的日志级别为 Debug
	// router.Use(middleware.Next()) //添加跨域处理

	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	router.MaxMultipartMemory = 32 << 20 // 32 MiB

	// v1
	v1 := router.Group("/v1")
	{
		api := v1.Group("/api")
		api.POST("/multiUpload", controller.MultiUpload)                     //多文件上传
		api.GET("/getCaptcha", controller.GenerateCaptchaHandler)            //获取图形验证码
		api.POST("/verifyCaptcha", controller.CaptchaVerifyHandle)           //验证图形验证码
		api.POST("/sendMsg", controller.SendMsg)                             //发送手机验证码
		api.POST("/setUserLocation", controller.SetUserLocation)             //保存用户位置信息
		api.POST("/getUserLocationRadius", controller.GetUserLocationRadius) //获取用户50km半径内照护员位置信息

		user := v1.Group("/user")
		user.POST("/login", controller.Login)                   //账号密码登录
		user.POST("/loginWithPhone", controller.LoginWithPhone) //手机号登录（自动注册）
		user.POST("/register", controller.Register)
		user.PUT("/updateUsername/:id", controller.UpdateUserName)
		user.PUT("/updatePassword/:id", controller.UpdatePassword)
		user.DELETE("/deleteUser/:id", controller.DeleteUser)
		user.GET("/getUserList", controller.GetUserList)

		pet := v1.Group("/pet")
		pet.POST("/add", controller.AddPet)
		pet.GET("/getPetInfoByID/:id", controller.GetPetInfoByID)
		pet.GET("/getPetList/:id", controller.GetPetList)
		pet.PUT("/updatePetInfo/:id", controller.UpdatePetInfo)
		pet.DELETE("deletePet/:id", controller.DeletePet)

		order := v1.Group("/order")
		order.POST("/create", controller.CreateOrder)
		order.GET("/getOrderList/:user_id", controller.GetOrderList)
		order.GET("getOrderOfPet/:pet_id", controller.GetOrderOfPet)
		order.PUT("/updateOrderInfo/:order_id", controller.UpdateOrderInfo)

		orderCmt := v1.Group("/orderCmt")
		orderCmt.POST("/create", controller.CreateOrderComment)
		orderCmt.GET("/getOrderCmtList/:order_id", controller.GetCommentListOfOrder)
		orderCmt.DELETE("/deleteOrderCmt/:order_cmt_id", controller.DeleteOrderCmt)
	}

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))

	err := router.Run(":8080")
	if err != nil {
		return nil
	}

	return router
}
