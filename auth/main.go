package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"auth-api/config"
	"auth-api/db"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cfg := config.Configuration{
		ServerPort:  os.Getenv("SERVER_PORT"),
		StoragePath: os.Getenv("STORAGE_PATH"),
		DB: config.DB{
			DbDatabase: os.Getenv("DB_DATABASE"),
			DbHost:     os.Getenv("DB_HOST"),
			DbPort:     os.Getenv("DB_PORT"),
			DbUsername: os.Getenv("DB_USERNAME"),
			DbPassword: os.Getenv("DB_PASSWORD"),
		},
	}

	err := db.InitClients(cfg)
	if err != nil {
		return
	}
	defer db.Client.Close()
	router := gin.Default()

	//loadSwagger(router)
	//setup.LoadUsers(router)

	serverPort := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
	err = router.Run(serverPort)
	if err != nil {
		return
	}
}

//func loadSwagger(router *gin.Engine) {
//	docs.SwaggerInfo.Title = "Swagger Users API"
//	docs.SwaggerInfo.Description = "This is a users microservice."
//	docs.SwaggerInfo.Version = "1.0"
//	docs.SwaggerInfo.Host = "localhost:9010"
//	docs.SwaggerInfo.BasePath = "/"
//	docs.SwaggerInfo.Schemes = []string{"http", "https"}
//
//	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
//
//	router.StaticFile("/docs/swagger3.json", "./docs/swagger3.json")
//}
//
