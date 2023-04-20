package routers

import "FurballCommunity_backend/docs"

func SetupSwagger() {
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Title = "毛球社区"
	docs.SwaggerInfo.Version = "1.0"

	docs.SwaggerInfo.Description = "This is a sample server."
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

}
