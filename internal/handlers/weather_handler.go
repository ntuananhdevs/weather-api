package handlers

import (
    "net/http"
    "weather-api/internal/services"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    r.GET("/weather/:city", WeatherHandler)
	// r.GET("", handlers ...gin.HandlerFunc)
}

func WeatherHandler(c *gin.Context) {
    city := c.Param("city")
    service := services.WeatherService{}

    weather, err := service.GetWeather(city)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, weather)
}
