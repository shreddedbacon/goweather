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
	usbDevice := hid.Enumerate(0x1941, 0x8021)
	wh1080, err := usbDevice[0].Open()
	if err != nil {
		log.Fatalln(err)
	}

	fullData := CollectData(wh1080)
	jsonFullData, _ := json.Marshal(fullData)
	fmt.Println(string(jsonFullData))
}

func CollectData(wh1080 *hid.Device) *FullData {
	serialBufferMain := Read(wh1080, 0x00, 0x100)
	mainData := ReturnMainData(serialBufferMain)
	serialBufferCurrent := Read(wh1080, mainData.State.CurrentPos, 0x20)
	currentData := ReturnCurrentData(serialBufferCurrent, mainData.State.CurrentPos)

	fullData := &FullData{
		CurrentData: *currentData,
		WH1080Data:  *mainData,
	}
	return fullData
}

func Read(dev *hid.Device, address int, offset int) []byte {
	data, _ := hex.DecodeString("a1")
	data2, _ := hex.DecodeString("20a1000020")
	v1 := byte(address / 0x100)
	v2 := byte(address % 0x100)
	bfr := make([]byte, 0)
	fulldata := make([]byte, 0)
	fulldata = append(fulldata, data...)
	fulldata = append(fulldata, v1)
	fulldata = append(fulldata, v2)
	fulldata = append(fulldata, data2...)
	chunks := offset / 32
	for i := 0; i < chunks+1; i++ {
		_, err := dev.Write(fulldata)
		if err != nil {
			log.Fatalln(err)
		}
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
		address += 0x20
		fulldata[1] = byte(address / 0x100)
		fulldata[2] = byte(address % 0x100)
	}
	return bfr

}

// return maindata
func ReturnMainData(data []byte) *WH1080Data {
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
	return retData
}

// Return CurrentData
func ReturnCurrentData(data []byte, cursor int) *CurrentData {
	retData := &CurrentData{
		IndoorHumidity:  int(data[1]),
		IndoorTemp:      BytesToShort(data[3], data[2]) * 0.1,
		OutdoorHumidity: int(data[4]),
		OutdoorTemp:     BytesToShort(data[6], data[5]) * 0.1,
		AbsPressure:     toFixed(float64(uint(data[8])<<8|uint(data[7]))*0.1, 1),
		AveWindSpeed:    toFixed(float64(uint(data[11]&0x0F)<<8|uint(data[9]))*0.1, 1),
		GustWindSpeed:   toFixed(float64(uint(data[11]&0xF0)<<8|uint(data[10]))*0.1, 1),
		WindDir:         toFixed(float64(data[12])*22.5, 1),
		RainTotal:       toFixed(float64(uint(data[14])<<8|uint(data[13]))*0.3, 1),
		StatusRCO:       (data[15] & 0x80) != 0,
		StatusLOC:       (data[15] & 0x40) != 0,
		Delay:           int(data[0]),
		Cursor:          cursor,
		TimeStr:         time.Now().UTC().Format(time.RFC3339),
	}
	return retData
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
