package openmeteo_test

import (
	"github.com/innotechdevops/openmeteo"
	"testing"
)

func TestParameter_ToQuery(t *testing.T) {
	expected := "hourly=temperature,humidity&daily=temperature_max,precipitation_sum&current=temperature_2m,relativehumidity_2m&temperature_unit=Celsius&precipitation_unit=millimeters&timeformat=YYYY-MM-DD HH:mm:ss&timezone=Europe/London&past_days=7&start_date=2023-06-01"
	param := openmeteo.Parameter{
		Latitude:          nil,
		Longitude:         nil,
		Hourly:            &[]string{"temperature", "humidity"},
		Elevation:         nil,
		Daily:             &[]string{"temperature_max", "precipitation_sum"},
		Current:		   &[]string{"temperature_2m", "relativehumidity_2m"},
		CurrentWeather:    nil,
		TemperatureUnit:   openmeteo.String("Celsius"),
		WindSpeedUnit:     nil,
		PrecipitationUnit: openmeteo.String("millimeters"),
		TimeFormat:        openmeteo.String("YYYY-MM-DD HH:mm:ss"),
		Timezone:          openmeteo.String("Europe/London"),
		PastDays:          openmeteo.Int(7),
		ForecastDays:      nil,
		StartDate:         openmeteo.String("2023-06-01"),
		EndDate:           nil,
		Models:            nil,
		CellSelection:     nil,
	}

	actual := param.ToQuery()

	if actual != expected {
		t.Errorf("Data not match: %s", actual)
	}
}

func BenchmarkParameter_ToQuery(b *testing.B) {
	param := openmeteo.Parameter{
		Latitude:          nil,
		Longitude:         nil,
		Hourly:            &[]string{"temperature", "humidity"},
		Elevation:         nil,
		Daily:             &[]string{"temperature_max", "precipitation_sum"},
		CurrentWeather:    nil,
		TemperatureUnit:   openmeteo.String("Celsius"),
		WindSpeedUnit:     nil,
		PrecipitationUnit: openmeteo.String("millimeters"),
		TimeFormat:        openmeteo.String("YYYY-MM-DD HH:mm:ss"),
		Timezone:          openmeteo.String("Europe/London"),
		PastDays:          openmeteo.Int(7),
		ForecastDays:      nil,
		StartDate:         openmeteo.String("2023-06-01"),
		EndDate:           nil,
		Models:            nil,
		CellSelection:     nil,
	}
	for i := 0; i < b.N; i++ {
		_ = param.ToQuery()
	}
}
