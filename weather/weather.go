package weather

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type LocationRequest struct {
	Name string `json:"name"`
}
type Main struct {
	Temp       float64 `json:"temp"`
	Feels_like float64 `json:"feels_like"`
	Temp_Min   float64 `json:"temp_min"`
	Temp_Max   float64 `json:"temp_max"`
}

type System struct {
	Country string `json:"country"`
	Sunrisr int    `json:"sunrise"`
	Sunset  int    `json:"sunset"`
}
type WeatherData struct {
	Name        string    `json:"name"`
	CurrentTime time.Time `json:"currentTime"`
	Main        Main      `json:"main"`
	Wind        Wind      `json:"wind"`
	TimeZone    int       `json:"timezone"`
	System      System    `json:"sys"`
}
type Wind struct {
	Speed float64 `json:"speed"`
}

func WeatherInfo(c *gin.Context) {
	//	var requestweatherInfo RequestWeatherinfo
	jsonbyte, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Errorf("error in reading", err)
	}
	weatherData, _ := GiveCalltopi(jsonbyte)

	err = json.Unmarshal(jsonbyte, &weatherData)
	if err != nil {
		fmt.Errorf("Unmarshal:", err)
	}
	fmt.Println("location::::", weatherData)

	outputWeatherInfo := ManipulationofData(weatherData)

	c.JSON(http.StatusOK, outputWeatherInfo)
}
func GiveCalltopi(locationName []byte) (WeatherData, error) {
	method := "GET"
	url := "https://api.openweathermap.org/data/2.5/weather?q=nagpur&appid=3a82acc915c2c64f57099d0b80eeb332"
	//url := "https://api.openweathermap.org/data/2.5/weather?q=Nagpur&appid=0b54e1265b60c95d70e64edb83d3c05b"
	body := locationName
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Errorf("error in making request", err)
	}
	Client := http.Client{}
	resp, err := Client.Do(req)
	if err != nil {
		fmt.Errorf("error in Client", err)
	}
	defer resp.Body.Close()
	var weatherData WeatherData
	if http.StatusOK == 200 {
		body, _ := io.ReadAll(resp.Body)
		json.Unmarshal(body, &weatherData)
	}
	fmt.Println("weatherData received from API:::", weatherData)
	return weatherData, nil
}
func ManipulationofData(weatherData WeatherData) WeatherData {
	//°C = 302.16K - 273.15 = 29.01°C
	// Apply conversions
	weatherData.Main.Temp = convertKelvinToCelsius(weatherData.Main.Temp)
	fmt.Println("weatherData.Main.Temp::", weatherData.Main.Temp)
	weatherData.Main.Feels_like = convertKelvinToCelsius(weatherData.Main.Feels_like)
	weatherData.Main.Temp_Min = convertKelvinToCelsius(weatherData.Main.Temp_Min)
	weatherData.Main.Temp_Max = convertKelvinToCelsius(weatherData.Main.Temp_Max)

	return weatherData
}
func convertKelvinToCelsius(kelvin float64) float64 {
	fmt.Println("weatherData.Main.Temp = convertKelvinToCelsius(weatherData.Main.Temp):", kelvin)
	return kelvin - 273.15

}
