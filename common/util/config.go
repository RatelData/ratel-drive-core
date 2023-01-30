package util

import (
	"encoding/json"
	"fmt"
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
	ServerPort      int    `json:"server_port"`
	ServerMode      string `json:"server_mode"`
	CentralServer   string `json:"central_server"`
	CentralPort     int    `json:"central_port"`
	CentralProtocol string `json:"central_protocol"`
	DatabasePath    string `json:"database_path"`
}

var appConfig *AppConfig
var appConfigFilePath string

func SetAppConfigFilePath(path string) {
	appConfigFilePath = path
}

func GetAppConfig() *AppConfig {
	if appConfig == nil {
		initAppConfig()
	}

	return appConfig
}

func initAppConfig() {
	if appConfigFilePath == "" {
		appConfigFilePath = "config/app.json"
	}
	jsonFile, err := os.Open(appConfigFilePath)
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
	case "development":
		return gin.DebugMode
	case "production":
		return gin.ReleaseMode
	case "test":
		return gin.TestMode
	default:
		return gin.DebugMode
	}
}

func (config *AppConfig) IsDebugMode() bool {
	return config.GetServerMode() == gin.DebugMode
}

func CentralHost() string {
	config := GetAppConfig()
	return fmt.Sprintf("%s://%s:%d", config.CentralProtocol, config.CentralServer, config.CentralPort)
}
