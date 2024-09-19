package main

import (
	"happy-gin/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/demo_ok", handler.DemoOk)
	r.GET("/demo_err", handler.DemoErr)
	r.GET("/demo_err_code", handler.DemoErrWithCode)

	r.Run()
}
