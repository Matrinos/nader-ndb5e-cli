package modbus

import (
	"bytes"
	"encoding/binary"
	"strconv"
	"time"

	"github.com/Matrinos/nader-ndb5e-cli/models"
)

// ConnectSlave try to connect to given slave id
// results, err := client.ReadDiscreteInputs(15, 2)
// 9600, 8, "N", 1
// address : "/dev/ttyUSB0"
func ConnectSlave(address string, slaveID uint8, protocol string, port uint8) (*ModbusClient, error) {
	// Modbus RTU/ASCII
	client, err := NewDeviceClient(&ConnectionInfo{
		Protocol:    protocol,
		Port:        502,
		BaudRate:    9600,
		DataBits:    8,
		Parity:      "N",
		StopBits:    1,
		UnitID:      slaveID,
		Timeout:     5,
		IdleTimeout: 3600, // TODO: right idle timeout?
		Address:     address,
	}, nil)

	if err != nil {
		return client, err
	}

	err = client.OpenConnection()

	if err != nil {
		return client, err
	}

	client.Logger.Info("Device connected")
	return client, nil
}

func ReadProduct(client *ModbusClient) (*models.Product, error) {
	p := models.Product{}

	results, err := client.ReadHoldingRegisters(models.PRODUCT_ADDR, models.PRODUCT_LEN)
	if err != nil {
		return &p, err
	}

	err = binary.Read(bytes.NewReader(results), binary.BigEndian, &p)
	if err != nil {
		return &p, err
	}

	return &p, nil
}

func ReadOpParameters(client *ModbusClient) (*models.OpParameters, error) {
	o := models.OpParameters{}

	results, err := client.ReadHoldingRegisters(models.OPPARAMETERS_ADDR, models.OPPARAMETERS_LEN)
	if err != nil {
		return &o, err
	}

	err = binary.Read(bytes.NewReader(results), binary.BigEndian, &o)
	if err != nil {
		return &o, err
	}

	return &o, nil
}

func SetOpParameters(client *ModbusClient) error {
	op := models.OpParameters{}
	currentTime := time.Now()

	yearMonthString := currentTime.Format("0601")
	yearMonth, err := strconv.ParseInt(yearMonthString, 16, 64)
	if err != nil {
		return err
	}
	op.YearMonth = uint16(yearMonth)

	dayHourString := currentTime.Format("0215")
	dayHour, err := strconv.ParseInt(dayHourString, 16, 64)
	if err != nil {
		return err
	}
	op.DayHour = uint16(dayHour)

	minuteSecondString := currentTime.Format("0405")
	minuteSecond, err := strconv.ParseInt(minuteSecondString, 16, 64)
	if err != nil {
		return err
	}
	op.MinuteSecond = uint16(minuteSecond)

	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, &op)
	if err != nil {
		return err
	}

	return client.WriteMultipleRegisters(models.OPPARAMETERS_ADDR, models.OPPARAMETERS_LEN, buf.Bytes())
}

// ReadRunStatus read run status of the device
func ReadRunStatus(client *ModbusClient) (*models.RunStatus, error) {
	r := models.RunStatus{}

	results, err := client.ReadHoldingRegisters(models.RUNSTATUS_ADDR, models.RUNSTATUS_LEN)
	if err != nil {
		return &r, err
	}

	err = binary.Read(bytes.NewReader(results), binary.BigEndian, &r)
	if err != nil {
		return &r, err
	}

	return &r, nil
}

func ReadData(client *ModbusClient) (*models.MetricalData, error) {
	d := models.MetricalData{}

	results, err := client.ReadHoldingRegisters(models.METRICALDATA_ADDR, models.METRICALDATA_LEN)
	if err != nil {
		return &d, err
	}

	err = binary.Read(bytes.NewReader(results), binary.BigEndian, &d)
	if err != nil {
		return &d, err
	}

	return &d, nil
}

func ReadLogs(client *ModbusClient, addr uint16) (*models.RecordLogsInfo, error) {
	l := models.RecordLogsInfo{}

	results, err := client.ReadHoldingRegisters(addr, models.RECORD_LOG_LEN)
	if err != nil {
		return &l, err
	}

	err = binary.Read(bytes.NewReader(results), binary.BigEndian, &l.Logs)
	if err != nil {
		return &l, err
	}

	return &l, nil
}

func ReadProtectParameters(client *ModbusClient) (*models.ProtectParameters, error) {
	p := models.ProtectParameters{}

	results, err := client.ReadHoldingRegisters(models.PRETECTPARAMETERS_ADDR, models.PRETECTPARAMETERS_LEN)
	if err != nil {
		return &p, err
	}

	err = binary.Read(bytes.NewReader(results), binary.BigEndian, &p)
	if err != nil {
		return &p, err
	}

	return &p, nil
}

func ReadRecord(client *ModbusClient, addr uint16) (*models.Record, error) {
	r := models.Record{}

	results, err := client.ReadHoldingRegisters(addr, models.RECORD_INFO_LEN)
	if err != nil {
		return &r, err
	}

	err = binary.Read(bytes.NewReader(results), binary.BigEndian, &r)
	if err != nil {
		return &r, err
	}

	return &r, nil
}

func ReadSummary1(client *ModbusClient) (*models.Summary1, error) {
	s := models.Summary1{}

	results, err := client.ReadHoldingRegisters(models.FE_TEMPERATURES_ADDR, models.FE_TEMPERATURES_LEN)
	if err != nil {
		return &s, err
	}

	err = binary.Read(bytes.NewReader(results), binary.BigEndian, &s)
	if err != nil {
		return &s, err
	}

	return &s, nil
}

func ReadSummary2(client *ModbusClient) (*models.Summary2, error) {
	s := models.Summary2{}

	results, err := client.ReadHoldingRegisters(models.FE_ENERGYPERHOUR_ADDR, models.FE_ENERGYPERHOUR_LEN)
	if err != nil {
		return &s, err
	}

	err = binary.Read(bytes.NewReader(results), binary.BigEndian, &s)
	if err != nil {
		return &s, err
	}

	return &s, nil
}

func ReadSummary3(client *ModbusClient) (*models.Summary3, error) {
	s := models.Summary3{}

	results, err := client.ReadHoldingRegisters(models.FE_ENERGYPERDAY_ADDR, models.FE_ENERGYPERDAY_LEN)
	if err != nil {
		return &s, err
	}

	err = binary.Read(bytes.NewReader(results), binary.BigEndian, &s)
	if err != nil {
		return &s, err
	}

	return &s, nil
}

func ReadSummary4(client *ModbusClient) (*models.Summary4, error) {
	s := models.Summary4{}

	results, err := client.ReadHoldingRegisters(models.FE_ENERGYPERMONTH_ADDR, models.FE_ENERGYPERMONTH_LEN)
	if err != nil {
		return &s, err
	}

	err = binary.Read(bytes.NewReader(results), binary.BigEndian, &s)
	if err != nil {
		return &s, err
	}

	return &s, nil
}

func ReadTimerParameters(client *ModbusClient) (*models.TimerControlParameter, error) {
	s := models.TimerControlParameter{}

	results, err := client.ReadHoldingRegisters(models.REMOTECONTROL_ADDR, models.REMOTECONTROL_LEN)
	if err != nil {
		return &s, err
	}

	err = binary.Read(bytes.NewReader(results), binary.BigEndian, &s)
	if err != nil {
		return &s, err
	}

	return &s, nil
}

func SwitchBreaker(client *ModbusClient, is_on bool) error {
	if is_on {
		return client.WriteSingleRegister(0x0400, 0xff)
	}

	return client.WriteSingleRegister(0x0400, 0xff00)
}

func SetRecordNo(client *ModbusClient, addr uint16, num uint16) error {

	return client.WriteSingleRegister(addr, num)
}

func SetTimerParameters(client *ModbusClient, jsonpath string) error {
	r := models.TimerControlParameter{}

	err := models.GetRemoteCtlSetting(jsonpath, &r)
	if err != nil {
		return err
	}
	//for test
	/*
		var w uint16 = TIMER_SUNDAY << 8 //week sunday
		var h uint16 = 0x15              //hour    15
		var t uint16 = 0x3500            //minute   20:00

		r.TimeOffDH0 = w | h
		r.TimeOffMS0 = t + 0x0000

		r.TimeOnDH0 = w | h
		r.TimeOnMS0 = t + 0x0100

		r.TimeOffDH1 = w | h
		r.TimeOffMS1 = t + 0x0200

		r.TimeOnDH1 = w | h
		r.TimeOnMS1 = t + 0x0300

		r.TimeOffDH2 = w | h
		r.TimeOffMS2 = t + 0x0400

		r.TimeOnDH2 = w | h
		r.TimeOnMS2 = t + 0x0500

		r.TimeOffDH3 = w | h
		r.TimeOffMS3 = 0x0600

		r.TimeOnDH3 = w | h
		r.TimeOnMS3 = t + 0x0700

		r.TimeOnDH4 = w | h
		r.TimeOnMS4 = t + 0x0800

		r.TimeOffDH4 = w | h
		r.TimeOffMS4 = t + 0x0900
	*/
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, &r)
	if err != nil {
		return err
	}
	return client.WriteMultipleRegisters(models.REMOTECONTROL_ADDR, models.REMOTECONTROL_LEN, buf.Bytes())
}

// func sample() {
// 	client, err := ConnectSlave("/dev/tty.usbserial-AG0JG5OU", 100)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	results, err := client.ReadHoldingRegisters(0x0210, 10)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// data := binary.BigEndian.Uint16(results)
// 	// fmt.Println(data)
// 	fmt.Print(string(results))
// }
