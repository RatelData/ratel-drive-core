package storage

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ratel-drive-core/service/common/util/config"
)

func QueryFilesHandler(c *gin.Context) {
	rootDir := config.GetStorageConfig().StorageRootDir
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
