package main

import (
	"context"
	"log"
	"time"
)

const (
	// Time interval before next DhtSendingJob iteration
	// if there were no error in previous iteration
	waitErrorTime = 30 * time.Second

	// Time interval before next DhtSendingJob iteration
	// if there were errors in previous iteration
	waitSuccessTime = 30 * time.Second
)

// DhtSendingJob periodically receives temperature and hummidity
// and sends it to influxdb
type DhtSendingJob struct {
	reader TempHumReader
	sender MeasurementSender
}

// NewDhtSendingJob returns new DhtSendingJob object
func NewDhtSendingJob(reader TempHumReader, sender MeasurementSender) *DhtSendingJob {
	return &DhtSendingJob{
		reader: reader,
		sender: sender,
	}
}

// Start starts DhtSendingJob
func (job *DhtSendingJob) Start(ctx context.Context) {
	for {
		err := job.send()

		waitTime := waitSuccessTime
		if err != nil {
			waitTime = waitErrorTime
		}

		select {
		case <-time.After(waitTime):
		case <-ctx.Done():
			break
		}
	}
}

func (job *DhtSendingJob) send() error {
	temperature, humidity, err := job.reader.Read()
	if err != nil {
		log.Printf("Error reading data from dht reader: %v\n", err)
		return err
	}

	log.Printf("temperature = %v humidity = %v\n", temperature, humidity)

	data := TempHumData{
		temperature: temperature,
		humidity:    humidity,
	}

	err = job.sender.Send(data)
	if err != nil {
		log.Printf("Error sending data to influxdb: %v\n", err)
		return err
	}

	log.Printf("Measurements sent to influxdb: t = %v h = %v\n", temperature, humidity)
	return nil
}
