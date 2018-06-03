// Package meteo provides functions to read temperature and humidity and send it to metrics server.
package meteo

import (
	"errors"
	"fmt"

	"github.com/skhoroshilov/home-dht/log"
)

// Reader provides interface for reading temperature and humidity from sensor.
type Reader interface {
	Read() (temperature float32, humidity float32, err error)
}

// Sender provides interface for sending temperature and humidity to metrics storage.
type Sender interface {
	Send(temperature float32, humidity float32) error
}

// Service reads meteo data from sensor and sends it to metrics storage.
type Service struct {
	log    log.Logger
	reader Reader
	sender Sender
}

// NewService creates new instance of Service type.
func NewService(log log.Logger, reader Reader, sender Sender) *Service {
	return &Service{
		log:    log,
		reader: reader,
		sender: sender,
	}
}

// Send reads meteo data from sensor and sends it to metrics server.
func (service *Service) Send() error {
	temperature, humidity, err := service.reader.Read()
	if err != nil {
		message := fmt.Sprintf("Error reading meteo data from reader: %v\n", err)
		service.log.Error(message)
		return errors.New(message)
	}

	service.log.Debugf("t = %v h = %v\n", temperature, humidity)

	err = service.sender.Send(temperature, humidity)
	if err != nil {
		message := fmt.Sprintf("Error sending meteo data to influxdb: %v\n", err)
		service.log.Error(message)
		return errors.New(message)
	}

	service.log.Debugf("Meteo data sent to influxdb: t = %v h = %v\n", temperature, humidity)
	return nil
}
