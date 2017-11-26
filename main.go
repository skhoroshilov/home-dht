package main

import (
	"context"
	"flag"
	"log"
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
	log.Printf("Using dht22 pin = %v", pin)
	log.Printf("Using influxdb address = %v", influxdbAddress)

	rd := NewReader(pin)
	sender, err := NewSender(influxdbAddress)

	if err != nil {
		log.Fatal(err)
	}

	defer sender.Close()

	job := NewDhtSendingJob(rd, sender)
	job.Start(context.TODO())

	log.Println("Done")
}
