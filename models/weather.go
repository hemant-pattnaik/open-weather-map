package models

// Coord represents the coordinates part of the JSON response
type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

// Weather represents the weather part of the JSON response
type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// Main represents the main part of the JSON response
type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
	SeaLevel  int     `json:"sea_level"`
	GrndLevel int     `json:"grnd_level"`
}

// Wind represents the wind part of the JSON response
type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
	Gust  float64 `json:"gust"`
}

// Rain represents the rain part of the JSON response
type Rain struct {
	OneH float64 `json:"1h"`
}

// Clouds represents the clouds part of the JSON response
type Clouds struct {
	All int `json:"all"`
}

// Sys represents the sys part of the JSON response
type Sys struct {
	Type    int    `json:"type"`
	ID      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

// WeatherResponse represents the full JSON response from the weather API
type WeatherResponse struct {
	Coord      Coord     `json:"coord"`
	Weather    []Weather `json:"weather"`
	Base       string    `json:"base"`
	Main       Main      `json:"main"`
	Visibility int       `json:"visibility"`
	Wind       Wind      `json:"wind"`
	Rain       Rain      `json:"rain"`
	Clouds     Clouds    `json:"clouds"`
	Dt         int64     `json:"dt"`
	Sys        Sys       `json:"sys"`
	Timezone   int       `json:"timezone"`
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Cod        int       `json:"cod"`
}

// WeatherResponse represents get curent weather JSON request
type GetCurrentWeatherRequest struct {
	Latitute  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

// WeatherResponse represents get curent weather JSON response
type GetCurrentWeatherResponse struct {
	Condition        string `json:"condition"`
	TemperatureRange string `json:"temperature_range"`
}
