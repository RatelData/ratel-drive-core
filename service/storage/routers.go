package storage

import (
	"github.com/gin-gonic/gin"
	"github.com/ratel-drive-core/service/common/util/config"
)

var storageConfig = config.GetStorageConfig()

func RegisterAllRouters(router *gin.RouterGroup) {
	QueryFilesRegister(router)
	DownloadFilesRegister(router)
	UploadFilesRegister(router)
	DeleteFilesRegister(router)
}

func QueryFilesRegister(router *gin.RouterGroup) {
	router.GET("/files", QueryFilesHandler)
}

func DownloadFilesRegister(router *gin.RouterGroup) {
	router.POST("/download", DownloadFilesHandler)
}

func UploadFilesRegister(router *gin.RouterGroup) {
	router.POST("/upload", UploadMultiFilesHandler)
}

func DeleteFilesRegister(router *gin.RouterGroup) {
	router.DELETE("/files", DeleteFilesHandler)
}
