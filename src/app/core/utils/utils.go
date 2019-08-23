package utils

import (
	"encoding/json"
	"log"
	"os"
)

type configuration struct {
	Server, MongoDBHost, DBUser, DBPwd, DBName string
}

// AppConfig holds the values of config.json file
var AppConfig configuration

// InitConfig set the config from config.json
func InitConfig() {
	loadAppConfig()
}

func loadAppConfig() {
	file, err := os.Open("src/appConfig.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
}
