package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	log "github.com/sirupsen/logrus"

	handler "github.com/KanishkaRR/web-scraper/handler"
	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()

	route.Use(cors.Default())
	handlerRegistration(route)

	port := "8081"
	if err := route.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Unable to start the server, %v", err)
	}

}

func handlerRegistration(router *gin.Engine) {
	handler.RegisterRoutes(router)
}
