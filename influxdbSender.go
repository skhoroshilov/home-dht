package main

import (
	"time"

	"github.com/influxdata/influxdb/client/v2"
)

// MeasurmentData represents a data that should be sent to influxdb
type MeasurmentData interface {
	Name() string
	Data() (tags map[string]string, fields map[string]interface{})
}

// MeasurementSender sends MeasurmentData to metrics storage
type MeasurementSender interface {
	Send(data MeasurmentData) (err error)
	Close() (err error)
}

// InfluxDbSender is MeasurementSender implementation for influxdb
type InfluxDbSender struct {
	client client.Client
}

const (
	databaseName = "data"
)

// NewSender creates new InfluxDbSender object
func NewSender(influxdbAddress string) (sender *InfluxDbSender, err error) {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: influxdbAddress,
	})

	if err == nil {
		sender = &InfluxDbSender{
			client: c,
		}
	}

	return
}

// Send sends MeasurmentData to influxdb
func (sender *InfluxDbSender) Send(data MeasurmentData) (err error) {
	batch, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database: databaseName,
	})

	if err != nil {
		return
	}

	tags, fields := data.Data()
	point, err := client.NewPoint(data.Name(), tags, fields, time.Now())
	if err != nil {
		return
	}

	batch.AddPoint(point)
	err = sender.client.Write(batch)

	return
}

// Close closes influxdb resources
func (sender *InfluxDbSender) Close() error {
	return sender.client.Close()
}
