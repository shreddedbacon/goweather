package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	//"strconv"
	"math"
	"time"

	"github.com/karalabe/hid"
)

func main() {

	device := hid.Enumerate(6465, 32801)
	dev, err := device[0].Open()
	if err != nil {
		log.Fatalln(err)
	}

	s := "a1"
	s2 := "20a1000020"
	data, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	data2, err := hex.DecodeString(s2)
	if err != nil {
		panic(err)
	}
	bfr := make([]byte, 0)
	bfr1 := Read(dev, 0x00, 0x100, data, data2)
	//bfr2 := Read(dev, 0x00 ,0x100, data, data2)
	bfr = append(bfr, bfr1...)
	//bfr = append(bfr, bfr2...)
	log.Println(len(bfr))
	log.Println(bfr)
	log.Println("-------------")
	ReturnJson(bfr)
	//ReturnWSJson(bfr2)
}

func Read(dev *hid.Device, address int, offset int, data []byte, data2 []byte) []byte {
	v1 := byte(address / offset)
	v2 := byte(address % offset)
	bfr := make([]byte, 0)
	fulldata := make([]byte, 0)
	fulldata = append(fulldata, data...)
	fulldata = append(fulldata, v1)
	fulldata = append(fulldata, v2)
	fulldata = append(fulldata, data2...)
	//log.Println(fulldata)
	chunks := offset / 32

	for i := 0; i <= chunks; i++ {
		if fulldata[1] > 0 {
			break
		}
		//fmt.Println(fulldata)
		_, err := dev.Write(fulldata)
		if err != nil {
			log.Fatalln(err)
		}
		time.Sleep(50 * time.Millisecond)
		for b := 0; b <= 3; b++ {
			// we need to loop to get all the bytes from our write/read operation, for some reason we can't get all 32 bytes at once
			buf := make([]byte, 8)
			_, err := dev.Read(buf)
			if err != nil {
				log.Fatalln(err)
			}
			bfr = append(bfr, buf...)
			time.Sleep(50 * time.Millisecond)
		}
		//log.Println(bfr)
		address += 0x20
		fulldata[1] = byte(address / offset)
		fulldata[2] = byte(address % offset)
	}
	return bfr

}

func ReturnJson(data []byte) {
	// https://stackoverflow.com/questions/28465098/golang-bitwise-calculation
	retData := &WH1080Data{
		State: State{
			ReadPeriod:       int(data[16]),
			TimeZone:         int(data[24]),
			NoOfReadings:     int(uint(data[28])<<8 | uint(data[27])),
			CurrentPos:       int(uint(data[31])<<8 | uint(data[30])),
			RelativePressure: toFixed(float64(uint(data[33])<<8|uint(data[32]))*0.1, 1),
			AbsolutePressure: toFixed(float64(uint(data[35])<<8|uint(data[34]))*0.1, 1),
			CurrentDateTime:  FromBCD(data[43:48]), //  # YY-MM-DD-HH-MM
		},
		UnitSetting: UnitSetting{
			IndoorTempF:   BitIsSset(int(data[17]), 0),
			OutdoorTempF:  BitIsSset(int(data[17]), 1),
			RainInch:      BitIsSset(int(data[17]), 2),
			PressureHPa:   BitIsSset(int(data[17]), 5),
			PressureInHg:  BitIsSset(int(data[17]), 6),
			PressureMmHg:  BitIsSset(int(data[17]), 7),
			WindspeedMs:   BitIsSset(int(data[18]), 0),
			WindspeedKmh:  BitIsSset(int(data[18]), 1),
			WindspeedKnot: BitIsSset(int(data[18]), 2),
			WindspeedMh:   BitIsSset(int(data[18]), 3),
			WindspeedBft:  BitIsSset(int(data[18]), 4),
		},
		DisplayOption: DisplayOption{
			PressureRelative: BitIsSset(int(data[19]), 0),
			WindSpeedGust:    BitIsSset(int(data[19]), 1),
			Time12:           BitIsSset(int(data[19]), 2),
			DateMonthDayYear: BitIsSset(int(data[19]), 3),
			TimeScale24Hour:  BitIsSset(int(data[19]), 4),
			DateShowYear:     BitIsSset(int(data[19]), 5),
			DateShowDayName:  BitIsSset(int(data[19]), 6),
			DateAlarmTime:    BitIsSset(int(data[19]), 7),
			ShowOutdoorTemp:  BitIsSset(int(data[20]), 0),
			ShowOutdoorChill: BitIsSset(int(data[20]), 1),
			ShowOutdoorDew:   BitIsSset(int(data[20]), 2),
			ShowRainHour:     BitIsSset(int(data[20]), 3),
			ShowRainDay:      BitIsSset(int(data[20]), 4),
			ShowRainWeek:     BitIsSset(int(data[20]), 5),
			ShowRainMonth:    BitIsSset(int(data[20]), 6),
			ShowRainTotal:    BitIsSset(int(data[20]), 7),
		},
		AlarmEnable: AlarmEnable{
			Time:                BitIsSset(int(data[21]), 1),
			WindDirection:       BitIsSset(int(data[21]), 2),
			IndoorHumidityLow:   BitIsSset(int(data[21]), 4),
			IndoorHumidityHigh:  BitIsSset(int(data[21]), 5),
			OutdoorHumidityLow:  BitIsSset(int(data[21]), 6),
			OutdoorHumidityHigh: BitIsSset(int(data[21]), 7),
			WindAverage:         BitIsSset(int(data[22]), 0),
			WindGust:            BitIsSset(int(data[22]), 1),
			RainHourly:          BitIsSset(int(data[22]), 2),
			RainDaily:           BitIsSset(int(data[22]), 3),
			AbsPressureLow:      BitIsSset(int(data[22]), 4),
			AbsPressureHigh:     BitIsSset(int(data[22]), 5),
			RelPressureLow:      BitIsSset(int(data[22]), 6),
			RelPressureHigh:     BitIsSset(int(data[22]), 7),
			IndoorTempLow:       BitIsSset(int(data[23]), 0),
			IndoorTempHigh:      BitIsSset(int(data[23]), 1),
			OutdoorTempLow:      BitIsSset(int(data[23]), 2),
			OutdoorTempHigh:     BitIsSset(int(data[23]), 3),
			WindChillLow:        BitIsSset(int(data[23]), 4),
			WindChillHigh:       BitIsSset(int(data[23]), 5),
			DewPointLow:         BitIsSset(int(data[23]), 6),
			DewPointHigh:        BitIsSset(int(data[23]), 7),
		},
		Alarm: Alarm{
			IndoorHumidityHigh:   int(data[48]),
			IndoorHumidityLow:    int(data[49]),
			IndoorTempHigh:       BytesToShort(data[51], data[50]) * 0.1,
			IndoorTempLow:        BytesToShort(data[53], data[52]) * 0.1,
			OutdoorHumidityHigh:  int(data[54]),
			OutdoorHumidityLow:   int(data[55]),
			OutdoorTempHigh:      BytesToShort(data[57], data[56]) * 0.1,
			OutdoorTempLow:       BytesToShort(data[59], data[58]) * 0.1,
			WindchillHigh:        BytesToShort(data[61], data[60]) * 0.1,
			WindchillLow:         BytesToShort(data[63], data[61]) * 0.1,
			DewpointHigh:         BytesToShort(data[65], data[64]) * 0.1,
			DewpointLow:          BytesToShort(data[67], data[66]) * 0.1,
			AbsolutePressureHigh: toFixed(float64(uint(data[69])<<8|uint(data[68]))*0.1, 1),
			AbsolutePressureLow:  toFixed(float64(uint(data[71])<<8|uint(data[70]))*0.1, 1),
			RelativePressureHigh: toFixed(float64(uint(data[73])<<8|uint(data[72]))*0.1, 1),
			RelativePressureLow:  toFixed(float64(uint(data[75])<<8|uint(data[74]))*0.1, 1),
			AverageWindSpeedBft:  int(data[76]),
			AverageWindSpeedMs:   toFixed(float64(uint(data[78])<<8|uint(data[77]))*0.1, 1),
			GustWindSpeedBft:     int(data[79]),
			GustWindSpeedMs:      toFixed(float64(uint(data[81])<<8|uint(data[80]))*0.1, 1),
			WindDirection:        toFixed(float64(data[82])*22.5, 1),
			RainHourly:           toFixed(float64(uint(data[84])<<8|uint(data[83]))*0.1, 1),
			RainDaily:            toFixed(float64(uint(data[86])<<8|uint(data[85]))*0.1, 1),
			Time:                 FromBCD(data[87:89]), // HH:MM
		},
		Minmax: Minmax{
			MaxIndoorHumidity:      int(data[98]),
			MinIndoorHumidity:      int(data[99]),
			MaxOutdoorHumidity:     int(data[100]),
			MinOutdoorHumidity:     int(data[101]),
			MaxIndoorTemp:          BytesToShort(data[103], data[102]) * 0.1,
			MinIndoorTemp:          BytesToShort(data[105], data[104]) * 0.1,
			MaxOutdoorTemp:         BytesToShort(data[107], data[106]) * 0.1,
			MinOutdoorTemp:         BytesToShort(data[109], data[108]) * 0.1,
			MaxWindChill:           BytesToShort(data[111], data[110]) * 0.1,
			MinWindChill:           BytesToShort(data[113], data[112]) * 0.1,
			MaxDewPoint:            BytesToShort(data[115], data[114]) * 0.1,
			MinDewPoint:            BytesToShort(data[117], data[116]) * 0.1,
			MaxAbsPressure:         toFixed(float64(uint(data[119])<<8|uint(data[118]))*0.1, 1),
			MinAbsPressure:         toFixed(float64(uint(data[121])<<8|uint(data[120]))*0.1, 1),
			MaxRelPressure:         toFixed(float64(uint(data[123])<<8|uint(data[122]))*0.1, 1),
			MinRelPressure:         toFixed(float64(uint(data[125])<<8|uint(data[124]))*0.1, 1),
			MaxAverageWindSpeed:    toFixed(float64(uint(data[127])<<8|uint(data[126]))*0.1, 1),
			MaxGustWindSpeed:       toFixed(float64(uint(data[129])<<8|uint(data[128]))*0.1, 1),
			MaxRainHourly:          toFixed(float64(uint(data[131])<<8|uint(data[130]))*0.1, 1),
			MaxRainDaily:           toFixed(float64(uint(data[133])<<8|uint(data[132]))*0.1, 1),
			MaxRainWeekly:          toFixed(float64(uint(data[135])<<8|uint(data[134]))*0.1, 1),
			MaxRainMonthly:         toFixed(float64(uint(data[137])<<8|uint(data[136]))*0.1, 1),
			MaxRainTotal:           toFixed(float64(uint(data[139])<<8|uint(data[138]))*0.1, 1),
			MaxRainMonthNibble:     int(data[140] >> 4),
			MaxRainTotalNibble:     int(data[140] & 0x0F),
			MaxIndoorHumidityDate:  FromBCD(data[141:146]),
			MinIndoorHumidityDate:  FromBCD(data[146:151]),
			MaxOutdoorHumidityDate: FromBCD(data[151:156]),
			MinOutdoorHumidityDate: FromBCD(data[156:161]),
			MaxIndoorTempDate:      FromBCD(data[161:166]),
			MinIndoorTempDate:      FromBCD(data[166:171]),
			MaxOutdoorTempDate:     FromBCD(data[171:176]),
			MinOutdoorTempDate:     FromBCD(data[176:181]),
			MaxWindChillDate:       FromBCD(data[181:186]),
			MinWindChillDate:       FromBCD(data[186:191]),
			MaxDewPointDate:        FromBCD(data[191:196]),
			MinDewPointDate:        FromBCD(data[196:201]),
			MaxAbsPressureDate:     FromBCD(data[201:206]),
			MinAbsPressureDate:     FromBCD(data[206:211]),
			MaxRelPressureDate:     FromBCD(data[211:216]),
			MinRelPressureDate:     FromBCD(data[216:221]),
			MaxAveWindSpeedDate:    FromBCD(data[221:226]),
			MaxGustWindSpeedDate:   FromBCD(data[226:231]),
			MaxRainHourlyDate:      FromBCD(data[231:236]),
			MaxRainDailyDate:       FromBCD(data[236:241]),
			MaxRainWeeklyDate:      FromBCD(data[241:246]),
			MaxRainMonthlyDate:     FromBCD(data[246:251]),
			MaxRainTotalDate:       FromBCD(data[251:256]),
		},
	}
	js, _ := json.Marshal(retData)
	fmt.Println(string(js))
}

// check bit is set / true/false
func BitIsSset(b int, bit uint) bool {
	val := b & (1 << bit)
	return (val > 0)
}

// print float to string
func Float2DecimalString(f float64) string {
	return fmt.Sprintf("%.2f", f)
}

// turn bytes to signed short
func BytesToShort(byte_high byte, byte_low byte) float64 {
	sign := 0
	if byte_high == 0x80 {
		sign = -1
	} else {
		sign = +1
	}
	retVal := float64(sign * int(uint16(byte_high&0x7F)<<8|uint16(byte_low)))
	return retVal
}

// turn bytes to ints
func FromBCD(bcd []byte) []int {
	retBytes := make([]int, 0)
	for i := 0; i < len(bcd); i++ {
		retBytes = append(retBytes, int((bcd[i]>>4)*10+(bcd[i]&0x0F)))
	}
	return retBytes
}

// turn uint to string
func UintToString(i uint) string {
	return fmt.Sprintf("%v", i)
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func ReturnWSJson(data []byte) {
	retString := `{
    "delay":` + string(int(data[0])) + `,
    "indoor_humidity":` + string(int(data[1])) + `
    }`
	fmt.Println(retString)
	// return {'delay': byte_str[0],
	//         'indoor_humidity': byte_str[1],
	//         'indoor_temp': to_signed_short(byte_str[3], byte_str[2]) * 0.1,
	//         'outdoor_humidity': byte_str[4],
	//         'outdoor_temp': to_signed_short(byte_str[6], byte_str[5]) * 0.1,
	//         'abs_pressure': (byte_str[8] << 8 | byte_str[7]) * 0.1,
	//         'ave_wind_speed': (((byte_str[11] & 0x0F) << 8) | byte_str[9]) * 0.1,
	//         'gust_wind_speed': (((byte_str[11] & 0xF0) << 8) | byte_str[10]) * 0.1,
	//         'wind_dir': byte_str[12] * 22.5,
	//         'rain_total': (byte_str[14] << 8 | byte_str[13]) * 0.3,
	//         'status_LOC': (byte_str[15] & 0x40) != 0,
	//         'status_RCO': (byte_str[15] & 0x80) != 0}
}
