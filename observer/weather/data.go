package weather

var _ Subject = (*WeatherData)(nil)

type WeatherData struct {
	Temperature float64
	Humidity    float64
	Pressure    float64

	observers map[Observer]chan<- *WeatherData
}

// NotifyObservers informs the registered observers of updates.
func (d *WeatherData) NotifyObservers() {
	for _, updateChan := range d.observers {
		updateChan <- d
	}
}

// RegisterObserver registers the provided observer for updates.
func (d *WeatherData) RegisterObserver(o Observer) {
	if d.observers == nil {
		d.observers = make(map[Observer]chan<- *WeatherData)
	}
	updateChan := make(chan *WeatherData)
	d.observers[o] = updateChan
	go func() {
		for data := range updateChan {
			o.Update(data)
		}
	}()
}

// RemoveObserver stops the provided observer from receiving further updates.
func (d *WeatherData) RemoveObserver(o Observer) {
	updateChan := d.observers[o]
	close(updateChan)
	delete(d.observers, o)
}

// SetMeasurements updates the weather measurement date for this weather station.
func (d *WeatherData) SetMeasurements(temperature float64, humidity float64, pressure float64) {
	d.Temperature = temperature
	d.Humidity = humidity
	d.Pressure = pressure

	d.measurementsChanged()
}

func (d *WeatherData) Stop() {
	for o := range d.observers {
		d.RemoveObserver(o)
	}
}

func (d *WeatherData) measurementsChanged() {
	d.NotifyObservers()
}
