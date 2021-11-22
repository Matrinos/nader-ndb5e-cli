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
const SWITCHRECORD_ADDR = 360
const RECORD_INFO_LEN = 16

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
		FaultRecord           uint16
		FaultReadNo           uint16
		FaultCategory         uint16
		_                     uint16
		FaultRecordParameter1 uint16
		_                     uint16
		_                     uint16
		_                     uint16
		_                     uint16
		_                     uint16
		_                     uint16
		_                     uint16
		_                     uint16
		FaultYearMonth        uint16
		FaultDayHour          uint16
		FaultMinuteSecond     uint16
	}

	RemoteControlParameter struct {
		//OperateCmd      uint16
		//_               uint16
		//_               uint16
		//_               uint16
		//_               uint16
		//_               uint16
		//_               uint16
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
		//_               uint16
		//_               uint16
		//_               uint16
		//_               uint16
		//OneKeySwitch    uint16
		//SwitchRegister1 uint16
		//SwitchRegister2 uint16
		//SwitchRegister3 uint16
		//SwitchRegister4 uint16
	}

	RecordLogs struct {
		Log [14]uint16
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

	return json.Marshal(p)
}

func (p *RemoteControlParameter) ToJson() ([]byte, error) {

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

func (p *RecordLogs) ToJson() ([]byte, error) {

	return json.Marshal(p)
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

func GetRemoteCtlSetting(jsonfile string, Params *RemoteControlParameter) error {
	//var strJson string = "{\"TimeOffDay0\":[\"Monday\",\"Sunday\"],\"TimeOffTime0\":\"15:40:34\",\"TimeOnDay0\":[\"Monday\",\"Sunday\"],\"TimeOnTime0\":\"15:41:34\",\"TimeOffDay1\":[\"Monday\",\"Sunday\"],\"TimeOffTime1\":\"15:42:34\",\"TimeOnDay1\":[\"Monday\",\"Sunday\"],\"TimeOnTime1\":\"15:43:34\",\"TimeOffDay2\":[\"Monday\",\"Sunday\"],\"TimeOffTime2\":\"15:44:34\",\"TimeOnDay2\":[\"Monday\",\"Sunday\"],\"TimeOnTime2\":\"15:45:34\",\"TimeOffDay3\":[\"Monday\",\"Sunday\"],\"TimeOffTime3\":\"15:46:34\",\"TimeOnDay3\":[\"Monday\",\"Sunday\"],\"TimeOnTime3\":\"15:47:34\",\"TimeOffDay4\":[\"Monday\",\"Sunday\"],\"TimeOffTime4\":\"15:48:34\",\"TimeOnDay4\":[\"Monday\",\"Sunday\"],\"TimeOnTime4\":\"15:49:34\"}"
	data, err := ioutil.ReadFile(jsonfile)

	if err != nil {
		return err
	}

	jsonMap := make(map[string]interface{})

	//json.Unmarshal([]byte(strJson), &jsonMap)
	json.Unmarshal(data, &jsonMap)

	//var Params RemoteControlParameter
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
		Params.TimeOffDH0 |= GROUPFULL_FLAG
	}
	return err
}
