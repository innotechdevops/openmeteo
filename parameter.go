package openmeteo

import (
	"fmt"
	"reflect"
	"strings"
)

type Parameter struct {
	Latitude          *float32  `json:"latitude"`
	Longitude         *float32  `json:"longitude"`
	Hourly            *[]string `json:"hourly"`
	Elevation         *float32  `json:"elevation"`
	Daily             *[]string `json:"daily"`
	Current           *[]string `json:"current"`
	CurrentWeather    *bool     `json:"current_weather"`
	TemperatureUnit   *string   `json:"temperature_unit"`
	WindSpeedUnit     *string   `json:"windspeed_unit"`
	PrecipitationUnit *string   `json:"precipitation_unit"`
	TimeFormat        *string   `json:"timeformat"`
	Timezone          *string   `json:"timezone"`
	PastDays          *int      `json:"past_days"`
	ForecastDays      *int      `json:"forecast_days"`
	StartDate         *string   `json:"start_date"`
	EndDate           *string   `json:"end_date"`
	Models            *[]string `json:"models"`
	CellSelection     *string   `json:"cell_selection"`
}

func (p Parameter) ToQuery() string {
	v := reflect.ValueOf(p)
	var queryParams []string

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		value := v.Field(i)

		if value.IsNil() {
			continue
		}

		fieldName := field.Tag.Get("json")
		fieldValue := fmt.Sprintf("%v", value.Elem().Interface())

		// Check if the field value is a slice
		if value.Elem().Kind() == reflect.Slice {
			slice := value.Elem()
			var sliceValues []string
			for j := 0; j < slice.Len(); j++ {
				sliceValue := fmt.Sprintf("%v", slice.Index(j).Interface())
				sliceValues = append(sliceValues, sliceValue)
			}
			fieldValue = strings.Join(sliceValues, ",")
		}

		queryParam := fieldName + "=" + fieldValue
		queryParams = append(queryParams, queryParam)
	}

	return strings.Join(queryParams, "&")
}
