package utils

import (
	"encoding/json"
	"io"
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

func DecodeJson(in io.Reader, out interface{}) (err error) {
	err = json.NewDecoder(in).Decode(out)
	return
}

func EndcodeJson(in interface{}) (result []byte, err error) {
	result, err = json.Marshal(in)
	return
}
