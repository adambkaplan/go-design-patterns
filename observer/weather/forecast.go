package weather

import (
	"io"
	"os"
)

// ForecastDisplay shows a projected forecast based on air pressure changes.
type ForecastDisplay struct {
	Writer io.StringWriter

	LastPressure    float64
	CurrentPressure float64
}

var _ DisplayElement = (*ForecastDisplay)(nil)

// NewForeastDisplay creates a new forecast display and registers it with the weather station.
func NewForecastDisplay(data *WeatherData) *ForecastDisplay {
	display := &ForecastDisplay{
		CurrentPressure: 29.92,
	}
	data.RegisterObserver(display)
	return display
}

// Display displays the weather forecast.
func (f *ForecastDisplay) Display() {
	if f.Writer == nil {
		f.Writer = os.Stdout
	}
	if f.pressureIncreased() {
		f.Writer.WriteString("Forecast: Improving weather on the way!\n")
	} else if f.pressureDecreased() {
		f.Writer.WriteString("Forecast: Watch out for rainy weather\n")
	} else {
		f.Writer.WriteString("Forecast: More of the same\n")
	}
}

func (f *ForecastDisplay) pressureIncreased() bool {
	return (f.CurrentPressure - f.LastPressure) > 0.5
}

func (f *ForecastDisplay) pressureDecreased() bool {
	return (f.CurrentPressure - f.LastPressure) < -0.5
}

// Update retrieves the data from the weather station and updates the display.
func (f *ForecastDisplay) Update(data *WeatherData) {
	f.LastPressure = f.CurrentPressure
	f.CurrentPressure = data.Pressure

	f.Display()
}
