package routes

import "github.com/gin-gonic/gin"

func RegisterNonAuthRoutes(router *gin.RouterGroup) {
	UsersNonAuthRoutesRegister(router)
}

func RegisterAuthRoutes(router *gin.RouterGroup) {
	StorageRoutesRegister(router.Group("/storage"))
}
