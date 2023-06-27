package openmeteo

func WeatherCodeName(code int) string {
	if code == 0 {
		return "clear-sky"
	} else if code == 1 {
		return "mainly-clear"
	} else if code == 2 {
		return "partly-cloudy"
	} else if code == 3 {
		return "overcast"
	} else if code == 48 {
		return "fog"
	} else if code == 49 {
		return "depositing-rime-fog"
	} else if code == 51 {
		return "drizzle-light"
	} else if code == 53 {
		return "drizzle-moderate"
	} else if code == 55 {
		return "drizzle-dense"
	} else if code == 56 {
		return "freezing-drizzle-light"
	} else if code == 57 {
		return "freezing-drizzle-dense"
	} else if code == 61 {
		return "rain-slight"
	} else if code == 63 {
		return "rain-moderate"
	} else if code == 65 {
		return "rain-heavy"
	} else if code == 66 {
		return "freezing-rain-light"
	} else if code == 67 {
		return "freezing-rain-heavy"
	} else if code == 71 {
		return "snow-fall-slight"
	} else if code == 73 {
		return "snow-fall-moderate"
	} else if code == 75 {
		return "snow-fall-heavy"
	} else if code == 77 {
		return "snow-grains"
	} else if code == 80 {
		return "rain-showers-slight"
	} else if code == 81 {
		return "rain-showers-moderate"
	} else if code == 82 {
		return "rain-showers-violent"
	} else if code == 85 {
		return "snow-showers-slight"
	} else if code == 86 {
		return "snow-showers-heavy"
	} else if code == 95 {
		return "thunderstorm-slight-or-moderate"
	} else if code == 96 || code == 99 {
		return "thunderstorm-slight-and-heavy-hail"
	}
	return "unknown"
}
