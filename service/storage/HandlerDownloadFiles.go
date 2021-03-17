package storage

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mholt/archiver"
	"github.com/ratel-drive-core/service/common/util/misc"
)

type DownloadParams struct {
	FilePaths []string `json:"file_paths"`
}

func DownloadSingleFileHandler(c *gin.Context) {
	rootDir := storageConfig.StorageRootDir
	path := c.Query("file")

	targetFilePath := fmt.Sprintf("%s/%s", rootDir, path)
	if misc.IsPathExists(targetFilePath) {
		c.File(targetFilePath)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "failed",
			"error":  "file is not existed",
		})
	}
}

func DownloadMultiFilesHandler(c *gin.Context) {
	var params DownloadParams
	if err := c.ShouldBindJSON(&params); err != nil {
		log.Panicln(err)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println(params.FilePaths)

	if len(params.FilePaths) <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "failed",
			"error":  "no files to download",
		})
		return
	}

	// if download multiple files or directories
	// zip them to a temporary file
	// serve this zipped file
	rootDir := storageConfig.StorageRootDir
	targetFilePath := fmt.Sprintf("/tmp/archive-%d.zip", time.Now().Unix())
	defer os.Remove(targetFilePath)

	var sourceFilesPaths []string
	for _, path := range params.FilePaths {
		sourceFilesPaths = append(sourceFilesPaths, rootDir+"/"+path)
	}

	err := archiver.Archive(sourceFilesPaths, targetFilePath)
	if err != nil {
		log.Panicln(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "failed",
			"error":  "internal issue",
		})
		return
	}

	if misc.IsPathExists(targetFilePath) {
		c.File(targetFilePath)
	} else {
		log.Panicln("[WARN] Something wrong while creating zipped file for downloading")
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "failed",
			"error":  "file is not existed",
		})
	}
}
