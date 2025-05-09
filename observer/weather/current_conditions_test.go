package weather_test

import (
	"bytes"
	"time"

	"github.com/adambkaplan/go-design-patterns/observer/weather"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Current Conditions Display", func() {

	When("weather data sets its measurements", func() {

		It("displays the current temperature, humidity, and pressure", func(ctx SpecContext) {
			data := &weather.WeatherData{}
			output := &bytes.Buffer{}
			display := weather.NewCurrentConditionsDisplay(data)
			display.Writer = output
			data.SetMeasurements(80, 90, 30.2)

			outputString := ""
			Eventually(ctx, func() string {
				outputString += output.String()
				return outputString
			}).Should(Equal("Current conditions: 80F, 90% humidity, 30.20 pressure\n"))
		}, SpecTimeout(1*time.Second))
	})
})
