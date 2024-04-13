package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/EdwardMelendezM/api-auth/authentication/setup"
	"github.com/EdwardMelendezM/api-info-shared/config"
	"github.com/EdwardMelendezM/api-info-shared/db"
)

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

	setup.LoadAuthentication(router)

	serverPort := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
	err = router.Run(serverPort)
	if err != nil {
		return
	}
}
