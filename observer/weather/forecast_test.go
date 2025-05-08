package weather_test

import (
	"bytes"

	"github.com/adambkaplan/go-design-patterns/observer/weather"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Forecast Display", func() {

	When("the weather data is updated", func() {

		var (
			output  *bytes.Buffer
			data    *weather.WeatherData
			display *weather.ForecastDisplay
		)

		BeforeEach(func() {
			output = &bytes.Buffer{}
			data = &weather.WeatherData{}
			display = weather.NewForecastDisplay(data)
			display.Writer = output
		})

		It("reports the weather is improving if pressure goes up", func() {
			data.SetMeasurements(80, 65, 32.10)
			Expect(output.String()).To(Equal("Forecast: Improving weather on the way!\n"))
		})

		It("reports the weather is getting worse if pressure goes down", func() {
			data.SetMeasurements(75, 95, 28.6)
			Expect(output.String()).To(Equal("Forecast: Watch out for rainy weather\n"))
		})

		It("reports the weather is stable if pressure remains constant", func() {
			data.SetMeasurements(65, 50, 30.0)
			Expect(output.String()).To(Equal("Forecast: More of the same\n"))
		})
	})

})
