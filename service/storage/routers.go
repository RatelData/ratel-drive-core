package storage

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ratel-drive-core/service/common/util/config"
)

var storageConfig = config.GetStorageConfig()

func GetFilesRegister(router *gin.RouterGroup) {
	router.GET("/files", GetFilesHandler)
}

func DownloadFilesRegister(router *gin.RouterGroup) {
	router.GET("/files", DownloadFilesHandler)
}

func UploadFilesRegister(router *gin.RouterGroup) {
	router.POST("/upload", UploadFilesHandler)
}

func DeleteFilesRegister(router *gin.RouterGroup) {
	router.DELETE("/files", DeleteFilesHandler)
}

func GetFilesHandler(c *gin.Context) {
	rootDir := storageConfig.StorageRootDir
	path := c.Query("path")

	files, err := ioutil.ReadDir(rootDir + "/" + path)
	if err != nil {
		log.Panic(err)

		c.JSON(http.StatusBadRequest, gin.H{
			"result": "failed",
			"error":  err.Error(),
		})
		return
	}

	var fiArray []gin.H
	for _, fi := range files {
		fiArray = append(fiArray, gin.H{
			"file_name": fi.Name(),
			"is_dir":    fi.IsDir(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "success",
		"data":   fiArray,
	})
}

func DownloadFilesHandler(c *gin.Context) {

}

func UploadFilesHandler(c *gin.Context) {
	if err := os.MkdirAll(storageConfig.StorageRootDir, os.ModePerm); err != nil {
		log.Panic(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "failed",
			"error":  err.Error(),
		})
		return
	}

	// Multipart form
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]

	hasErr := false
	for _, file := range files {
		log.Println(file.Filename)

		// Upload the file to specific dst.
		dst := fmt.Sprintf("%s/%s", storageConfig.StorageRootDir, file.Filename)
		err := c.SaveUploadedFile(file, dst)
		if err != nil {
			hasErr = true

			log.Println(err)

			c.JSON(http.StatusBadRequest, gin.H{
				"result": "failed",
				"error":  err.Error(),
			})
		}
	}

	if !hasErr {
		c.JSON(http.StatusOK, gin.H{
			"result": "success",
			"msg":    fmt.Sprintf("%d files uploaded!", len(files)),
		})
	}
}

func DeleteFilesHandler(c *gin.Context) {

}
