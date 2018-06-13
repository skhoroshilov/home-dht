package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/skhoroshilov/home-dht/influxdb"
	"github.com/skhoroshilov/home-dht/task"

	"github.com/skhoroshilov/home-dht/log"
	stdlog "github.com/skhoroshilov/home-dht/log/std"

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
	log := stdlog.NewLogger()

	log.Info("Starting")

	pin, influxdbAddress := parseSettings()
	log.Infof("Using influxdb address = %v", influxdbAddress)
	log.Infof("Using dht22 pin = %v", pin)

	ctx, cancel := context.WithCancel(context.TODO())
	tg := task.NewTasksGroup()

	influxdbClient, err := influxdb.NewClient(influxdbAddress)
	if err != nil {
		log.Fatalf("Error creating influxdb client: %v\n", err)
	}
	defer influxdbClient.Close()

	meteoService := createMeteoService(log, pin, influxdbClient)
	tg.Start(ctx, jobInterval, func() {
		_ = meteoService.Send()
	})

	setupCancellation(log, cancel)
	tg.WaitAll()

	log.Info("Done")
}

func createMeteoService(log log.Logger, pin int, influxdbClient influxdb.Sender) *meteo.Service {
	reader := dht22.NewReader(pin)
	sender := meteosender.NewSender(influxdbClient)

	return meteo.NewService(log, reader, sender)
}

func setupCancellation(log log.Logger, cancel func()) {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signalChan
		log.Info("Interrupt signal received. Stopping...")
		cancel()
	}()
}
