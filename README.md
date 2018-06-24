# home-dht

[![Build Status](https://travis-ci.org/skhoroshilov/home-dht.svg?branch=master)](https://travis-ci.org/skhoroshilov/home-dht)
[![Go Report Card](https://goreportcard.com/badge/github.com/skhoroshilov/home-dht)](https://goreportcard.com/report/github.com/skhoroshilov/home-dht)
[![GoDoc](https://godoc.org/github.com/skhoroshilov/home-dht?status.svg)](https://godoc.org/github.com/skhoroshilov/home-dht)

A simple app that gathers temperature and humidity from DHT22 sensor and sends it to influxdb for using at grafana dashboards.

![Dashboard example](doc/images/example.png)

## Requirements

* DHT22 temperature and humidity sensor
* linux pc (e.g. Raspberry Pi) for gathering data from DHT22 sensor
* ![influxdb](https://github.com/influxdata/influxdb)
* ![grafana](https://github.com/grafana/grafana)

## Building home-dht

`make`

You need golang installed to build this app. This app can be built only on linux machine because the library used for reading DHT22 data can be build properly only for linux. If you build the app on windows it will use fake DHT22 data.

## Getting Started

1. Install and run influxdb [![influxdb installation](https://github.com/influxdata/influxdb#installation)]

2. Install and run grafana [![grafana installation](https://github.com/grafana/grafana#installation)]

3. Create dashboard in grafana. You can import ![existing dashboard](third_party/grafana/dashboard.json) or create your own. The app uses influxdb `autogen.data` database, fields `temperature` and `humidity`.

4. Start home-dht app

`./home-dht <dht22 pin number> <influxdb address>`
