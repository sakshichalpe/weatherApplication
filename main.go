package main

import (
	"weathertracker/weather"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/getLocation", weather.WeatherInfo)
	r.Run("localhost:8080")
}
