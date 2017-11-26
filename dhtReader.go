package main

import "github.com/d2r2/go-dht"

// TempHumReader reads temperature and humidity from temperature sensor
type TempHumReader interface {
	Read() (temperature float32, humidity float32, err error)
}

// Dht22Reader implements TempHumReader using DHT22 sensor
type Dht22Reader struct {
	pin int
}

// NewReader returns new Dht22Reader object
//   pin - DHT22 pin address
func NewReader(pin int) (reader *Dht22Reader) {
	return &Dht22Reader{
		pin: pin,
	}
}

func (rd Dht22Reader) Read() (temperature float32, humidity float32, err error) {
	temperature, humidity, _, err = dht.ReadDHTxxWithRetry(dht.DHT22, rd.pin, true, 10)
	return
}
