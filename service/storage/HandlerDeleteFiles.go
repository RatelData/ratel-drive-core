package storage

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func DeleteFileHandler(c *gin.Context) {
	rootDir := storageConfig.StorageRootDir
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
