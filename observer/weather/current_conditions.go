package weather

import (
	"fmt"
	"io"
	"os"
)

// CurrentConditionsDisplay displays the current conditions from the weather station.
type CurrentConditionsDisplay struct {
	Writer      io.StringWriter
	Temperature float64
	Humidity    float64
	Pressure    float64
}

var _ DisplayElement = (*CurrentConditionsDisplay)(nil)

// Display reports the current conditions. By default this writes to `os.Stdout` unless a different
// Writer has been set.
func (c *CurrentConditionsDisplay) Display() (err error) {
	if c.Writer == nil {
		c.Writer = os.Stdout
	}
	_, err = c.Writer.WriteString(fmt.Sprintf("Current conditions: %.fF, %.f%% humidity, %.2f pressure\n", c.Temperature, c.Humidity, c.Pressure))
	return
}

// Update retrieves the data from the weather station and updates the display.
func (c *CurrentConditionsDisplay) Update(data *WeatherData) error {
	c.Temperature = data.Temperature
	c.Humidity = data.Humidity
	c.Pressure = data.Pressure

	return c.Display()
}

// NewCurrentConditionsDisplay creates a new CurrentConditionsDisplay, and registers it for updates
// with the provided weather station.
func NewCurrentConditionsDisplay(data *WeatherData) *CurrentConditionsDisplay {
	display := &CurrentConditionsDisplay{}
	data.RegisterObserver(display)
	return display
}
