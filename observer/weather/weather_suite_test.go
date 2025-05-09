package weather_test

import (
	"bytes"
	"io"
	"sync"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestWeather(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Weather Suite")
}

type ConcurrentWriter struct {
	mutex    *sync.Mutex
	buf      *bytes.Buffer
	initOnce sync.Once
}

func (c *ConcurrentWriter) init() {
	c.initOnce.Do(func() {
		c.mutex = &sync.Mutex{}
		c.buf = &bytes.Buffer{}
	})
}

// WriteString writes the string to a buffer in a thread-safe manner.
func (c *ConcurrentWriter) WriteString(s string) (int, error) {
	c.init()
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.buf.WriteString(s)
}

func (c *ConcurrentWriter) String() string {
	c.init()
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.buf.String()
}

var _ io.StringWriter = (*ConcurrentWriter)(nil)
