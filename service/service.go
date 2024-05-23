package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hemant-pattnaik/open-weather-map/models"
)

const (
	MIN_COLD = 50
	MIN_HOT  = 70
)

// WeatherService defines the interface for getting weather data
type WeatherService interface {
	GetWeather(req models.GetCurrentWeatherRequest) (models.GetCurrentWeatherResponse, error)
}

// OpenWeatherMapService implements the WeatherService interface
type OpenWeatherMapService struct {
	APIKey string
}

// GetWeather fetches the weather data from OpenWeather API
func (s *OpenWeatherMapService) GetWeather(req models.GetCurrentWeatherRequest) (models.GetCurrentWeatherResponse, error) {
	var response models.GetCurrentWeatherResponse
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s&units=Imperial", req.Latitute, req.Longitude, s.APIKey)
	resp, err := http.Get(url)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	var wr models.WeatherResponse
	err = json.Unmarshal(body, &wr)
	if err != nil {
		return response, err
	}

	condition := wr.Weather[0].Main
	temp := wr.Main.Temp
	var tr string
	if temp < MIN_COLD {
		tr = "cold"
	} else if temp >= MIN_COLD && temp <= MIN_HOT {
		tr = "moderate"
	} else {
		tr = "hot"
	}
	return models.GetCurrentWeatherResponse{Condition: condition, TemperatureRange: tr}, nil
}
