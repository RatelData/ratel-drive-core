package storage

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DownloadParams struct {
	FilePaths []string `json:"file_paths"`
}

func DownloadFilesHandler(c *gin.Context) {
	var params DownloadParams
	if err := c.ShouldBindJSON(&params); err != nil {
		log.Panicln(err)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println(params.FilePaths)

	c.JSON(http.StatusOK, gin.H{
		"result": "success",
	})
}
