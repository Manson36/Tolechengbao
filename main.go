package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lechengbao/handler"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*")

	r.Any("/login", handler.LoginHandler)
	r.Any("/register", handler.RegisterHandler)

	r.Run()
	var v = make(map[string]map[string]string)
}
