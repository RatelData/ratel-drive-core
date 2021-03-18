package storage

import (
	"github.com/gin-gonic/gin"
)

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
	router.GET("/download", DownloadSingleFileHandler)
	router.POST("/download", DownloadMultiFilesHandler)
}

func UploadFilesRegister(router *gin.RouterGroup) {
	router.POST("/upload", UploadFilesHandler)
}

func DeleteFilesRegister(router *gin.RouterGroup) {
	router.DELETE("/files/:path", DeleteFileHandler)
	router.DELETE("/trash", EmptyTrashHandler)
}

func UpdateFilesRegister(router *gin.RouterGroup) {
	router.PATCH("/files", UpdateFilesHandler)
}
