package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func testRouteInitializer(router *gin.RouterGroup) {
	router.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	})
}

func initializerRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		testRouteInitializer(v1)
	}
}
