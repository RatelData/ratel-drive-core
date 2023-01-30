package routes

import (
	"github.com/RatelData/ratel-drive-core/app/controllers"
	"github.com/gin-gonic/gin"
)

func StorageRoutesRegister(router *gin.RouterGroup) {
	QueryFilesRegister(router)
	DownloadFilesRegister(router)
	UploadFilesRegister(router)
	DeleteFilesRegister(router)
}

func QueryFilesRegister(router *gin.RouterGroup) {
	router.GET("/files", controllers.QueryFilesHandler)
}

func DownloadFilesRegister(router *gin.RouterGroup) {
	router.GET("/download", controllers.DownloadSingleFileHandler)
	router.POST("/download", controllers.DownloadMultiFilesHandler)
}

func UploadFilesRegister(router *gin.RouterGroup) {
	router.POST("/upload", controllers.UploadFilesHandler)
}

func DeleteFilesRegister(router *gin.RouterGroup) {
	router.DELETE("/files/:path", controllers.DeleteFileHandler)
	router.DELETE("/trash", controllers.EmptyTrashHandler)
}

func UpdateFilesRegister(router *gin.RouterGroup) {
	router.PATCH("/files", controllers.UpdateFilesHandler)
}
