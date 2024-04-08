package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// listen any route
	r.Any("/*path", Handler)

	r.Run()
}
