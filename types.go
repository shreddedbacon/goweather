package goweather

type State struct {
	NoOfReadings     int     `json:"no_of_readings"`
	CurrentPos       int     `json:"current_pos"`
	AbsolutePressure float64 `json:"absolute_pressure"`
	TimeZone         int     `json:"time_zone"`
	ReadPeriod       int     `json:"read_Period"`
	RelativePressure float64 `json:"relative_pressure"`
	CurrentDateTime  []int   `json:"current_date_time"`
}

type DisplayOption struct {
	ShowRainMonth    bool `json:"show_rain_month"`
	DateAlarmTime    bool `json:"date_alarm_time"`
	ShowOutdoorTemp  bool `json:"show_outdoor_temp"`
	DateMonthDayYear bool `json:"date_month-day-year"`
	ShowRainHour     bool `json:"show_rain_hour"`
	ShowRainDay      bool `json:"show_rain_day"`
	TimeScale24Hour  bool `json:"time_scale_24_hour"`
	ShowOutdoorChill bool `json:"show_outdoor_chill"`
	ShowRainTotal    bool `json:"show_rain_total"`
	WindSpeedGust    bool `json:"wind_speed_gust"`
	PressureRelative bool `json:"pressure_relative"`
	ShowRainWeek     bool `json:"show_rain_week"`
	DateShowYear     bool `json:"date_show_year"`
	ShowOutdoorDew   bool `json:"show_outdoor_dew"`
	Time12           bool `json:"time_12"`
	DateShowDayName  bool `json:"date_show_day_name"`
}

type Alarm struct {
	AbsolutePressureHigh float64 `json:"absolute_pressure_high"`
	OutdoorTempLow       float64 `json:"outdoor_temp_low"`
	GustWindSpeedMs      float64 `json:"gust_wind_speed_ms"`
	AverageWindSpeedMs   float64 `json:"average_wind_speed_ms"`
	WindDirection        float64 `json:"wind_direction"`
	OutdoorHumidityHigh  int     `json:"outdoor_humidity_high"`
	WindchillLow         float64 `json:"windchill_low"`
	AbsolutePressureLow  float64 `json:"absolute_pressure_low"`
	OutdoorHumidityLow   int     `json:"outdoor_humidity_low"`
	RelativePressureHigh float64 `json:"relative_pressure_high"`
	IndoorTempHigh       float64 `json:"indoor_temp_high"`
	RainDaily            float64 `json:"rain_daily"`
	GustWindSpeedBft     int     `json:"gust_wind_speed_bft"`
	AverageWindSpeedBft  int     `json:"average_wind_speed_bft"`
	DewpointLow          float64 `json:"dewpoint_low"`
	IndoorTempLow        float64 `json:"indoor_temp_low"`
	WindchillHigh        float64 `json:"windchill_high"`
	IndoorHumidityHigh   int     `json:"indoor_humidity_high"`
	IndoorHumidityLow    int     `json:"indoor_humidity_low"`
	RelativePressureLow  float64 `json:"relative_pressure_low"`
	DewpointHigh         float64 `json:"dewpoint_high"`
	RainHourly           float64 `json:"rain_hourly"`
	Time                 []int   `json:"time"`
	OutdoorTempHigh      float64 `json:"outdoor_temp_high"`
}
type Minmax struct {
	MinWindChill           float64 `json:"min_wind_chill"`
	MinIndoorHumidity      int     `json:"min_indoor_humidity"`
	MinOutdoorTempDate     []int   `json:"min_outdoor_temp_date"`
	MinAbsPressure         float64 `json:"min_abs_pressure"`
	MaxDewPointDate        []int   `json:"max_dew_point_date"`
	MaxIndoorTemp          float64 `json:"max_indoor_temp"`
	MaxOutdoorHumidity     int     `json:"max_outdoor_humidity"`
	MinDewPoint            float64 `json:"min_dew_point"`
	MaxRainTotal           float64 `json:"max_rain_total"`
	MinOutdoorTemp         float64 `json:"min_outdoor_temp"`
	MaxRainMonthNibble     int     `json:"max_rain_month_nibble"`
	MaxRainWeekly          float64 `json:"max_rain_weekly"`
	MaxRainDaily           float64 `json:"max_rain_daily"`
	MaxAbsPressureDate     []int   `json:"max_abs_pressure_date"`
	MinRelPressureDate     []int   `json:"min_rel_pressure_date"`
	MaxWindChill           float64 `json:"max_wind_chill"`
	MaxRainDailyDate       []int   `json:"max_rain_daily_date"`
	MinOutdoorHumidity     int     `json:"min_outdoor_humidity"`
	MinOutdoorHumidityDate []int   `json:"min_outdoor_humidity_date"`
	MinRelPressure         float64 `json:"min_rel_pressure"`
	MaxRainMonthly         float64 `json:"max_rain_monthly"`
	MinIndoorTemp          float64 `json:"min_indoor_temp"`
	MaxAbsPressure         float64 `json:"max_abs_pressure"`
	MaxRelPressure         float64 `json:"max_rel_pressure"`
	MaxRainMonthlyDate     []int   `json:"max_rain_monthly_date"`
	MaxWindChillDate       []int   `json:"max_wind_chill_date"`
	MaxRainHourlyDate      []int   `json:"max_rain_hourly_date"`
	MaxRainHourly          float64 `json:"max_rain_hourly"`
	MaxOutdoorTemp         float64 `json:"max_outdoor_temp"`
	MaxRelPressureDate     []int   `json:"max_rel_pressure_date"`
	MaxRainTotalNibble     int     `json:"max_rain_total_nibble"`
	MaxOutdoorHumidityDate []int   `json:"max_outdoor_humidity_date"`
	MaxIndoorTempDate      []int   `json:"max_indoor_temp_date"`
	MaxOutdoorTempDate     []int   `json:"max_outdoor_temp_date"`
	MaxAveWindSpeedDate    []int   `json:"max_ave_wind_speed_date"`
	MaxRainWeeklyDate      []int   `json:"max_rain_weekly_date"`
	MaxIndoorHumidityDate  []int   `json:"max_indoor_humidity_date"`
	MaxGustWindSpeed       float64 `json:"max_gust_wind_speed"`
	MaxRainTotalDate       []int   `json:"max_rain_total_date"`
	MinDewPointDate        []int   `json:"min_dew_point_date"`
	MaxDewPoint            float64 `json:"max_dew_point"`
	MinAbsPressureDate     []int   `json:"min_abs_pressure_date"`
	MaxAverageWindSpeed    float64 `json:"max_average_wind_speed"`
	MaxGustWindSpeedDate   []int   `json:"max_gust_wind_speed_date"`
	MinIndoorTempDate      []int   `json:"min_indoor_temp_date"`
	MinWindChillDate       []int   `json:"min_wind_chill_date"`
	MaxIndoorHumidity      int     `json:"max_indoor_humidity"`
	MinIndoorHumidityDate  []int   `json:"min_indoor_humidity_date"`
}

type AlarmEnable struct {
	OutdoorTempLow      bool `json:"outdoor_temp_low"`
	WindAverage         bool `json:"wind_average"`
	WindDirection       bool `json:"wind_direction"`
	OutdoorHumidityHigh bool `json:"outdoor_humidity_high"`
	AbsPressureLow      bool `json:"abs_pressure_low"`
	OutdoorHumidityLow  bool `json:"outdoor_humidity_low"`
	IndoorTempHigh      bool `json:"indoor_temp_high"`
	RelPressureHigh     bool `json:"rel_pressure_high"`
	RainDaily           bool `json:"rain_daily"`
	AbsPressureHigh     bool `json:"abs_pressure_high"`
	WindChillLow        bool `json:"wind_chill_low"`
	WindChillHigh       bool `json:"wind_chill_high"`
	DewPointLow         bool `json:"dew_point_low"`
	DewPointHigh        bool `json:"dew_point_high"`
	IndoorTempLow       bool `json:"indoor_temp_low"`
	WindGust            bool `json:"wind_gust"`
	IndoorHumidityHigh  bool `json:"indoor_humidity_high"`
	IndoorHumidityLow   bool `json:"indoor_humidity_low"`
	RainHourly          bool `json:"rain_hourly"`
	Time                bool `json:"time"`
	OutdoorTempHigh     bool `json:"outdoor_temp_high"`
	RelPressureLow      bool `json:"rel_pressure_low"`
}

type UnitSetting struct {
	WindspeedKnot bool `json:"windspeed_knot"`
	WindspeedBft  bool `json:"windspeed_bft"`
	OutdoorTempF  bool `json:"outdoor_Temp_F"`
	RainInch      bool `json:"rain_inch"`
	PressureInHg  bool `json:"pressure_inHg"`
	PressureMmHg  bool `json:"pressure_mmHg"`
	IndoorTempF   bool `json:"indoor_Temp_F"`
	PressureHPa   bool `json:"pressure_hPa"`
	WindspeedMs   bool `json:"windspeed_ms"`
	WindspeedKmh  bool `json:"windspeed_kmh"`
	WindspeedMh   bool `json:"windspeed_mh"`
}
type WH1080Data struct {
	DisplayOption DisplayOption `json:"display_option"`
	Alarm         Alarm         `json:"alarm"`
	State         State         `json:"state"`
	Minmax        Minmax        `json:"minmax"`
	AlarmEnable   AlarmEnable   `json:"alarm_enable"`
	UnitSetting   UnitSetting   `json:"unit_setting"`
}

type CurrentData struct {
	IndoorHumidity  int     `json:"indoor_humidity"`
	IndoorTemp      float64 `json:"indoor_temp"`
	StatusRCO       bool    `json:"status_RCO"`
	WindDir         float64 `json:"wind_dir"`
	Cursor          int     `json:"cursor"`
	StatusLOC       bool    `json:"status_LOC"`
	AveWindSpeed    float64 `json:"ave_wind_speed"`
	OutdoorHumidity int     `json:"outdoor_humidity"`
	GustWindSpeed   float64 `json:"gust_wind_speed"`
	Delay           int     `json:"delay"`
	AbsPressure     float64 `json:"abs_pressure"`
	Time            float64 `json:"time"`
	RainTotal       float64 `json:"rain_total"`
	OutdoorTemp     float64 `json:"outdoor_temp"`
	TimeStr         string  `json:"time_str"`
}

type FullData struct {
	CurrentData CurrentData `json:"current_data"`
	WH1080Data  WH1080Data  `json:"main_data"`
}
