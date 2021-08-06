package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"

	"github.com/yerden/go-util/bcd"
)

const PRODUCT_ADDR = 0x200 // modbus start address
const PRODUCT_LEN = 26     // number of registers
const MANUFACTURE_DATE = "ManufactureDate"

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

	Data struct {
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
		Temprature          int16
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

	Logs1 struct {
		FaultLogs1 [20]uint16
	}

	Summary1 struct {
		Temprature [144]int8
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

func (p *Logs1) ToJson() ([]byte, error) {

	return json.Marshal(p)
}

func (d *Data) ToJson() ([]byte, error) {
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
