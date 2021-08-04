package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {

	var p = Product{

		Specification:           1,
		Model:                   2,
		ShellCurrent:            5,
		Ampacity:                4,
		Usage:                   1,
		RatedVoltage:            2250,
		Frequency:               500,
		Type:                    2,
		Functions:               3,
		Poles:                   4,
		FaultProtection:         3,
		GroundLeakageProtection: 4,
		SlaveId:                 100,
		BaudRate:                9600,
		ManufactureYear:         33,
		ManufactureMonthDay:     0x0610,
		SerialNumber:            [20]uint8{97, 98, 99, 100, 101},
	}

	data, err := p.ToJson()
	if err != nil {
		t.Error()
	}

	jsonMap := make(map[string]interface{})
	json.Unmarshal(data, &jsonMap)

	assert.Equal(t, "00210610", jsonMap[MANUFACTURE_DATE], "they should be equal")
	assert.Equal(t, "abcde", fmt.Sprintf("%v", jsonMap["SerialNumber"]))
}
