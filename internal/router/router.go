package router

import (
	"github.com/gin-gonic/gin"
)

func Initializer() {
	router := gin.Default()

	initializerRoutes(router)

	router.Run()
}
