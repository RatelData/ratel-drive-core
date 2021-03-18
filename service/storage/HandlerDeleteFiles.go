package storage

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ratel-drive-core/service/common/util/config"
)

func DeleteFileHandler(c *gin.Context) {
	rootDir := config.GetStorageConfig().StorageRootDir
	pathToDel := rootDir + "/" + c.Param("path")

	if err := os.RemoveAll(pathToDel); err != nil {
		log.Panicln("[WARN]", err)

		c.JSON(http.StatusBadRequest, gin.H{
			"result": "failed",
			"error":  err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "success",
	})
}

func EmptyTrashHandler(c *gin.Context) {

}
