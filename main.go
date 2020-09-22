package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*")

	r.Any("/login", loginHandler)
	r.Any("/register", registerHandler)

	r.Run()
}
