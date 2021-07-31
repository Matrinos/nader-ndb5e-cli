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

	GateWay struct {
	}
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
