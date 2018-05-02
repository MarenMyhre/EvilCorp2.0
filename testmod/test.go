package main
(
import "github.com/ramsgoli/openweathermap"
)

owm := openweathermap.OpenWeatherMap{API_KEY: os.Getenv("OWM_APP_ID")}