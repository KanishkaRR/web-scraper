package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	route.Use(cors.Default())

	asa := &scraperHandler{}
	print(asa)
}
