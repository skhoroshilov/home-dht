package main

// TempHumData is MeasurmentData implementation for
// sending temperature and humidity
type TempHumData struct {
	temperature float32
	humidity    float32
}

// Name is a name of measurement data
func (data TempHumData) Name() string {
	return "data"
}

// Data returns tags and fields with temperature and humidity
func (data TempHumData) Data() (tags map[string]string, fields map[string]interface{}) {
	fields = map[string]interface{}{
		"temperature": data.temperature,
		"humidity":    data.humidity,
	}

	return
}
