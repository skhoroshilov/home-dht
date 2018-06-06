package meteo

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/skhoroshilov/home-dht/log"
)

// Dht22Mock mocks Reader interface
type Dht22Mock struct {
	t   float32
	h   float32
	err error
}

func (mock Dht22Mock) Read() (temperature float32, humidity float32, err error) {
	return mock.t, mock.h, mock.err
}

// InfluxDbSenderMock mocks Sender interface
type InfluxDbSenderMock struct {
	sent bool
	t    float32
	h    float32
}

func (mock *InfluxDbSenderMock) Send(temperature float32, humidity float32) error {
	mock.t = temperature
	mock.h = humidity
	mock.sent = true
	return nil
}

// Test checks Send() func where service reads data using Reader and
// sends it to influxdb using Sender
func Test_Send_Success(t *testing.T) {
	require := require.New(t)

	// arrange

	var tExpected, hExpected float32 = 35.5, 48.2
	log := log.NewMock()
	reader := &Dht22Mock{t: tExpected, h: hExpected}
	sender := &InfluxDbSenderMock{}
	service := NewService(log, reader, sender)

	// act

	err := service.Send()

	// assert

	require.NoError(err)
	require.True(sender.sent)
	require.Equal(tExpected, sender.t)
	require.Equal(hExpected, sender.h)
}

// Test checks Send() func where service cannot read data using Reader
func Test_Send_ReadFailed(t *testing.T) {
	require := require.New(t)

	// arrange

	expectedError := errors.New("Some erorr reading data from dht reader")
	log := log.NewMock()
	reader := &Dht22Mock{err: expectedError}
	sender := &InfluxDbSenderMock{}
	service := NewService(log, reader, sender)

	// act

	err := service.Send()

	// assert

	require.Error(err)
	require.False(sender.sent)
}
