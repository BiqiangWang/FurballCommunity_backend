package routers

import "FurballCommunity_backend/docs"

func SetupSwagger() {
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Title = "毛球社区"
	docs.SwaggerInfo.Version = "1.0"
}
