package openmeteo_test

import (
	"fmt"
	"github.com/innotechdevops/openmeteo"
	"testing"
)

func TestOpenMeteo_Execute(t *testing.T) {
	param := openmeteo.Parameter{
		Latitude:  openmeteo.Float32(13.7248785),
		Longitude: openmeteo.Float32(100.4683022),
		Hourly: &[]string{
			openmeteo.HourlyTemperature2m,
			openmeteo.HourlyRelativeHumidity2m,
			openmeteo.HourlyWindSpeed10m,
		},
		Current: &[]string{
			openmeteo.CurrentTemperature2m,
			openmeteo.CurrentRelativeHumidity2m,
			openmeteo.CurrentWindSpeed10m,
		},
		CurrentWeather: openmeteo.Bool(true),
	}
	m := openmeteo.New()

	actual, err := m.Execute(param)

	if err != nil {
		t.Errorf("Get weather error: %s", err)
	}
	fmt.Println(actual)
}
