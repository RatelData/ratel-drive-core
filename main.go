package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/RatelData/ratel-drive-core/app/models"
	docs "github.com/RatelData/ratel-drive-core/docs"
	"github.com/RatelData/ratel-drive-core/routes"
	"github.com/gin-contrib/cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/RatelData/ratel-drive-core/common/util"
	"github.com/gin-gonic/gin"
)

func AutoMigrate() {
	models.AutoMigrate()
}

// @title RatelDriveCore API
// @version 1.0
// @description RatelDriveCore server
// @termsOfService https://rateldrive.io/terms/

// @contact.name API Support
// @contact.url https://rateldrive.io/support
// @contact.email support@rateldrive.io

// @license.name GNU AFFERO GENERAL PUBLIC LICENSE 3.0
// @license.url https://www.gnu.org/licenses/agpl-3.0.en.html

// @host localhost:8666
// @BasePath /
func main() {
	util.InitLogger()
	defer util.GetLogger().Sync()

	AutoMigrate()

	appConfig := util.GetAppConfig()
	gin.SetMode(appConfig.GetServerMode())

	util.CheckCreateDataDirectory()

	r := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	if !appConfig.IsDebugMode() {
		r.Static("/app", "./ui")
		r.GET("/", func(c *gin.Context) {
			c.Redirect(http.StatusPermanentRedirect, "/app")
		})
	}

	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")
	api_v1 := api.Group("/v1")
	routes.RegisterNonAuthRoutes(api_v1)
	routes.RegisterAuthRoutes(api_v1)

	r.Run(fmt.Sprintf(":%d", appConfig.ServerPort))
}
