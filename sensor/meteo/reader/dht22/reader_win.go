// +build windows

package dht22

// Reader implements meteo.Reader interface for windows returning fake data.
// It is used for developing only (in case of dht library does not compile on Windows).
type Reader struct {
}

// NewReader creates new Reader instance.
func NewReader(pin int) *Reader {
	return &Reader{}
}

const retryCount = 10

func (reader *Reader) Read() (temperature float32, humidity float32, err error) {
	temperature = 25.65
	humidity = 42.21

	return
}
