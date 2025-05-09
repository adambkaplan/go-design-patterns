package weather_test

import (
	"fmt"
	"io"
	"time"

	"github.com/adambkaplan/go-design-patterns/observer/weather"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("WeatherData", func() {

	When("an observer reports an update error", func() {

		var (
			subject    *weather.WeatherData
			testWriter io.StringWriter
			observer   *weather.CurrentConditionsDisplay
		)

		BeforeEach(func() {
			subject = &weather.WeatherData{}
			testWriter = &errorWriter{}
			observer = weather.NewCurrentConditionsDisplay(subject)
			observer.Writer = testWriter
		})

		It("stores the error for future reference", func(ctx SpecContext) {
			subject.SetMeasurements(80, 62, 31.56)

			Eventually(ctx, len(subject.Errors())).ShouldNot(BeZero())
		}, SpecTimeout(1*time.Second))
	})
})

type errorWriter struct{}

// WriteString always returns an error
func (e *errorWriter) WriteString(_ string) (int, error) {
	return 0, fmt.Errorf("a test error occurred")
}

var _ io.StringWriter = (*errorWriter)(nil)
