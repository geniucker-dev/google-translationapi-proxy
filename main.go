package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// disable debug mode
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// listen any route
	r.Any("/*path", Handler)

	r.Run()
}
