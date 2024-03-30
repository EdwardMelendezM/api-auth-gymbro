package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/EdwardMelendezM/api-auth/auth/setup"
	"github.com/EdwardMelendezM/api-auth/config"
	"github.com/EdwardMelendezM/api-auth/db"
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
	setup.LoadAuth(router)

	serverPort := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
	err = router.Run(serverPort)
	if err != nil {
		return
	}
}
