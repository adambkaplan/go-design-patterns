package weather

// DisplayElements are capable of displaying weather data.
type DisplayElement interface {
	Display() error
}

type Subject interface {
	RegisterObserver(o Observer)

	RemoveObserver(o Observer)

	NotifyObservers()
}

type Observer interface {
	Update(data *WeatherData) error
}
