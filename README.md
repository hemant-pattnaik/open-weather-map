# Weather Service API

This is a simple Weather Service API built with Go and Gin that interacts with the OpenWeatherMap API to provide current weather information based on latitude and longitude.

## Features

- Get current weather by providing latitude and longitude as query parameters.
- Easy to set up and run.
- Well-documented API using Swagger.

## Prerequisites

- Go 1.16 or higher
- An OpenWeatherMap API key (ref: https://openweathermap.org/faq)

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/hemant-pattnaik/open-weather-map.git
    cd open-weather-map
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

3. Set up environment variables:

    ```sh
    export API_KEY=your_openweathermap_api_key
    ```

## Usage

1. Run the application:

    ```sh
    go run main.go
    ```

3. Access the weather endpoint:

    ```
    GET /weather?lat={latitude}&lon={longitude}
    ```

    Example:

    ```sh
    curl "http://localhost:8080/weather?lat=35.6895&lon=139.6917"
    ```


