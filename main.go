package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

const httpPort = ":8080"

func main() {
	fmt.Println("Server is running on", httpPort)

	server := gin.New()

	server.GET("/health", healthHandler)

	log.Fatal(server.Run(httpPort))
}

func healthHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Everything is good!")
}