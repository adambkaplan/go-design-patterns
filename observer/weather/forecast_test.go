package weather_test

import (
	"time"

	"github.com/adambkaplan/go-design-patterns/observer/weather"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Forecast Display", func() {

	When("the weather data is updated", func() {

		var (
			output  *ConcurrentWriter
			data    *weather.WeatherData
			display *weather.ForecastDisplay
		)

		BeforeEach(func() {
			output = &ConcurrentWriter{}
			data = &weather.WeatherData{}
			display = weather.NewForecastDisplay(data)
			display.Writer = output
		})

		It("reports the weather is improving if pressure goes up", func(ctx SpecContext) {
			data.SetMeasurements(80, 65, 32.10)
			outputString := ""
			Eventually(ctx, func() string {
				outputString += output.String()
				return outputString
			}).Should(Equal("Forecast: Improving weather on the way!\n"))
		}, SpecTimeout(1*time.Second))

		It("reports the weather is getting worse if pressure goes down", func(ctx SpecContext) {
			data.SetMeasurements(75, 95, 28.6)
			outputString := ""
			Eventually(ctx, func() string {
				outputString += output.String()
				return outputString
			}).Should(Equal("Forecast: Watch out for rainy weather\n"))
		}, SpecTimeout(1*time.Second))

		It("reports the weather is stable if pressure remains constant", func(ctx SpecContext) {
			data.SetMeasurements(65, 50, 30.0)
			outputString := ""
			Eventually(ctx, func() string {
				outputString += output.String()
				return outputString
			}).Should(Equal("Forecast: More of the same\n"))
		}, SpecTimeout(1*time.Second))
	})

})
