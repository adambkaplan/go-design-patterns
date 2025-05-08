package weather

import "github.com/adambkaplan/go-design-patterns/observer"

var _ observer.Subject = (*WeatherData)(nil)

type WeatherData struct {
	Temperature float64
	Humidity    float64
	Pressure    float64

	observers []observer.Observer
}

// NotifyObservers informs the registered observers of updates.
func (d *WeatherData) NotifyObservers() {
	for _, obs := range d.observers {
		obs.Update()
	}
}

// RegisterObserver registers the provided observer for updates.
func (d *WeatherData) RegisterObserver(o observer.Observer) {
	d.observers = append(d.observers, o)
}

// RemoveObserver stops the provided observer from receiving further updates.
func (d *WeatherData) RemoveObserver(o observer.Observer) {
	for i, observer := range d.observers {
		if observer == o {
			d.observers = append(d.observers[:i], d.observers[i+1:]...)
		}
	}
}

// SetMeasurements updates the weather measurement date for this weather station.
func (d *WeatherData) SetMeasurements(temperature float64, humidity float64, pressure float64) {
	d.Temperature = temperature
	d.Humidity = humidity
	d.Pressure = pressure

	d.measurementsChanged()
}

func (d *WeatherData) measurementsChanged() {
	d.NotifyObservers()
}
