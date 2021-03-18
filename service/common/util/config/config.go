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
	TempDir        string `json:"temp-dir"`
}

var storageConfig *StorageConfig

func GetStorageConfig() *StorageConfig {
	if storageConfig == nil {
		initStorageConfig()
	}
	return storageConfig
}

func initStorageConfig() {
	jsonFile, err := os.Open("config/storage.json")
	var config StorageConfig

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("[SUCCESS] open storage config file")
	defer jsonFile.Close()

	bytes, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(bytes, &config)

	storageConfig = &config
}

type AppConfig struct {
	ServerPort int    `json:"server_port"`
	ServerMode string `json:"server_mode"`
}

var appConfig *AppConfig

func GetServerConfig() *AppConfig {
	if appConfig == nil {
		initServerConfig()
	}

	return appConfig
}

func initServerConfig() {
	jsonFile, err := os.Open("config/app.json")
	var config AppConfig

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("[SUCCESS] open app config file")
	defer jsonFile.Close()

	bytes, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(bytes, &config)

	appConfig = &config
}

func (config *AppConfig) GetServerMode() string {
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
