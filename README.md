# GoWeather

A way to read data from a USB weatherstation based on the FineOffset WH1080 or compatible device.

This is meant to be used as a way to read data from the device and return something that can be used for other applications

# Usage
Use this to get data from a weather station, you need to make sure that you allow read access to the WH1080 device via USB by adding the following udev rule
```
# cat /etc/udev/rules.d/99-wh1080.rules 
SUBSYSTEM=="usb", ATTRS{idVendor}=="1941", ATTRS{idProduct}=="8021", MODE="0666"
```
## Example program
```
package main

import (
	"fmt"
	"log"
	"encoding/json"
	weather "github.com/shreddedbacon/goweather"
)

func main() {
	wh1080, err := weather.New() //open the usb connection and hold it here
	if err != nil {
		log.Fatalln(err)
	}	
	serialBufferMain := wh1080.Read(0x00, 0x100)
	mainData := wh1080.ReturnMainData(serialBufferMain)
	serialBufferCurrent := wh1080.Read(mainData.State.CurrentPos, 0x20)
	currentData := wh1080.ReturnCurrentData(serialBufferCurrent, mainData.State.CurrentPos)
	fullData := &weather.FullData{
		CurrentData: *currentData,
		WH1080Data:  *mainData,
	}
	jsonData, _ := json.Marshal(fullData)
	fmt.Println(string(jsonData))
}
```
## Output
```
{
	"current_data": {
		"indoor_humidity": 41,
		"indoor_temp": 21.5,
		"status_RCO": false,
		"wind_dir": 315,
		"cursor": 51312,
		"status_LOC": false,
		"ave_wind_speed": 0,
		"outdoor_humidity": 94,
		"gust_wind_speed": 0,
		"delay": 1,
		"abs_pressure": 972.9,
		"time": 0,
		"rain_total": 1887.3,
		"outdoor_temp": 4.4,
		"time_str": "2019-06-27T11:14:18Z"
	},
	"main_data": {
		"display_option": {
			"show_rain_month": false,
			"date_alarm_time": false,
			"show_outdoor_temp": true,
			"date_month-day-year": false,
			"show_rain_hour": true,
			"show_rain_day": false,
			"time_scale_24_hour": false,
			"show_outdoor_chill": false,
			"show_rain_total": false,
			"wind_speed_gust": false,
			"pressure_relative": true,
			"show_rain_week": false,
			"date_show_year": true,
			"show_outdoor_dew": false,
			"time_12": false,
			"date_show_day_name": false
		},
		"alarm": {
			"absolute_pressure_high": 1040,
			"outdoor_temp_low": -10,
			"gust_wind_speed_ms": 36,
			"average_wind_speed_ms": 18,
			"wind_direction": 0,
			"outdoor_humidity_high": 70,
			"windchill_low": 0,
			"absolute_pressure_low": 960,
			"outdoor_humidity_low": 45,
			"relative_pressure_high": 1040,
			"indoor_temp_high": 20,
			"rain_daily": 50,
			"gust_wind_speed_bft": 0,
			"average_wind_speed_bft": 0,
			"dewpoint_low": -10,
			"indoor_temp_low": 0,
			"windchill_high": 20,
			"indoor_humidity_high": 65,
			"indoor_humidity_low": 35,
			"relative_pressure_low": 960,
			"dewpoint_high": 10,
			"rain_hourly": 1,
			"time": [12, 0],
			"outdoor_temp_high": 30
		},
		"state": {
			"no_of_readings": 3192,
			"current_pos": 51312,
			"absolute_pressure": 972.9,
			"time_zone": 0,
			"read_Period": 5,
			"relative_pressure": 1025.5,
			"current_date_time": [19, 6, 27, 21, 12]
		},
		"minmax": {
			"min_wind_chill": -7,
			"min_indoor_humidity": 11,
			"min_outdoor_temp_date": [18, 7, 21, 6, 23],
			"min_abs_pressure": 933.6,
			"max_dew_point_date": [18, 1, 19, 16, 54],
			"max_indoor_temp": 36.4,
			"max_outdoor_humidity": 99,
			"min_dew_point": -11.8,
			"max_rain_total": 585.3,
			"min_outdoor_temp": -7,
			"max_rain_month_nibble": 0,
			"max_rain_weekly": 64.5,
			"max_rain_daily": 35.1,
			"max_abs_pressure_date": [10, 1, 1, 12, 0],
			"min_rel_pressure_date": [16, 9, 30, 5, 15],
			"max_wind_chill": 44.400000000000006,
			"max_rain_daily_date": [15, 1, 11, 13, 38],
			"min_outdoor_humidity": 10,
			"min_outdoor_humidity_date": [10, 1, 1, 12, 0],
			"min_rel_pressure": 991.7,
			"max_rain_monthly": 99.9,
			"min_indoor_temp": 40,
			"max_abs_pressure": 1018.3,
			"max_rel_pressure": 1042.6,
			"max_rain_monthly_date": [17, 11, 14, 10, 19],
			"max_wind_chill_date": [39, 1, 14, 13, 51],
			"max_rain_hourly_date": [15, 1, 24, 13, 1],
			"max_rain_hourly": 848.3,
			"max_outdoor_temp": 44.400000000000006,
			"max_rel_pressure_date": [15, 10, 8, 8, 48],
			"max_rain_total_nibble": 0,
			"max_outdoor_humidity_date": [15, 11, 26, 6, 33],
			"max_indoor_temp_date": [15, 3, 11, 16, 53],
			"max_outdoor_temp_date": [19, 1, 14, 13, 51],
			"max_ave_wind_speed_date": [18, 6, 2, 21, 1],
			"max_rain_weekly_date": [10, 1, 6, 20, 46],
			"max_indoor_humidity_date": [17, 3, 21, 10, 53],
			"max_gust_wind_speed": 54.4,
			"max_rain_total_date": [15, 12, 26, 10, 20],
			"min_dew_point_date": [90, 1, 1, 12, 0],
			"max_dew_point": 37,
			"min_abs_pressure_date": [16, 9, 30, 5, 15],
			"max_average_wind_speed": 43.5,
			"max_gust_wind_speed_date": [17, 1, 20, 18, 30],
			"min_indoor_temp_date": [15, 1, 24, 10, 19],
			"min_wind_chill_date": [18, 7, 21, 6, 23],
			"max_indoor_humidity": 92,
			"min_indoor_humidity_date": [14, 12, 16, 16, 42]
		},
		"alarm_enable": {
			"outdoor_temp_low": false,
			"wind_average": false,
			"wind_direction": false,
			"outdoor_humidity_high": false,
			"abs_pressure_low": false,
			"outdoor_humidity_low": false,
			"indoor_temp_high": false,
			"rel_pressure_high": false,
			"rain_daily": false,
			"abs_pressure_high": false,
			"wind_chill_low": false,
			"wind_chill_high": false,
			"dew_point_low": false,
			"dew_point_high": false,
			"indoor_temp_low": false,
			"wind_gust": false,
			"indoor_humidity_high": false,
			"indoor_humidity_low": false,
			"rain_hourly": false,
			"time": false,
			"outdoor_temp_high": false,
			"rel_pressure_low": false
		},
		"unit_setting": {
			"windspeed_knot": false,
			"windspeed_bft": false,
			"outdoor_Temp_F": false,
			"rain_inch": false,
			"pressure_inHg": false,
			"pressure_mmHg": false,
			"indoor_Temp_F": false,
			"pressure_hPa": true,
			"windspeed_ms": false,
			"windspeed_kmh": true,
			"windspeed_mh": false
		}
	}
}
```
