package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/yerden/go-util/bcd"
)

const PRODUCT_ADDR = 0x200 // modbus start address
const PRODUCT_LEN = 26     // number of registers
const MANUFACTURE_DATE = "ManufactureDate"

const FAULTRECORD_ADDR = 0x340
const ALARMRECORD_ADDR = 0x350
const SWITCHRECORD_ADDR = 0x360
const RECORD_INFO_LEN = 16
const FAULTRECORD_NUM_ADDR = 0x341
const ALARMRECORD_NUM_ADDR = 0x351
const SWITCHRECORD_NUM_ADDR = 0x361

const FAULT_TYPE = 0
const ALARM_TYPE = 1
const SWITCH_TYPE = 2

const OPPARAMETERS_ADDR = 0x240
const OPPARAMETERS_LEN = 3

const RUNSTATUS_ADDR = 0x250
const RUNSTATUS_LEN = 14

const METRICALDATA_ADDR = 0x0260
const METRICALDATA_LEN = 80

const PRETECTPARAMETERS_ADDR = 0x0300
const PRETECTPARAMETERS_LEN = 46

const REMOTECONTROL_ADDR = 0x0407
const REMOTECONTROL_LEN = 20 //36

const FE_TEMPERATURES_ADDR = 0x0500
const FE_TEMPERATURES_LEN = 72

const FE_ENERGYPERHOUR_ADDR = 0x0548
const FE_ENERGYPERHOUR_LEN = 48

const FE_ENERGYPERDAY_ADDR = 0x0578
const FE_ENERGYPERDAY_LEN = 62

const FE_ENERGYPERMONTH_ADDR = 0x05B6
const FE_ENERGYPERMONTH_LEN = 24

const FAULTRECORDLOG_ADDR = 0x0800
const ALARMRECORDLOG_ADDR = 0x0918
const SWITCHRECORDLOG_ADDR = 0x0A30
const MAX_FAULTRECORDLOG_NUM = 20
const MAX_ALARMRECORDLOG_NUM = 20
const MAX_SWITCHRECORDLOG_NUM = 30
const RECORD_LOG_LEN = 14
const LOGSINGROUP_NUM = 5

const TIMER_MONDAY = 0x01
const TIMER_TUESDAY = TIMER_MONDAY << 1
const TIMER_WEDNESDAY = TIMER_MONDAY << 2
const TIMER_THURSDAY = TIMER_MONDAY << 3
const TIMER_FIRDAY = TIMER_MONDAY << 4
const TIMER_SATURDAY = TIMER_MONDAY << 5
const TIMER_SUNDAY = TIMER_MONDAY << 6

const GROUPFULL_FLAG = 0x8000

type (
	Product struct {
		Specification           uint16
		Model                   uint16
		ShellCurrent            uint16
		Ampacity                uint16
		Usage                   uint16
		RatedVoltage            uint16
		Frequency               uint16
		Type                    uint16
		Functions               uint16
		Poles                   uint16
		FaultProtection         uint16
		GroundLeakageProtection uint16
		SlaveId                 uint16
		BaudRate                uint16
		ManufactureYear         uint16
		ManufactureMonthDay     uint16
		SerialNumber            [20]uint8
	}

	OpParameters struct {
		YearMonth    uint16
		DayHour      uint16
		MinuteSecond uint16
	}

	RunStatus struct {
		OperateTime          uint16
		ContactWear          uint16
		SelfDiagnosticAlarm  uint16
		AlarmStatus          uint16
		RunStatus1           uint16
		RunStatus2           uint16
		AlarmIncreaseing     uint16
		_                    uint16
		_                    uint16
		BrakePositionChanges uint16
		_                    uint16
		CurrentTripTimes     uint16
		OtherTripTimes       uint16
		_                    uint16
	}

	GateWay struct {
	}

	MetricalData struct {
		ACurrent            uint32
		BCurrent            uint32
		CCurrent            uint32
		_                   uint32
		_                   uint16
		_                   uint16
		_                   uint32
		_                   uint32
		_                   uint32
		_                   uint32
		_                   uint32
		_                   uint32
		AVoltage            uint16
		BVoltage            uint16
		CVoltage            uint16
		ABVoltage           uint16
		BCVoltage           uint16
		CAVoltage           uint16
		_                   uint16
		_                   uint16
		AFrequency          uint16
		BFrequency          uint16
		CFrequency          uint16
		PhaseState          uint16
		ARealPower          uint16
		BRealPower          uint16
		CRealPower          uint16
		RealPowerTotal      uint16
		AReactivePower      uint16
		BReactivePower      uint16
		CReactivePower      uint16
		ReactivePowerTotal  uint16
		AApparentPower      uint16
		BApparentPower      uint16
		CApparentPower      uint16
		ApparentPowerTotal  uint16
		_                   uint16
		_                   uint16
		_                   uint16
		TotalUsagePower     uint16
		APowerFactor        uint16
		BPowerFactor        uint16
		CPowerFactor        uint16
		TotalPowerFactor    uint16
		AActiveEnergy       uint32
		BActiveEnergy       uint32
		CActiveEnergy       uint32
		ActiveEnergyTotal   uint32
		AReactiveEnergy     uint32
		BReactiveEnergy     uint32
		CReactiveEnergy     uint32
		ReactiveEnergyTotal uint32
		AApparentEnergy     uint32
		BApparentEnergy     uint32
		CApparentEnergy     uint32
		ApparentEnergyTotal uint32
		Temperature         int16
		LeakageCurrent      uint16
	}

	ProtectParameters struct {
		CurrentSettingValue    uint16
		TimeSettingValue       uint16
		_                      uint16
		_                      uint16
		_                      uint16
		_                      uint16
		_                      uint16
		_                      uint16
		_                      uint16
		_                      uint16
		_                      uint16
		_                      uint16
		_                      uint16
		_                      uint16
		_                      uint16
		_                      uint16
		_                      uint16
		TotalPowerSettingValue uint16
		_                      uint16
		_                      uint16
		_                      uint16
		_                      uint16
		_                      uint16
		PhaseSequenceSwitch    uint16
		_                      uint16
		_                      uint16
		_                      uint16
		_                      uint16
		_                      uint16
		_                      uint16
		_                      uint16
		_                      uint16
		_                      uint16
		_                      uint16
		TemperatureSwitch      uint16
		_                      uint16
		_                      uint16
		_                      uint16
		_                      uint16
		OverVoltageSwitch      uint16
		UnderVoltageSwitch     uint16
		CurrentSwitch          uint16
		_                      uint16
		PhaseLossSwitch        uint16
		_                      uint16
		RequiredValueSwitch    uint16
	}

	Record struct {
		Record       uint16
		ReadNo       uint16
		Category     uint16
		_            uint16
		RecordParams uint16
		_            uint16
		_            uint16
		_            uint16
		_            uint16
		_            uint16
		_            uint16
		_            uint16
		_            uint16
		YearMonth    uint16
		DayHour      uint16
		MinuteSecond uint16
	}

	RecordJson struct {
		Record       uint16
		RecordNo     uint16
		RecordType   uint16
		RecordCode   uint16
		RecordParams uint16
		Description  string
		Date         string
		Time         string
	}

	TimerControlParameter struct {
		TimeOffDH0 uint16
		TimeOffMS0 uint16
		TimeOnDH0  uint16
		TimeOnMS0  uint16
		TimeOffDH1 uint16
		TimeOffMS1 uint16
		TimeOnDH1  uint16
		TimeOnMS1  uint16
		TimeOffDH2 uint16
		TimeOffMS2 uint16
		TimeOnDH2  uint16
		TimeOnMS2  uint16
		TimeOffDH3 uint16
		TimeOffMS3 uint16
		TimeOnDH3  uint16
		TimeOnMS3  uint16
		TimeOffDH4 uint16
		TimeOffMS4 uint16
		TimeOnDH4  uint16
		TimeOnMS4  uint16
	}

	TimerControlJson struct {
		TimeOffDay0  []string
		TimeOffTime0 string
		TimeOnDay0   []string
		TimeOnTime0  string
		TimeOffDay1  []string
		TimeOffTime1 string
		TimeOnDay1   []string
		TimeOnTime1  string
		TimeOffDay2  []string
		TimeOffTime2 string
		TimeOnDay2   []string
		TimeOnTime2  string
		TimeOffDay3  []string
		TimeOffTime3 string
		TimeOnDay3   []string
		TimeOnTime3  string
		TimeOffDay4  []string
		TimeOffTime4 string
		TimeOnDay4   []string
		TimeOnTime4  string
	}

	RecordLogsInfo struct {
		LogNo   uint16
		LogType uint16
		Logs    RecordLogs
	}
	RecordLogs struct {
		LogRecord    uint16
		_            uint16
		LogParams    uint16
		_            uint16
		_            uint16
		_            uint16
		_            uint16
		_            uint16
		_            uint16
		_            uint16
		_            uint16
		YearMonth    uint16
		DayHour      uint16
		MinuteSecond uint16
	}

	Summary1 struct {
		Temperature [144]int8
	}

	Summary2 struct {
		ElectricEnergyPerHour [24]uint32
	}

	Summary3 struct {
		ElectricEnergyPerDay [31]uint32
	}

	Summary4 struct {
		ElectricEnergyPerMonth [12]uint32
	}

	MultipleTenuinnt uint16
)

type JsonMarshal interface {
	ToJson() ([]byte, error)
}

func (p *Product) SerialNumberStr() string {
	return string(bytes.Trim(p.SerialNumber[:], "\x00"))
}

func (p *Product) ManufactureDate() (string, error) {
	year, err := UintToBCDString(p.ManufactureYear)
	if err != nil {
		return "", err
	}

	monthDay, err := UintToBCDString(p.ManufactureMonthDay)
	if err != nil {
		return "", err
	}
	result := fmt.Sprintf("%s%s",
		year,
		monthDay)

	return result, nil
}

func (p *Product) ToJson() ([]byte, error) {
	jp, _ := json.Marshal(p)
	var m map[string]interface{}
	json.Unmarshal(jp, &m)

	date, err := p.ManufactureDate()
	if err != nil {
		//TODO: handle error correctly
		return []byte{}, err
	}

	m[MANUFACTURE_DATE] = date

	for key := range m {
		if key == "SerialNumber" {
			m[key] = p.SerialNumberStr()
		}
		if key == "RatedVoltage" {
			m[key] = float32(p.RatedVoltage) * 0.1
		}
	}

	return json.Marshal(m)
}

func (p *OpParameters) ToJson() ([]byte, error) {

	return json.Marshal(p)
}

func (p *RunStatus) ToJson() ([]byte, error) {

	return json.Marshal(p)
}

func (p *ProtectParameters) ToJson() ([]byte, error) {

	return json.Marshal(p)
}

func (p *Record) ToJson() ([]byte, error) {

	rJson := RecordJson{}

	recordType := (p.Category >> 5) & 0x7
	recordCode := (p.Category & 0x1F)

	rJson.Record = p.Record
	rJson.RecordNo = p.ReadNo
	rJson.RecordCode = recordCode
	rJson.RecordType = recordType
	rJson.RecordParams = p.RecordParams
	rJson.Description = GetRecordDescription(p.ReadNo, recordType, recordCode, p.RecordParams)
	rJson.Date = GetDate(p.YearMonth, p.DayHour)
	rJson.Time = GetTime(p.DayHour, p.MinuteSecond)

	return json.Marshal(rJson)
}

func (p *TimerControlParameter) ToJson() ([]byte, error) {

	var TimeJson TimerControlJson
	TimeJson.TimeOffDay0 = GetDay(p.TimeOffDH0)
	TimeJson.TimeOffTime0 = GetTime(p.TimeOffDH0, p.TimeOffMS0)
	TimeJson.TimeOnDay0 = GetDay(p.TimeOnDH0)
	TimeJson.TimeOnTime0 = GetTime(p.TimeOnDH0, p.TimeOnMS0)

	TimeJson.TimeOffDay1 = GetDay(p.TimeOffDH1)
	TimeJson.TimeOffTime1 = GetTime(p.TimeOffDH1, p.TimeOffMS1)
	TimeJson.TimeOnDay1 = GetDay(p.TimeOnDH1)
	TimeJson.TimeOnTime1 = GetTime(p.TimeOnDH1, p.TimeOnMS1)

	TimeJson.TimeOffDay2 = GetDay(p.TimeOffDH2)
	TimeJson.TimeOffTime2 = GetTime(p.TimeOffDH2, p.TimeOffMS2)
	TimeJson.TimeOnDay2 = GetDay(p.TimeOnDH2)
	TimeJson.TimeOnTime2 = GetTime(p.TimeOnDH2, p.TimeOnMS2)

	TimeJson.TimeOffDay3 = GetDay(p.TimeOffDH3)
	TimeJson.TimeOffTime3 = GetTime(p.TimeOffDH3, p.TimeOffMS3)
	TimeJson.TimeOnDay3 = GetDay(p.TimeOnDH3)
	TimeJson.TimeOnTime3 = GetTime(p.TimeOnDH3, p.TimeOnMS3)

	TimeJson.TimeOffDay4 = GetDay(p.TimeOffDH4)
	TimeJson.TimeOffTime4 = GetTime(p.TimeOffDH4, p.TimeOffMS4)
	TimeJson.TimeOnDay4 = GetDay(p.TimeOnDH4)
	TimeJson.TimeOnTime4 = GetTime(p.TimeOnDH4, p.TimeOnMS4)
	return json.Marshal(TimeJson)
}

func (p *TimerControlJson) ToJson() ([]byte, error) {

	return json.Marshal(p)
}

func (p *Summary1) ToJson() ([]byte, error) {

	return json.Marshal(p)
}

func (p *Summary2) ToJson() ([]byte, error) {

	return json.Marshal(p)
}

func (p *Summary3) ToJson() ([]byte, error) {

	return json.Marshal(p)
}

func (p *Summary4) ToJson() ([]byte, error) {

	return json.Marshal(p)
}

func (p *RecordLogsInfo) ToJson() ([]byte, error) {

	rJson := RecordJson{}

	LogNo := p.LogNo
	LogType := p.LogType
	LogNum := ((p.Logs.LogRecord >> 8) & 0xFF)
	LogCode := (p.Logs.LogRecord & 0x1F)
	LogParams := p.Logs.LogParams

	rJson.Record = LogNum
	rJson.RecordNo = LogNo
	rJson.RecordCode = LogCode
	rJson.RecordType = LogType
	rJson.RecordParams = LogParams
	rJson.Description = GetRecordDescription(LogNo, LogType, LogCode, LogParams)
	rJson.Date = GetDate(p.Logs.YearMonth, p.Logs.DayHour)
	rJson.Time = GetTime(p.Logs.DayHour, p.Logs.MinuteSecond)

	return json.Marshal(rJson)
}

func (d *MetricalData) ToJson() ([]byte, error) {
	jp, _ := json.Marshal(d)
	var m map[string]interface{}
	json.Unmarshal(jp, &m)

	//TODO: custom logic here.
	return json.Marshal(m)
}

func UintToBCDString(data uint16) (string, error) {

	dec := bcd.NewDecoder(bcd.Standard)
	src := make([]byte, 2)
	binary.BigEndian.PutUint16(src, data)

	dst := make([]byte, bcd.DecodedLen(len(src)))

	n, err := dec.Decode(dst, src)
	if err != nil {
		return "", err
	}

	return string(dst[:n]), nil
}

func UintToBCD(data uint16) uint16 {

	return (((data & 0x3f) / 10) << 4) | ((data & 0x3f) % 10)
}

func BCDToUint(data uint16) uint16 {

	return ((data>>4)*10 + (data & 0x0f))
}

func GetDayHour(arrDay []interface{}, strTime string) (uint16, error) {
	mapDays := map[string]uint16{
		"Monday":    TIMER_MONDAY,
		"Tuesday":   TIMER_TUESDAY,
		"Wednesday": TIMER_WEDNESDAY,
		"Thursday":  TIMER_THURSDAY,
		"Firday":    TIMER_FIRDAY,
		"Saturday":  TIMER_SATURDAY,
		"Sunday":    TIMER_SUNDAY,
	}

	var nDay uint16 = 0
	for i := 0; i < len(arrDay); i++ {
		strDay := arrDay[i].(string)
		if v, ok := mapDays[strDay]; ok {
			fmt.Println(v)
			nDay |= v
		}
	}

	var h, m, s uint16 = 0, 0, 0
	_, err := fmt.Sscanf(strTime, "%d:%d:%d", &h, &m, &s)

	if err != nil {
		return 0, err
	}

	nDay <<= 8
	nDay |= UintToBCD(h)
	return nDay, err
}

func GetMinute(strTime string) (uint16, error) {
	var h, m, s uint16 = 0, 0, 0
	_, err := fmt.Sscanf(strTime, "%d:%d:%d", &h, &m, &s)
	if err != nil {
		return 0, err
	}

	var nMinute uint16 = 0
	nMinute = UintToBCD(m)

	return (nMinute << 8), err
}

func GetDay(DayHour uint16) []string {
	mapDays := map[uint16]string{
		0: "Monday",
		1: "Tuesday",
		2: "Wednesday",
		3: "Thursday",
		4: "Firday",
		5: "Saturday",
		6: "Sunday",
	}

	days := (DayHour >> 8) & 0x7F

	buf := make([]string, 0, 7)

	for i := uint16(0); i < 7; i++ {
		if (days & 0x1) != 0 {
			buf = append(buf, mapDays[i])
		}
		days >>= 1
	}
	return buf
}

func GetDate(YearMonth uint16, DayHour uint16) string {

	year := BCDToUint(YearMonth >> 8)
	month := BCDToUint(YearMonth & 0xFF)
	Hour := BCDToUint(DayHour >> 8)
	return fmt.Sprintf("20%02d-%02d-%02d", year, month, Hour)
}

func GetTime(DayHour uint16, MinuteSecond uint16) string {
	hour := DayHour & 0xFF
	min := (MinuteSecond >> 8) & 0xFF
	second := MinuteSecond & 0xFF

	if hour > 0x23 || min > 0x59 {
		return ""
	}

	return fmt.Sprintf("%02d:%02d:%02d", BCDToUint(hour), BCDToUint(min), BCDToUint(second))
}

func GetRemoteCtlSetting(jsonfile string, Params *TimerControlParameter) error {
	//var strJson string = "{\"TimeOffDay0\":[\"Monday\",\"Sunday\"],\"TimeOffTime0\":\"15:40:34\",\"TimeOnDay0\":[\"Monday\",\"Sunday\"],\"TimeOnTime0\":\"15:41:34\",\"TimeOffDay1\":[\"Monday\",\"Sunday\"],\"TimeOffTime1\":\"15:42:34\",\"TimeOnDay1\":[\"Monday\",\"Sunday\"],\"TimeOnTime1\":\"15:43:34\",\"TimeOffDay2\":[\"Monday\",\"Sunday\"],\"TimeOffTime2\":\"15:44:34\",\"TimeOnDay2\":[\"Monday\",\"Sunday\"],\"TimeOnTime2\":\"15:45:34\",\"TimeOffDay3\":[\"Monday\",\"Sunday\"],\"TimeOffTime3\":\"15:46:34\",\"TimeOnDay3\":[\"Monday\",\"Sunday\"],\"TimeOnTime3\":\"15:47:34\",\"TimeOffDay4\":[\"Monday\",\"Sunday\"],\"TimeOffTime4\":\"15:48:34\",\"TimeOnDay4\":[\"Monday\",\"Sunday\"],\"TimeOnTime4\":\"15:49:34\"}"
	data, err := ioutil.ReadFile(jsonfile)

	if err != nil {
		return err
	}

	jsonMap := make(map[string]interface{})

	//json.Unmarshal([]byte(strJson), &jsonMap)
	json.Unmarshal(data, &jsonMap)

	//var Params TimerControlParameter
	var bAllGroups bool = true
	//Group 0
	if jsonMap["TimeOffDay0"] != nil {
		nDayHour, err1 := GetDayHour(jsonMap["TimeOffDay0"].([]interface{}), jsonMap["TimeOffTime0"].(string))
		dMinute, err2 := GetMinute(jsonMap["TimeOffTime0"].(string))
		if err1 != nil || err2 != nil {
			Params.TimeOffDH0 = 0
			Params.TimeOffMS0 = 0
			bAllGroups = false
		} else {
			Params.TimeOffDH0 = nDayHour
			Params.TimeOffMS0 = dMinute
		}
	} else {
		Params.TimeOffDH0 = 0
		Params.TimeOffMS0 = 0
		bAllGroups = false
	}

	if jsonMap["TimeOnDay0"] != nil {
		nDayHour, err1 := GetDayHour(jsonMap["TimeOnDay0"].([]interface{}), jsonMap["TimeOnTime0"].(string))
		dMinute, err2 := GetMinute(jsonMap["TimeOnTime0"].(string))
		if err1 != nil || err2 != nil {
			Params.TimeOnDH0 = 0
			Params.TimeOnMS0 = 0
			bAllGroups = false
		} else {
			Params.TimeOnDH0 = nDayHour
			Params.TimeOnMS0 = dMinute
		}
	} else {
		Params.TimeOnDH0 = 0
		Params.TimeOnMS0 = 0
		bAllGroups = false
	}
	//Group 1
	if jsonMap["TimeOffDay1"] != nil {
		nDayHour, err1 := GetDayHour(jsonMap["TimeOffDay1"].([]interface{}), jsonMap["TimeOffTime1"].(string))
		dMinute, err2 := GetMinute(jsonMap["TimeOffTime1"].(string))
		if err1 != nil || err2 != nil {
			Params.TimeOffDH1 = 0
			Params.TimeOffMS1 = 0
			bAllGroups = false
		} else {
			Params.TimeOffDH1 = nDayHour
			Params.TimeOffMS1 = dMinute
		}

	} else {
		Params.TimeOffDH1 = 0
		Params.TimeOffMS1 = 0
		bAllGroups = false
	}

	if jsonMap["TimeOnDay1"] != nil {
		nDayHour, err1 := GetDayHour(jsonMap["TimeOnDay1"].([]interface{}), jsonMap["TimeOnTime1"].(string))
		dMinute, err2 := GetMinute(jsonMap["TimeOnTime1"].(string))
		if err1 != nil || err2 != nil {
			Params.TimeOnDH1 = 0
			Params.TimeOnMS1 = 0
			bAllGroups = false
		} else {
			Params.TimeOnDH1 = nDayHour
			Params.TimeOnMS1 = dMinute
		}
	} else {
		Params.TimeOnDH1 = 0
		Params.TimeOnMS1 = 0
		bAllGroups = false
	}

	//Group 2
	if jsonMap["TimeOffDay2"] != nil {
		nDayHour, err1 := GetDayHour(jsonMap["TimeOffDay2"].([]interface{}), jsonMap["TimeOffTime2"].(string))
		dMinute, err2 := GetMinute(jsonMap["TimeOffTime2"].(string))
		if err1 != nil || err2 != nil {
			Params.TimeOffDH2 = 0
			Params.TimeOffMS2 = 0
			bAllGroups = false
		} else {
			Params.TimeOffDH2 = nDayHour
			Params.TimeOffMS2 = dMinute
		}
	} else {
		Params.TimeOffDH2 = 0
		Params.TimeOffMS2 = 0
		bAllGroups = false
	}

	if jsonMap["TimeOnDay2"] != nil {
		nDayHour, err1 := GetDayHour(jsonMap["TimeOnDay2"].([]interface{}), jsonMap["TimeOnTime2"].(string))
		dMinute, err2 := GetMinute(jsonMap["TimeOnTime2"].(string))
		if err1 != nil || err2 != nil {
			Params.TimeOnDH2 = 0
			Params.TimeOnMS2 = 0
			bAllGroups = false
		} else {
			Params.TimeOnDH2 = nDayHour
			Params.TimeOnMS2 = dMinute
		}
	} else {
		Params.TimeOnDH2 = 0
		Params.TimeOnMS2 = 0
		bAllGroups = false
	}

	//Group 3
	if jsonMap["TimeOffDay3"] != nil {
		nDayHour, err1 := GetDayHour(jsonMap["TimeOffDay3"].([]interface{}), jsonMap["TimeOffTime3"].(string))
		dMinute, err2 := GetMinute(jsonMap["TimeOffTime3"].(string))
		if err1 != nil || err2 != nil {
			Params.TimeOffDH3 = 0
			Params.TimeOffMS3 = 0
			bAllGroups = false
		} else {
			Params.TimeOffDH3 = nDayHour
			Params.TimeOffMS3 = dMinute
		}
	} else {
		Params.TimeOffDH3 = 0
		Params.TimeOffMS3 = 0
		bAllGroups = false
	}

	if jsonMap["TimeOnDay3"] != nil {
		nDayHour, err1 := GetDayHour(jsonMap["TimeOnDay3"].([]interface{}), jsonMap["TimeOnTime3"].(string))
		dMinute, err2 := GetMinute(jsonMap["TimeOnTime3"].(string))
		if err1 != nil || err2 != nil {
			Params.TimeOnDH3 = 0
			Params.TimeOnMS3 = 0
			bAllGroups = false
		} else {
			Params.TimeOnDH3 = nDayHour
			Params.TimeOnMS3 = dMinute
		}
	} else {
		Params.TimeOnDH3 = 0
		Params.TimeOnMS3 = 0
		bAllGroups = false
	}

	//Group 4
	if jsonMap["TimeOffDay4"] != nil {
		nDayHour, err1 := GetDayHour(jsonMap["TimeOffDay4"].([]interface{}), jsonMap["TimeOffTime4"].(string))
		dMinute, err2 := GetMinute(jsonMap["TimeOffTime4"].(string))
		if err1 != nil || err2 != nil {
			Params.TimeOffDH4 = 0
			Params.TimeOffMS4 = 0
			bAllGroups = false
		} else {
			Params.TimeOffDH4 = nDayHour
			Params.TimeOffMS4 = dMinute
		}
	} else {
		Params.TimeOffDH4 = 0
		Params.TimeOffMS4 = 0
		bAllGroups = false
	}

	if jsonMap["TimeOnDay4"] != nil {
		nDayHour, err1 := GetDayHour(jsonMap["TimeOnDay4"].([]interface{}), jsonMap["TimeOnTime4"].(string))
		dMinute, err2 := GetMinute(jsonMap["TimeOnTime4"].(string))
		if err1 != nil || err2 != nil {
			Params.TimeOnDH4 = 0
			Params.TimeOnMS4 = 0
			bAllGroups = false
		} else {
			Params.TimeOnDH4 = nDayHour
			Params.TimeOnMS4 = dMinute
		}
	} else {
		Params.TimeOnDH4 = 0
		Params.TimeOnMS4 = 0
		bAllGroups = false
	}

	if bAllGroups {
		// TODO: need check with vendor to know the mearning of this bit
		//Params.TimeOffDH0 |= GROUPFULL_FLAG  //? sometimes exception error happans when set 15bit as 1, illegal data.
	}
	return err
}

func GetRecordDescription(RecordNo uint16, RecordType uint16, RecordCode uint16, RecordParams uint16) string {
	mapFaultCode := map[uint16]string{
		2:  "漏电故障",     //00010-漏电故障
		5:  "过载长延时",    //00101-过载长延时
		6:  "瞬时",       //00110-瞬时
		7:  "缺相",       //00111-缺相
		8:  "欠压",       //01000-欠压
		9:  "过压",       //01001-过压
		12: "相序",       //01100-相序
		26: "IU 电流不平衡", //11010-IU 电流不平衡
		27: "功率需用值保护",  //11011-功率需用值保护
		30: "温度保护",     //11110-温度保护
	}

	mapAlarmCode := map[uint16]string{
		1:  "漏电自检",     //00001-漏电自检
		3:  "功率需用值预报警", //00011-功率需用值预报警
		21: "过载预报警",    //10101-过载预报警
		31: "温度预报警",    //11111-温度预报警
	}

	mapCauseCode := map[uint16]string{
		1:  "漏电自检",     //00001-漏电自检
		2:  "漏电故障",     //00010-漏电故障
		5:  "过载长延时",    //00101-过载长延时
		6:  "瞬时",       //00110-瞬时
		7:  "缺相",       //00111-缺相
		8:  "欠压",       //01000-欠压
		9:  "过压",       //01001-过压
		12: "相序",       //01100-相序
		18: "手动分/合闸",   //10010-手动分/合闸
		20: "定时分/合闸",   //10100-定时分/合闸
		24: "重合闸",      //11000-重合闸
		25: "打火(预留)",   //11001-打火(预留)
		26: "IU 电流不平衡", //11010-IU 电流不平衡
		27: "功率需用值保护",  //11011-功率需用值保护
		29: "远程分/合闸",   //11101-远程分/合闸
		30: "温度保护",     //11110-温度保护

	}

	Description := string("")

	if RecordType == FAULT_TYPE {
		v, ok := mapFaultCode[RecordCode]
		if ok {
			Description = fmt.Sprintf("故障记录第%d条: %s", RecordNo, v)
		} else {
			Description = fmt.Sprintf("故障记录第%d条: 故障不明", RecordNo)
		}
	} else if RecordType == ALARM_TYPE {
		v, ok := mapAlarmCode[RecordCode]
		if ok {
			Description = fmt.Sprintf("报警记录第%d条: %s", RecordNo, v)
		} else {
			Description = fmt.Sprintf("报警记录第%d条: 报警不明", RecordNo)
		}
	} else if RecordType == SWITCH_TYPE {
		causeCode := (RecordParams & 0x1F)
		status := (RecordParams >> 5) & 0x1
		statusStr := string("成功")
		if status == 0 {
			statusStr = string("失败")
		}
		if RecordCode == 0x1 { //switch on
			Description = fmt.Sprintf("变位记录第%d条: 合闸, 变位%s", RecordNo, statusStr)
		} else if RecordCode == 0x2 { //switch off
			v, ok := mapCauseCode[causeCode]
			if ok {
				Description = fmt.Sprintf("变位记录第%d条: 分闸, 变位%s, 分闸原因:%s", RecordNo, statusStr, v)
			} else {
				Description = fmt.Sprintf("变位记录第%d条: 分闸, 变位%s, 分闸原因不明", RecordNo, statusStr)
			}
		}

	} else {
		return string("")
	}
	return Description
}
