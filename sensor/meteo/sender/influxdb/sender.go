// Package influxdb contains meteo.Sender implementation for sending temperature
// and humidity to influxdb.
package influxdb

import (
	client "github.com/skhoroshilov/home-dht/influxdb"
)

// Sender implents meteo.Sender interface sending data to influxdb.
type Sender struct {
	client client.Sender
}

// NewSender creates new instance of Sender type.
func NewSender(client client.Sender) *Sender {
	return &Sender{
		client: client,
	}
}

// Send sends temperature and humidity to influxdb.
func (sender Sender) Send(temperature float32, humidity float32) error {
	return sender.client.Send(&client.Data{
		Point: "data",
		Fields: map[string]interface{}{
			"temperature": temperature,
			"humidity":    humidity,
		},
	})
}
