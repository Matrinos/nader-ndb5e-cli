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

func TestTimerControlStruct(t *testing.T) {

	var myTime = TimerControlJson{
		TimeOffDay0:  []string{"Monday", "Sunday"},
		TimeOffTime0: "13:00:34",
		TimeOnDay0:   []string{"Monday", "Tuesday", "Sunday"},
		TimeOnTime0:  "13:01:34",
		TimeOffDay1:  []string{"Monday", "Wednesday", "Sunday"},
		TimeOffTime1: "13:02:34",
		TimeOnDay1:   []string{"Monday", "Thursday", "Sunday"},
		TimeOnTime1:  "13:03:34",
		TimeOffDay2:  []string{"Monday", "Firday", "Sunday"},
		TimeOffTime2: "13:04:34",
		TimeOnDay2:   []string{"Monday", "Saturday", "Sunday"},
		TimeOnTime2:  "13:05:34",
		TimeOffDay3:  []string{"Monday", "Wednesday", "Sunday"},
		TimeOffTime3: "13:06:34",
		TimeOnDay3:   []string{"Monday", "Thursday", "Sunday"},
		TimeOnTime3:  "13:07:34",
		TimeOffDay4:  []string{"Monday", "Tuesday", "Wednesday", "Sunday"},
		TimeOffTime4: "13:08:34",
		TimeOnDay4:   []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Sunday"},
		TimeOnTime4:  "13:09:34",
	}
	data, err := myTime.ToJson()
	if err != nil {
		t.Error()
	}
	fmt.Println(data)

	jsonMap := make(map[string]interface{})

	json.Unmarshal(data, &jsonMap)
	var Params TimerControlParameter
	var bAllGroups bool = true
	//Group 0
	nDayHour, err1 := GetDayHour(jsonMap["TimeOffDay0"].([]interface{}), jsonMap["TimeOffTime0"].(string))
	dMinute, err2 := GetMinute(jsonMap["TimeOffTime0"].(string))
	if err1 != nil || err2 != nil {
		Params.TimeOffDH0 = nDayHour
		Params.TimeOffMS0 = dMinute
		bAllGroups = false
	} else {
		Params.TimeOffDH0 = nDayHour
		Params.TimeOffMS0 = dMinute
	}

	nDayHour, err1 = GetDayHour(jsonMap["TimeOnDay0"].([]interface{}), jsonMap["TimeOnTime0"].(string))
	dMinute, err2 = GetMinute(jsonMap["TimeOnTime0"].(string))
	if err1 != nil || err2 != nil {
		Params.TimeOnDH0 = nDayHour
		Params.TimeOnMS0 = dMinute
		bAllGroups = false
	} else {
		Params.TimeOnDH0 = nDayHour
		Params.TimeOnMS0 = dMinute
	}

	//Group 1
	nDayHour, err1 = GetDayHour(jsonMap["TimeOffDay1"].([]interface{}), jsonMap["TimeOffTime1"].(string))
	dMinute, err2 = GetMinute(jsonMap["TimeOffTime1"].(string))
	if err1 != nil || err2 != nil {
		Params.TimeOffDH1 = nDayHour
		Params.TimeOffMS1 = dMinute
		bAllGroups = false
	} else {
		Params.TimeOffDH1 = nDayHour
		Params.TimeOffMS1 = dMinute
	}

	nDayHour, err1 = GetDayHour(jsonMap["TimeOnDay1"].([]interface{}), jsonMap["TimeOnTime1"].(string))
	dMinute, err2 = GetMinute(jsonMap["TimeOnTime1"].(string))
	if err1 != nil || err2 != nil {
		Params.TimeOnDH1 = nDayHour
		Params.TimeOnMS1 = dMinute
		bAllGroups = false
	} else {
		Params.TimeOnDH1 = nDayHour
		Params.TimeOnMS1 = dMinute
	}

	//Group 2
	nDayHour, err1 = GetDayHour(jsonMap["TimeOffDay2"].([]interface{}), jsonMap["TimeOffTime2"].(string))
	dMinute, err2 = GetMinute(jsonMap["TimeOffTime2"].(string))
	if err1 != nil || err2 != nil {
		Params.TimeOffDH2 = nDayHour
		Params.TimeOffMS2 = dMinute
		bAllGroups = false
	} else {
		Params.TimeOffDH2 = nDayHour
		Params.TimeOffMS2 = dMinute
	}

	nDayHour, err1 = GetDayHour(jsonMap["TimeOnDay2"].([]interface{}), jsonMap["TimeOnTime2"].(string))
	dMinute, err2 = GetMinute(jsonMap["TimeOnTime2"].(string))
	if err1 != nil || err2 != nil {
		Params.TimeOnDH2 = nDayHour
		Params.TimeOnMS2 = dMinute
		bAllGroups = false
	} else {
		Params.TimeOnDH2 = nDayHour
		Params.TimeOnMS2 = dMinute
	}

	//Group 3
	nDayHour, err1 = GetDayHour(jsonMap["TimeOffDay3"].([]interface{}), jsonMap["TimeOffTime3"].(string))
	dMinute, err2 = GetMinute(jsonMap["TimeOffTime3"].(string))
	if err1 != nil || err2 != nil {
		Params.TimeOffDH3 = nDayHour
		Params.TimeOffMS3 = dMinute
		bAllGroups = false
	} else {
		Params.TimeOffDH3 = nDayHour
		Params.TimeOffMS3 = dMinute
	}

	nDayHour, err1 = GetDayHour(jsonMap["TimeOnDay3"].([]interface{}), jsonMap["TimeOnTime3"].(string))
	dMinute, err2 = GetMinute(jsonMap["TimeOnTime3"].(string))
	if err1 != nil || err2 != nil {
		Params.TimeOnDH3 = nDayHour
		Params.TimeOnMS3 = dMinute
		bAllGroups = false
	} else {
		Params.TimeOnDH3 = nDayHour
		Params.TimeOnMS3 = dMinute
	}

	//Group 4
	nDayHour, err1 = GetDayHour(jsonMap["TimeOffDay4"].([]interface{}), jsonMap["TimeOffTime4"].(string))
	dMinute, err2 = GetMinute(jsonMap["TimeOffTime4"].(string))
	if err1 != nil || err2 != nil {
		Params.TimeOffDH4 = nDayHour
		Params.TimeOffMS4 = dMinute
		bAllGroups = false
	} else {
		Params.TimeOffDH4 = nDayHour
		Params.TimeOffMS4 = dMinute
	}

	nDayHour, err1 = GetDayHour(jsonMap["TimeOnDay4"].([]interface{}), jsonMap["TimeOnTime4"].(string))
	dMinute, err2 = GetMinute(jsonMap["TimeOnTime4"].(string))
	if err1 != nil || err2 != nil {
		Params.TimeOnDH4 = nDayHour
		Params.TimeOnMS4 = dMinute
		bAllGroups = false
	} else {
		Params.TimeOnDH4 = nDayHour
		Params.TimeOnMS4 = dMinute
	}

	if bAllGroups {
		Params.TimeOffDH0 |= GROUPFULL_FLAG
	}

	assert.Equal(t, "abcde", fmt.Sprintf("%v", jsonMap["SerialNumber"]))
}

func TestTimerControlJson(t *testing.T) {

	var param = TimerControlParameter{

		TimeOffDH0: 0xC015,
		TimeOffMS0: 0x3500,
		TimeOnDH0:  0xC015,
		TimeOnMS0:  0x3600,
		TimeOffDH1: 0x4015,
		TimeOffMS1: 0x3700,
		TimeOnDH1:  0x4015,
		TimeOnMS1:  0x3800,
		TimeOffDH2: 0x4015,
		TimeOffMS2: 0x3900,
		TimeOnDH2:  0x4015,
		TimeOnMS2:  0x4000,
		TimeOffDH3: 0x4015,
		TimeOffMS3: 0x4100,
		TimeOnDH3:  0x4015,
		TimeOnMS3:  0x4200,
		TimeOffDH4: 0x4015,
		TimeOffMS4: 0x4300,
		TimeOnDH4:  0x4015,
		TimeOnMS4:  0x4400,
	}

	var TimeJson TimerControlJson
	TimeJson.TimeOffDay0 = GetDay(param.TimeOffDH0)
	TimeJson.TimeOffTime0 = GetTime(param.TimeOffDH0, param.TimeOffMS0)
	TimeJson.TimeOnDay0 = GetDay(param.TimeOnDH0)
	TimeJson.TimeOnTime0 = GetTime(param.TimeOnDH0, param.TimeOnMS0)

	TimeJson.TimeOffDay1 = GetDay(param.TimeOffDH1)
	TimeJson.TimeOffTime1 = GetTime(param.TimeOffDH1, param.TimeOffMS1)
	TimeJson.TimeOnDay1 = GetDay(param.TimeOnDH1)
	TimeJson.TimeOnTime1 = GetTime(param.TimeOnDH1, param.TimeOnMS1)

	TimeJson.TimeOffDay2 = GetDay(param.TimeOffDH2)
	TimeJson.TimeOffTime2 = GetTime(param.TimeOffDH2, param.TimeOffMS2)
	TimeJson.TimeOnDay2 = GetDay(param.TimeOnDH2)
	TimeJson.TimeOnTime2 = GetTime(param.TimeOnDH2, param.TimeOnMS2)

	TimeJson.TimeOffDay3 = GetDay(param.TimeOffDH3)
	TimeJson.TimeOffTime3 = GetTime(param.TimeOffDH3, param.TimeOffMS3)
	TimeJson.TimeOnDay3 = GetDay(param.TimeOnDH3)
	TimeJson.TimeOnTime3 = GetTime(param.TimeOnDH3, param.TimeOnMS3)

	TimeJson.TimeOffDay4 = GetDay(param.TimeOffDH4)
	TimeJson.TimeOffTime4 = GetTime(param.TimeOffDH4, param.TimeOffMS4)
	TimeJson.TimeOnDay4 = GetDay(param.TimeOnDH4)
	TimeJson.TimeOnTime4 = GetTime(param.TimeOnDH4, param.TimeOnMS4)

	data, err := TimeJson.ToJson()
	if err != nil {
		t.Error()
	}

	jsonMap := make(map[string]interface{})
	json.Unmarshal(data, &jsonMap)

}

func TestRecord(t *testing.T) {

	var p = Record{
		Record:       0xa,
		ReadNo:       0xa,
		Category:     0x2,
		RecordParams: 0xca,
		YearMonth:    0x1,
		DayHour:      0x100,
		MinuteSecond: 0x57,
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
