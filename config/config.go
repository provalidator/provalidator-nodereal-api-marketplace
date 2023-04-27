package config

import (
	"github.com/joho/godotenv"
	"github.com/provalidator-nodereal-api-marketplace/log"
)

func LoadEnv(envName string) {
	err := godotenv.Load(envName)

	if err != nil {
		log.Logger.Error.Fatal("Error loading " + envName)
	}

	log.Logger.Info.Println(envName)
}
