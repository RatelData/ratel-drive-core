package storage

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type UploadMetaData struct {
	Dst map[string]string `json:"dst"`
}

func UploadMultiFilesHandler(c *gin.Context) {
	// Multipart form
	form, err_upload := c.MultipartForm()
	if err_upload != nil {
		log.Panicln(err_upload)

		c.JSON(http.StatusBadRequest, gin.H{
			"result": "failed",
			"error":  err_upload.Error(),
		})
		return
	}

	files := form.File["upload[]"]
	extraData := form.Value["meta"]

	if len(extraData) <= 0 {
		log.Panicln("[WARN] Upload should have meta data")
	}

	var uploadMeta UploadMetaData
	json.Unmarshal([]byte(extraData[0]), &uploadMeta)

	hasErr := false
	for _, file := range files {
		// Upload the file to specific dst.
		// if cannot find destination path specified in metadata,
		// simply use file name as the dst.
		relativeDst := file.Filename
		if len(extraData) > 0 {
			relativeDst = uploadMeta.Dst[file.Filename]
		}
		dst := fmt.Sprintf("%s/%s", storageConfig.StorageRootDir, relativeDst)
		if err := os.MkdirAll(filepath.Dir(dst), os.ModePerm); err != nil {
			log.Panicln(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"result": "failed",
				"error":  err.Error(),
			})
			return
		}

		if err := c.SaveUploadedFile(file, dst); err != nil {
			hasErr = true
			log.Panicln(err)
		}
	}

	if !hasErr {
		c.JSON(http.StatusOK, gin.H{
			"result": "success",
			"msg":    fmt.Sprintf("%d files uploaded!", len(files)),
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "failed",
			"error":  "failed to save uploaded files!",
		})
	}
}
