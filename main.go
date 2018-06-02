package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/skhoroshilov/home-dht/influxdb"
	"github.com/skhoroshilov/home-dht/task"

	// meteo
	"github.com/skhoroshilov/home-dht/sensor/meteo"
	"github.com/skhoroshilov/home-dht/sensor/meteo/reader/dht22"
	meteosender "github.com/skhoroshilov/home-dht/sensor/meteo/sender/influxdb"
)

const (
	jobInterval = 30 * time.Second
)

func parseSettings() (pin int, influxdbAddress string) {
	flag.IntVar(&pin, "pin", 4, "Dht22 pin address")
	flag.StringVar(&influxdbAddress, "influxdb", "http://192.168.1.92:8086", "Influxdb address")

	flag.Parse()

	return
}

func main() {
	log.Println("Starting")

	pin, influxdbAddress := parseSettings()
	log.Printf("Using influxdb address = %v", influxdbAddress)
	log.Printf("Using dht22 pin = %v", pin)

	ctx := context.TODO()
	tg := task.NewTasksGroup()

	influxdbClient, err := influxdb.NewClient(influxdbAddress)
	if err != nil {
		log.Fatalf("Error creating influxdb client: %v\n", err)
	}

	meteoService := createMeteoService(pin, influxdbClient)
	tg.Start(ctx, jobInterval, func() { meteoService.Send() })

	tg.WaitAll()

	log.Println("Done")
}

func createMeteoService(pin int, influxdbClient influxdb.Sender) *meteo.Service {
	reader := dht22.NewReader(pin)
	sender := meteosender.NewSender(influxdbClient)

	return meteo.NewService(reader, sender)
}
