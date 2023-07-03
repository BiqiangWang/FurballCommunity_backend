package routers

import (
	"FurballCommunity_backend/controller"
	_ "FurballCommunity_backend/docs"
	"FurballCommunity_backend/middleware"

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
	router.Static("/public/img", "../img")
	router.Use(gin.Logger()) // 设置 gin 的日志级别为 Debug
	// router.Use(middleware.Next()) //添加跨域处理

	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	router.MaxMultipartMemory = 32 << 20 // 32 MiB

	// v1
	v1 := router.Group("/v1")
	{
		api := v1.Group("/api")
		api.GET("/getCaptcha", controller.GenerateCaptchaHandler)            //获取图形验证码
		api.POST("/verifyCaptcha", controller.CaptchaVerifyHandle)           //验证图形验证码
		api.POST("/sendMsg", controller.SendMsg)                             //发送手机验证码
		api.POST("/multiUpload", controller.MultiUpload)                     //多文件上传
		api.POST("/setUserLocation", controller.SetUserLocation)             //保存用户位置信息
		api.POST("/getUserLocationRadius", controller.GetUserLocationRadius) //获取用户50km半径内照护员位置信息

		user := v1.Group("/user")
		user.POST("/login", controller.Login)                   //账号密码登录
		user.POST("/loginWithPhone", controller.LoginWithPhone) //手机号登录（自动注册）
		user.POST("/register", controller.Register)
		// 【需要token】中间件验证的路由
		user.Use(middleware.CheckTokenAuth())
		{
			user.PUT("/updateUsername/:id", controller.UpdateUserName)
			user.PUT("/updatePassword/:id", controller.UpdatePassword)
			user.PUT("/updateUserInfo/:id", controller.UpdateUserInfo)
			user.DELETE("/deleteUser/:id", controller.DeleteUser)
			user.GET("/getUserList", controller.GetUserList)
		}

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
		order.GET("/getOrderInfoById/:order_id", controller.GetOrderInfoById)
		order.PUT("/updateOrderInfo/:order_id", controller.UpdateOrderInfo)
		order.DELETE("/delete/:order_id", controller.DeleteOrder)

		orderCmt := v1.Group("/orderCmt")
		orderCmt.POST("/create", controller.CreateOrderComment)
		orderCmt.GET("/getOrderCmtList/:order_id", controller.GetCommentListOfOrder)
		orderCmt.DELETE("/deleteOrderCmt/:order_cmt_id", controller.DeleteOrderCmt)

		blog := v1.Group("/blog")
		blog.POST("/create", controller.CreateBlog)
		blog.GET("getBlogList", controller.GetBlogList)
		blog.GET("/getUserBlog/:id", controller.GetBlogListOfUser)
		blog.GET("/info/:id", controller.GetBlogInfo)
		blog.PUT("/info/:id", controller.UpdateBlog)
		blog.PUT("/like", controller.LikeBlog)
		blog.DELETE("/delete/:blog_id", controller.DeleteBlog)

		blogCmt := v1.Group("/blogCmt")
		blogCmt.POST("/create", controller.CreateBlogComment)
		blogCmt.GET("/getList", controller.GetCommentListOfBlog)
		blogCmt.DELETE("/delete", controller.DeleteBlogCmt)
	}

	// 第二次迭代
	v2 := router.Group("/v2")
	{
		user := v2.Group("/user")
		user.GET("/getUserById/:id", controller.GetUserInfo)
		user.PUT("/updateUserInfo/:id", controller.UpdateUserInfo)
	}

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))

	err := router.Run(":8080")
	if err != nil {
		return nil
	}

	return router
}
