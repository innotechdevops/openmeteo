# Open-Meteo for Golang

[Open-Meteo](https://open-meteo.com/en/docs) is an open-source weather API with free access for non-commercial use. No API key is required. You can use it immediately!

## Install

```shell
go get github.com/innotechdevops/openmeteo
```

## How to use

API Document https://open-meteo.com/en/docs

```go
param := openmeteo.Parameter{
    Latitude:       openmeteo.Float32(13.7248785),
    Longitude:      openmeteo.Float32(100.4683022),
    Hourly:         &[]string{
        openmeteo.HourlyTemperature2m,
        openmeteo.HourlyRelativeHumidity2m,
        openmeteo.HourlyWindSpeed10m,
    },
    CurrentWeather: openmeteo.Bool(true),
}

m := openmeteo.New()
resp, err := m.Execute(param)
```
Output

```json
{
  "latitude":13.75,
  "longitude":100.5,
  "generationtime_ms":0.18894672393798828,
  "utc_offset_seconds":0,
  "timezone":"GMT",
  "timezone_abbreviation":"GMT",
  "elevation":5.0,
  "current_weather":{
    "temperature":30.6,
    "windspeed":5.6,
    "winddirection":195.0,
    "weathercode":3,
    "is_day":0,
    "time":"2023-06-02T15:00"
  },
  "hourly_units":{
    "time":"iso8601",
    "temperature_2m":"°C",
    "relativehumidity_2m":"%",
    "windspeed_10m":"km/h"
  },
  "hourly":{
    "time":["2023-06-02T00:00"],
    "temperature_2m":[28.8],
    "relativehumidity_2m":[77],
    "windspeed_10m":[3.3]
  }
}
```