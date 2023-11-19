package interfaces

import (
	// "github.com/the-ress/prusalink-screen/octoprintApis"
	"github.com/the-ress/prusalink-screen/octoprintApis/dataModels"
)

type ITemperatureDataDisplay interface {
	UpdateTemperatureData(temperatureData map[string]dataModels.TemperatureData)
}
