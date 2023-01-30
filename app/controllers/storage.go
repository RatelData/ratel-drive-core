package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/RatelData/ratel-drive-core/common/util"
	"github.com/gin-gonic/gin"
	"github.com/mholt/archiver"
)

type FileInfo struct {
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
	IsDir    bool   `json:"is_dir"`
}

// Files godoc
// @tags files
// @summary Retrieve files information
// @description get files by specified path
// @accept  json
// @produce json
// @param   path path string true "the path that you want to list the files"
// @success 200 {object} types.JSONResult{data=[]FileInfo}
// @failure 400 {object} types.ErrorResult{error=string}
// @router /api/storage/files [get]
func QueryFilesHandler(c *gin.Context) {
	rootDir := util.GetStorageConfig().StorageRootDir
	path := c.Query("path")

	path = strings.TrimSuffix(path, "/")

	files, err := ioutil.ReadDir(rootDir + "/" + path)
	if err != nil {
		util.GetLogger().Panic(err.Error())

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var fiArray []FileInfo
	for _, fi := range files {
		fiArray = append(fiArray, FileInfo{fi.Name(), fmt.Sprintf("%s/%s", path, fi.Name()), fi.IsDir()})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": fiArray,
	})
}

func UpdateFilesHandler(c *gin.Context) {
}

func DeleteFileHandler(c *gin.Context) {
	rootDir := util.GetStorageConfig().StorageRootDir
	pathToDel := rootDir + "/" + c.Param("path")

	if err := os.RemoveAll(pathToDel); err != nil {
		util.GetLogger().Panic(err.Error())

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

type DownloadParams struct {
	FilePaths []string `json:"file_paths"`
}

// Files godoc
// @tags files
// @summary Download a single file
// @description Download a single file by the specified file path
// @accept  json
// @produce octet-stream
// @param   path path string true "the file that you want to download"
// @success 200 {file} binary
// @failure 400 {object} types.ErrorResult{error=string}
// @router /api/storage/download [get]
func DownloadSingleFileHandler(c *gin.Context) {
	rootDir := util.GetStorageConfig().StorageRootDir
	path := c.Query("path")

	targetFilePath := fmt.Sprintf("%s/%s", rootDir, path)
	if util.IsPathExists(targetFilePath) {
		c.File(targetFilePath)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "file is not existed",
		})
	}
}

// Files godoc
// @tags files
// @summary Download multiple files
// @description Download files by the specified file paths, will be zipped
// @accept  json
// @produce octet-stream
// @param   files body DownloadParams true "the files that you want to download"
// @success 200 {file} binary
// @failure 400 {object} types.ErrorResult{error=string}
// @failure 500 {object} types.ErrorResult{error=string}
// @router /api/storage/download [post]
func DownloadMultiFilesHandler(c *gin.Context) {
	var params DownloadParams
	if err := c.ShouldBindJSON(&params); err != nil {
		util.GetLogger().Warn(err.Error())

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(params.FilePaths) <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no files to download",
		})
		return
	}

	rootDir := util.GetStorageConfig().StorageRootDir

	// if there is only one file, just serve it.
	if len(params.FilePaths) == 1 {
		path := params.FilePaths[0]
		targetFilePath := fmt.Sprintf("%s/%s", rootDir, path)

		if !util.IsPathDir(targetFilePath) {
			if util.IsPathExists(targetFilePath) {
				c.File(fmt.Sprintf("%s/%s", rootDir, path))
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "file is not existed",
				})
			}
			return
		}
	}

	// if download multiple files or directories
	// zip them to a temporary file
	// serve this zipped file
	tempDir := util.GetStorageConfig().TempDir
	targetFilePath := fmt.Sprintf("%s/archive-%d.zip", tempDir, time.Now().Unix())
	defer os.Remove(targetFilePath)

	var sourceFilesPaths []string
	for _, path := range params.FilePaths {
		sourceFilesPaths = append(sourceFilesPaths, fmt.Sprintf("%s/%s", rootDir, path))
	}

	err := archiver.Archive(sourceFilesPaths, targetFilePath)
	if err != nil {
		util.GetLogger().Panic(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal issue",
		})
		return
	}

	if util.IsPathExists(targetFilePath) {
		c.File(targetFilePath)
	} else {
		util.GetLogger().Warn("Something wrong while creating zipped file for downloading")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "file is not existed",
		})
	}
}

type UploadMetaData struct {
	Dst map[string]string `json:"dst"`
}

func UploadFilesHandler(c *gin.Context) {
	// Multipart form
	form, err_upload := c.MultipartForm()
	if err_upload != nil {
		util.GetLogger().Panic(err_upload.Error())

		c.JSON(http.StatusBadRequest, gin.H{
			"result": "failed",
			"error":  err_upload.Error(),
		})
		return
	}

	files := form.File["upload[]"]
	extraData := form.Value["meta"]

	if len(extraData) <= 0 {
		util.GetLogger().Panic("[WARN] Upload should have meta data")
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
		dst := fmt.Sprintf("%s/%s", util.GetStorageConfig().StorageRootDir, relativeDst)
		if err := os.MkdirAll(filepath.Dir(dst), os.ModePerm); err != nil {
			util.GetLogger().Panic(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"result": "failed",
				"error":  err.Error(),
			})
			return
		}

		if err := c.SaveUploadedFile(file, dst); err != nil {
			hasErr = true
			util.GetLogger().Panic(err.Error())
		}
	}

	if !hasErr {
		c.JSON(http.StatusCreated, gin.H{
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
