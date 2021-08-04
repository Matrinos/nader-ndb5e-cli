package main

import (
	"encoding/binary"
	"fmt"

	"github.com/yerden/go-util/bcd"
)

const PRODUCT_ADDR = 0x200 // modbus start address
const PRODUCT_LEN = 26     // number of registers

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

	MultipleTenuinnt uint16
)

func (p Product) SerialNumberStr() string {
	return string(p.SerialNumber[:])
}

func (p Product) ManufactureDate() (string, error) {
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
