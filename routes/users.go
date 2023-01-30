package routes

import (
	"github.com/RatelData/ratel-drive-core/app/controllers"
	"github.com/gin-gonic/gin"
)

func UsersNonAuthRoutesRegister(router *gin.RouterGroup) {
	router.POST("/login", controllers.UserLogin)
}
