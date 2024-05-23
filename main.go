package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hemant-pattnaik/open-weather-map/handler"
	"github.com/hemant-pattnaik/open-weather-map/service"
)

func main() {
	// get the open weather api key to interact with the open weather api
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY environment variable is not set")
	}

	// create the weather service
	weatherService := &service.OpenWeatherMapService{APIKey: apiKey}

	// create the weather http router
	router := gin.Default()
	router.GET("/weather", handler.WeatherHandler(weatherService))

	// set the port at which app will run
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port
	}

	// start the server
	log.Printf("Starting server on port %s...\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err.Error())
	}
}
