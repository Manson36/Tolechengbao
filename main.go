package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lechengbao/handler"
	"net/http"
)

func main() {
	r := gin.Default()
	//r.LoadHTMLGlob("template/*")

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.Group("/user").
		POST("/login", handler.Login).
		Any("/register", handler.Register)

	r.Run()
}
