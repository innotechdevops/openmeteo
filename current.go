package openmeteo

// Note: Current conditions are based on 15-minutely weather model data.
// Every weather variable available in hourly data, is available as current condition as well.
const (
	CurrentIsDayOrNight             = "is_day"
	CurrentSnow                     = "snowfall"
	CurrentTemperature2m            = "temperature_2m"
	CurrentRelativeHumidity2m       = "relativehumidity_2m"
	CurrentDewPoint2m               = "dewpoint_2m"
	CurrentApparentTemperature      = "apparent_temperature"
	CurrentPressureMsl              = "pressure_msl"
	CurrentSurfacePressure          = "surface_pressure"
	CurrentCloudCover               = "cloudcover"
	CurrentCloudCoverLow            = "cloudcover_low"
	CurrentCloudCoverMid            = "cloudcover_mid"
	CurrentCloudCoverHigh           = "cloudcover_high"
	CurrentWindSpeed10m             = "windspeed_10m"
	CurrentWindSpeed80m             = "windspeed_80m"
	CurrentWindSpeed120m            = "windspeed_120m"
	CurrentWindSpeed180m            = "windspeed_180m"
	CurrentWindDirection10m         = "winddirection_10m"
	CurrentWindDirection80m         = "winddirection_80m"
	CurrentWindDirection120m        = "winddirection_120m"
	CurrentWindDirection180m        = "winddirection_180m"
	CurrentWindGusts10m             = "windgusts_10m"
	CurrentShortwaveRadiation       = "shortwave_radiation"
	CurrentDirectRadiation          = "direct_radiation"
	CurrentDirectNormalIrradiance   = "direct_normal_irradiance"
	CurrentDiffuseRadiation         = "diffuse_radiation"
	CurrentVaporPressureDeficit     = "vapor_pressure_deficit"
	CurrentCape                     = "cape"
	CurrentEvapotranspiration       = "evapotranspiration"
	CurrentEt0FaoEvapotranspiration = "et0_fao_evapotranspiration"
	CurrentPrecipitation            = "precipitation"
	CurrentSnowfall                 = "snowfall"
	CurrentPrecipitationProbability = "precipitation_probability"
	CurrentRain                     = "rain"
	CurrentShowers                  = "showers"
	CurrentWeatherCode              = "weather_code"
	CurrentSnowDepth                = "snow_depth"
	CurrentFreezingLevelHeight      = "freezinglevel_height"
	CurrentVisibility               = "visibility"
	CurrentSoilTemperature0cm       = "soil_temperature_0cm"
	CurrentSoilTemperature6cm       = "soil_temperature_6cm"
	CurrentSoilTemperature18cm      = "soil_temperature_18cm"
	CurrentSoilTemperature54cm      = "soil_temperature_54cm"
	CurrentSoilMoisture01cm         = "soil_moisture_0_1cm"
	CurrentSoilMoisture13cm         = "soil_moisture_1_3cm"
	CurrentSoilMoisture39cm         = "soil_moisture_3_9cm"
	CurrentSoilMoisture927cm        = "soil_moisture_9_27cm"
	CurrentSoilMoisture2781cm       = "soil_moisture_27_81cm"
	CurrentIsDay                    = "is_day"
)
