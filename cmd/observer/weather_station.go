package observer

import (
	"time"

	"github.com/adambkaplan/go-design-patterns/observer/weather"
)

func RunWeatherStation() {
	weatherData := &weather.WeatherData{}

	weather.NewCurrentConditionsDisplay(weatherData)
	// Not implemented - other displays
	// weather.NewStatisticsDisplay(weatherData)
	weather.NewForecastDisplay(weatherData)

	// Display run concurrently, therefore we need to introduce some delay between each measurement
	// update.
	weatherData.SetMeasurements(80, 65, 30.8)
	time.Sleep(10 * time.Millisecond)
	weatherData.SetMeasurements(82, 70, 29.2)
	time.Sleep(10 * time.Millisecond)
	weatherData.SetMeasurements(78, 90, 29.2)
	time.Sleep(10 * time.Millisecond)
}
