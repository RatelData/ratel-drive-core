package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type StorageConfig struct {
	StorageRootDir string `json:"storage-root-dir"`
}

func GetStorageConfig() StorageConfig {
	jsonFile, err := os.Open("config/storage.json")
	var config StorageConfig

	if err != nil {
		log.Println(err)
		return config
	}

	log.Println("[SUCCESS] open storage config file")
	defer jsonFile.Close()

	bytes, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(bytes, &config)

	return config
}

type AppConfig struct {
	ServerPort int    `json:"server_port"`
	ServerMode string `json:"server_mode"`
}

func GetServerConfig() AppConfig {
	jsonFile, err := os.Open("config/app.json")
	var config AppConfig

	if err != nil {
		log.Println(err)
		return config
	}

	log.Println("[SUCCESS] open app config file")
	defer jsonFile.Close()

	bytes, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(bytes, &config)

	return config
}

func (config AppConfig) GetServerMode() string {
	switch mode := config.ServerMode; mode {
	case "debug":
		return gin.DebugMode
	case "release":
		return gin.ReleaseMode
	case "test":
		return gin.TestMode
	default:
		return gin.DebugMode
	}
}
