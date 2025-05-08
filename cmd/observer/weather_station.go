package observer

import (
	"github.com/adambkaplan/go-design-patterns/observer/weather"
)

func RunWeatherStation() {
	weatherData := &weather.WeatherData{}

	weather.NewCurrentConditionsDisplay(weatherData)
	// Not implemented - other displays
	// weather.NewStatisticsDisplay(weatherData)
	weather.NewForecastDisplay(weatherData)

	weatherData.SetMeasurements(80, 65, 30.8)
	weatherData.SetMeasurements(82, 70, 29.2)
	weatherData.SetMeasurements(78, 90, 29.2)
}
