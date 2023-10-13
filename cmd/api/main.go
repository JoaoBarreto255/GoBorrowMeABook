package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	count := 0
	r := gin.Default()
	fmt.Println("Hello world!")
	r.GET("/test", func(ctx *gin.Context) {
		count += 1
		fmt.Printf("Request [%4d]...\n", count)
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	r.Run()
}
