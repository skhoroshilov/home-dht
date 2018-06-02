// Package influxdb provides functionality for sending data to influxdb.
// It uses github.com/influxdata/influxdb/client/v2 as influxdb client.
package influxdb

import (
	"time"

	"github.com/influxdata/influxdb/client/v2"
)

const (
	// database is influxdb database name used for all data.
	database = "data"
)

// Data represents a data sending to influxdb.
type Data struct {
	Point  string                 // point name in influxdb
	Tags   map[string]string      // tags sending to influxdb
	Fields map[string]interface{} // fields sending to influxdb
}

// Sender is the interface that wraps the basic Send method.
//
// Send sends data to influxdb.
type Sender interface {
	Send(data *Data) error
}

// Client implements Sender interface for sending data to influxdb.
type Client struct {
	client client.Client
}

// NewClient creates new client.
// Parameters:
//   address - influxdb address.
func NewClient(address string) (sender *Client, err error) {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: address,
	})

	if err == nil {
		sender = &Client{
			client: c,
		}
	}

	return
}

// Send sends data to influxdb (Sender interface implementation).
func (sender *Client) Send(data *Data) error {
	batch, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database: database,
	})

	if err != nil {
		return err
	}

	point, err := client.NewPoint(data.Point, data.Tags, data.Fields, time.Now())
	if err != nil {
		return err
	}

	batch.AddPoint(point)
	err = sender.client.Write(batch)

	return err
}

// Close closes resources allocated for Client.
func (sender *Client) Close() error {
	return sender.Close()
}
