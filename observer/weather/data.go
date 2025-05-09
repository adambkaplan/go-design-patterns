package weather

import "sync"

var _ Subject = (*WeatherData)(nil)

type WeatherData struct {
	Temperature float64
	Humidity    float64
	Pressure    float64

	observers     map[Observer]chan<- *WeatherData
	observerLock  *sync.RWMutex
	observersOnce sync.Once

	errChan      chan error
	updateErrors []error
	errorsLock   *sync.RWMutex
	errorsOnce   sync.Once
}

func (d *WeatherData) ensureObserversMap() {
	d.observersOnce.Do(func() {
		d.observers = make(map[Observer]chan<- *WeatherData)
		d.observerLock = &sync.RWMutex{}
	})

}

func (d *WeatherData) ensureErrorHandling() {
	d.errorsOnce.Do(func() {
		d.errChan = make(chan error)
		d.errorsLock = &sync.RWMutex{}
		go func() {
			errChan := d.errChan
			for err := range errChan {
				d.errorsLock.Lock()
				d.updateErrors = append(d.updateErrors, err)
				d.errorsLock.Unlock()
			}
		}()

	})
}

func (d *WeatherData) Errors() []error {
	d.ensureErrorHandling()

	d.errorsLock.RLock()
	defer d.errorsLock.RUnlock()
	// Return a copy of the errors to ensure the returned value does not modify the original array.
	// This could lead to race conditions.
	errorsCopy := make([]error, len(d.updateErrors))
	copy(errorsCopy, d.updateErrors)
	return errorsCopy
}

// NotifyObservers informs the registered observers of updates.
func (d *WeatherData) NotifyObservers() {
	d.ensureObserversMap()

	d.observerLock.RLock()
	defer d.observerLock.RUnlock()
	for _, updateChan := range d.observers {
		updateChan <- d
	}
}

// RegisterObserver registers the provided observer for updates.
func (d *WeatherData) RegisterObserver(o Observer) {
	d.ensureObserversMap()
	d.ensureErrorHandling()

	d.observerLock.Lock()
	defer d.observerLock.Unlock()

	// Exit early if we already registered this observer. We do not want duplicate goroutines for
	// the same observer.
	if _, ok := d.observers[o]; ok {
		return
	}

	updateChan := make(chan *WeatherData)
	d.observers[o] = updateChan
	go func() {
		for data := range updateChan {
			// If an error occurs during update, send it to the error chanel for further processing.
			if err := o.Update(data); err != nil {
				errChan := d.errChan
				errChan <- err
			}
		}
	}()
}

// RemoveObserver stops the provided observer from receiving further updates.
func (d *WeatherData) RemoveObserver(o Observer) {
	d.ensureObserversMap()

	d.observerLock.Lock()
	defer d.observerLock.Unlock()
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
	d.observerLock.Lock()
	defer d.observerLock.Unlock()
	for o := range d.observers {
		d.RemoveObserver(o)
	}
}

func (d *WeatherData) measurementsChanged() {
	d.NotifyObservers()
}
