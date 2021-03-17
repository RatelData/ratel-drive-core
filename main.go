package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ratel-drive-core/service/common/util/config"
	"github.com/ratel-drive-core/service/storage"
)

func main() {
	appConfig := config.GetServerConfig()
	gin.SetMode(appConfig.GetServerMode())

	r := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	v1 := r.Group("/api")
	v1_storage := v1.Group("/storage")
	storage.RegisterAllRouters(v1_storage)

	r.Run(fmt.Sprintf(":%d", appConfig.ServerPort))
}
