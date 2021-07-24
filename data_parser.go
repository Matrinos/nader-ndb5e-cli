package main

import (
	"encoding/hex"
	"strconv"
)

// Declare a new type named status which will unify our enum values
// It has an underlying type of unsigned integer (uint).
type DataTag int

// Declare typed constants each with type of status
const (
	SimpleFormat DataTag = iota
	FullFormat
)

// String returns the string value of the status
func (s DataTag) String() string {
	strings := [...]string{"SimpleFormat", "FullFormat"}

	if s < SimpleFormat || s > FullFormat {
		return "Unknown"
	}

	return strings[s]
}

type MegTag int

const (
	Product MegTag = iota
	Operational
	Status
	Data
	ProtectedParameters
	Log
	RemoteParameters
)

// String returns the string value of the status
func (s MegTag) String() string {
	strings := [...]string{"Product", "Operational", "Status", "Data", "ProtectedParameters", "Log", "RemoteParameters"}

	if s < Product || s > RemoteParameters {
		return "Unknown"
	}

	return strings[s]
}

// const (
// 	ProductLen    = 20
// 	DataAddrLe    = 4
// 	DataNumberLen = 2
// 	DataLen       = 52
// )

// type (
// 	fileInfo struct {
// 		FileName string
// 		HostID   string
// 	}
// )

type (
	ProductInfomation struct {
		Model        string
		ShellCurrent uint16
		Ampacity     uint16
		Usage        uint16
		RatedVoltage uint16
		Exception    uint16
	}
)

func SplitProductData(data string) error {
	addr := data[:2]

	var p = new(ProductInfomation)
	ParseProductInfo(data[2:42], p)

	dataAddr := data[42:46]

	dataLen, err := hexStringToDec(data[46:48])
	if err != nil {
		return err
	}

	return nil
}

func hexStringToDec(num string) (uint16, error) {
	result, err := strconv.ParseUint(num, 16, 16)
	if err != nil {
		return 0, err
	}

	return uint16(result), nil
}

func ParseProductInfo(data string, product *ProductInfomation) error {
	brandName := data[:20]
	decoded, err := hex.DecodeString(brandName)
	if err != nil {
		return err
	}

	product.Model = string(decoded)

	decValue, err := hexStringToDec(data[20:24])
	if err != nil {
		return err
	}
	product.ShellCurrent = decValue

	decValue, err = hexStringToDec(data[24:28])
	if err != nil {
		return err
	}
	product.Ampacity = decValue

	decValue, err = hexStringToDec(data[28:32])
	if err != nil {
		return err
	}
	product.Usage = decValue

	decValue, err = hexStringToDec(data[32:36])
	if err != nil {
		return err
	}
	product.RatedVoltage = decValue

	decValue, err = hexStringToDec(data[36:40])
	if err != nil {
		return err
	}
	product.Exception = decValue

	return nil
}
