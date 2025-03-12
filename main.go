package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/override/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title           XYZ Multifinance
// @version			0.0.1
// @description     Swagger documentation for XYZ Multifinance API

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

type Object struct {
	Str string
	Num int
}

func main() {
	fmt.Println(fmt.Sprintf("Welcome to XYZ Multifinance!"))

	router := gin.Default()

	// swagger basepath and host
	docs.SwaggerInfo.BasePath = os.Getenv("SWAGGER_BASE_PATH")
	docs.SwaggerInfo.Host = os.Getenv("SWAGGER_HOST")
	fmt.Println("Host: ", docs.SwaggerInfo.Host)

	// connect MySQL
	// connMySQL := conn.ConnectionMapMySQL[os.Getenv("XYZ_MULTIFINANCE_PLATFORM_ENVIRONMENT")]
	// fmt.Println("MySQL DB: " + connMySQL.Database)

	// initialize router for modules
	// loan.InitializeLoan(router, &connMySQL)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := http.ListenAndServe(fmt.Sprintf(":%s", "3000"), router); err != nil {
		log.Fatal(err)
	}
}
