package web

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MyMockedObjectConfig struct {
	mock.Mock
}

func (m *MyMockedObjectConfig) GetPort() (string, error) {
	return "1111", nil
}

func (m *MyMockedObjectConfig) GetUrl() (string, error) {
	return "localhost", nil
}

func (m *MyMockedObjectConfig) GetEnvironment() (string, error) {
	return "local", nil
}

func (m *MyMockedObjectConfig) GetRegion() (string, error) {
	return "east", nil
}

func (m *MyMockedObjectConfig) GetVersionApp() (string, error) {
	return "1.0", nil
}

func Test_Process(t *testing.T) {
	t.Parallel()

	// create an instance of our test object
	testObj := new(MyMockedObjectConfig)

	// setup expectations
	testObj.On("GetPort").Return("9191", nil)

	var err error
	c := make(chan string)
	go func() {
		c <- "StartServer"
		err = NewWebServer(testObj).Start()
	}()
	time.Sleep(2 * time.Second)

	x := <-c
	t.Run("Validate Server", func(t *testing.T) {
		assert.Equal(t, nil, err)
		assert.Equal(t, "StartServer", x)
	})
}
