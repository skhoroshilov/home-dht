package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Dht22Mock mocks TempHumReader interface
type Dht22Mock struct {
	t   float32
	h   float32
	err error
}

func (mock Dht22Mock) Read() (temperature float32, humidity float32, err error) {
	return mock.t, mock.h, mock.err
}

// InfluxDbSenderMock mocks MeasurementSender interface
type InfluxDbSenderMock struct {
	data MeasurmentData
}

func (mock *InfluxDbSenderMock) Send(data MeasurmentData) (err error) {
	mock.data = data
	return nil
}

func (mock *InfluxDbSenderMock) Close() (err error) {
	return nil
}

// Test checks send() func - one successfull step of DhtSendingJob
// where job reads data using TempHumReader and sends it to influxdb
// using MeasurementSender
func Test_send(t *testing.T) {
	require := require.New(t)

	// arrange

	var tExpected, hExpected float32 = 35.5, 48.2
	dhtMock := &Dht22Mock{t: tExpected, h: hExpected}
	senderMock := &InfluxDbSenderMock{}

	job := DhtSendingJob{reader: dhtMock, sender: senderMock}

	// act

	actualErr := job.send()

	// assert

	require.NoError(actualErr)
	require.NotNil(senderMock.data)

	tags, fields := senderMock.data.Data()
	require.Equal(0, len(tags))
	require.Equal(tExpected, fields["temperature"])
	require.Equal(hExpected, fields["humidity"])
}
