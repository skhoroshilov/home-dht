// +build !windows

// package dht22 provides meteo.Reader interface reading meteo data from dht sensor.
package dht22

import dht "github.com/d2r2/go-dht"

// Reader implements meteo.Reader interface reading data from dht sensor.
type Reader struct {
	pin int
}

// NewReader creates new Reader instance using pin number used by sensor.
func NewReader(pin int) *Reader {
	return &Reader{
		pin: pin,
	}
}

const retryCount = 10

// Read reads temperature and humidity from dht sensor.
func (reader *Reader) Read() (temperature float32, humidity float32, err error) {
	temperature, humidity, _, err = dht.ReadDHTxxWithRetry(dht.DHT22, rd.pin, true, retryCount)
	return
}
