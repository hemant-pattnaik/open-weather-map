package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hemant-pattnaik/open-weather-map/models"
	"github.com/hemant-pattnaik/open-weather-map/service"
)

func WeatherHandler(service service.WeatherService) gin.HandlerFunc {
	return func(c *gin.Context) {
		latParam := c.Query("lat")
		lonParam := c.Query("lon")

		// basic error checks
		if latParam == "" || lonParam == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "lat and lon parameters are required"})
			return
		}

		lat, err := strconv.ParseFloat(latParam, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid lat parameter"})
			return
		}
		lon, err := strconv.ParseFloat(lonParam, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid lon parameter"})
			return
		}

		// call the weather service
		req := models.GetCurrentWeatherRequest{
			Latitute:  lat,
			Longitude: lon,
		}
		response, err := service.GetWeather(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, response)
	}
}
