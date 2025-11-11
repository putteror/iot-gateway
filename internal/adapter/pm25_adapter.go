package adapter

import (
	"fmt"
	"math/rand/v2"
)

type GetPM25 interface {
	GetPM25Value() (float64, float64, float64, float64, error)
}

type GetPM25Impl struct {
}

func NewGetPM25Impl() GetPM25 {
	return &GetPM25Impl{}
}

func (s *GetPM25Impl) GetPM25Value() (float64, float64, float64, float64, error) {

	minPm25 := 29.0
	maxPm25 := 33.0

	diff := maxPm25 - minPm25
	randomValue := rand.Float64()*diff + minPm25

	var pm25 float64
	_, err := fmt.Sscanf(fmt.Sprintf("%.2f", randomValue), "%f", &pm25)
	if err != nil {
		return 0.0, 0.0, 0.0, 0.0, err
	}

	minPm10 := 58.0
	maxPm10 := 62.0

	pm10Diff := maxPm10 - minPm10
	randomPm10Value := rand.Float64()*pm10Diff + maxPm10

	var pm10 float64
	_, err = fmt.Sscanf(fmt.Sprintf("%.2f", randomPm10Value), "%f", &pm10)
	if err != nil {
		return 0.0, 0.0, 0.0, 0.0, err
	}

	minHumidity := 60.0
	maxHumidity := 70.0

	humidityDiff := maxHumidity - minHumidity
	randomHumidityValue := rand.Float64()*humidityDiff + maxHumidity

	var humidity float64
	_, err = fmt.Sscanf(fmt.Sprintf("%.2f", randomHumidityValue), "%f", &humidity)
	if err != nil {
		return 0.0, 0.0, 0.0, 0.0, err
	}

	minTemperature := 30.0
	maxTemperature := 31.0

	temperatureDiff := maxTemperature - minTemperature
	randomTemperatureValue := rand.Float64()*temperatureDiff + maxTemperature

	var tempeerature float64
	_, err = fmt.Sscanf(fmt.Sprintf("%.2f", randomTemperatureValue), "%f", &tempeerature)
	if err != nil {
		return 0.0, 0.0, 0.0, 0.0, err
	}

	return pm25, pm10, humidity, tempeerature, nil

}
