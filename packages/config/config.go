package config

import (
	"os"
	"path/filepath"
	"github.com/apex/log"
	"github.com/joho/godotenv"
)

const (
	MYSQL_USER        = "MYSQL_USER"
	MYSQL_PASSWORD    = "MYSQL_PASSWORD"
	MYSQL_DB          = "MYSQL_DB"
	MYSQL_URL		  = "MYSQL_URL"
	CLIENT_URL           = "CLIENT_URL"
	PORT          		 = "PORT"
	JWT_KEY              = "JWT_KEY"
	RUN_MIGRATION        = "RUN_MIGRATION"
	MYSQL_SERVER_HOST = "POSTGRES_SERVER_HOST"
	ENVIRONMEMT           = "ENV"
)

type ConfigType map[string]string

var Config = ConfigType{
	MYSQL_USER:     	  "",
	MYSQL_PASSWORD:    	  "",
	MYSQL_DB:        	  "",
	MYSQL_URL:			  "",
	CLIENT_URL:           "",
	PORT:          		  "",
	JWT_KEY:              "",
	RUN_MIGRATION:        "",
	MYSQL_SERVER_HOST: "localhost",
}

func InitConfig() {
	environment, exists := os.LookupEnv(ENVIRONMEMT)
	var envFilePath string
	if !(exists && environment == "prod") {
		envFilePath, _ = filepath.Abs("../.env")
		if err := godotenv.Load(envFilePath);  environment != "prod" && err != nil {
			log.WithField("reason", err.Error()).Fatal("No .env file found")
		}
	}

	required := map[string]bool{
		MYSQL_USER:     true,
		MYSQL_PASSWORD: true,
		MYSQL_DB:       true,
		MYSQL_URL:		true,
		CLIENT_URL:        true,
		PORT:       	   true,
		RUN_MIGRATION:     true,
	}

	for key := range Config {
		envVal, exists := os.LookupEnv(key)
		if !exists {
			if required[key] {
				log.Fatal(key + " not found in env")
			}
			continue
		}
		if _, ok := Config[key]; ok {
			Config[key] = envVal
		}
	}

	log.Info("All config & secrets set")
}