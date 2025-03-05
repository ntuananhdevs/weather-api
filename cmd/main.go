package main

import (
	"weather-api/config"
	"weather-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	r := gin.Default()
	handlers.SetupRoutes(r)

	r.Run(":8080")
}
