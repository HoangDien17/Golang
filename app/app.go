package app

import (
	db "employee/connect"
	"employee/docs"
	route "employee/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Routes struct {
}

func (app Routes) Bootstrap() {

	docs.SwaggerInfo.Title = "Swagger Example Employee API"
	docs.SwaggerInfo.Description = "This is a Employee server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:4000"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.Default()

	r.Use(gin.Logger())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	publicRoute := r.Group("/api/v1")
	resource, err := db.InitResource()
	if err != nil {
		logrus.Error(err)
	}
	defer resource.Close()

	route.UserRoute(publicRoute, resource)
	route.RoleRoute(publicRoute, resource)
	r.Run(":" + os.Getenv("PORT"))
}
