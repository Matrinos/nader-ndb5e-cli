package main

import (
	"bytes"
	"encoding/binary"
)

// ConnectSlave try to connect to given slave id
// results, err := client.ReadDiscreteInputs(15, 2)
// 9600, 8, "N", 1
// address : "/dev/ttyUSB0"
func ConnectSlave(address string, slaveID uint8) (*ModbusClient, error) {
	// Modbus RTU/ASCII
	client, err := NewDeviceClient(&ConnectionInfo{
		BaudRate:    9600,
		DataBits:    8,
		Parity:      "N",
		StopBits:    1,
		UnitID:      slaveID,
		Timeout:     5,
		IdleTimeout: 3600, // TODO: right idle timeout?
		Address:     address,
	})

	if err != nil {
		return client, err
	}

	err = client.OpenConnection()

	if err != nil {
		return client, err
	}

	Logger.Println("Device connected")
	return client, nil
}

func ReadProduct(client *ModbusClient) (*Product, error) {
	p := Product{}

	results, err := client.ReadHoldingRegisters(PRODUCT_ADDR, PRODUCT_LEN)
	if err != nil {
		return &p, err
	}

	err = binary.Read(bytes.NewReader(results), binary.BigEndian, &p)
	if err != nil {
		return &p, err
	}

	return &p, nil
}

func ReadOpParameters(client *ModbusClient) (*OpParameters, error) {
	o := OpParameters{}

	results, err := client.ReadHoldingRegisters(OPPARAMETERS_ADDR, OPPARAMETERS_LEN)
	if err != nil {
		return &o, err
	}

	err = binary.Read(bytes.NewReader(results), binary.BigEndian, &o)
	if err != nil {
		return &o, err
	}

	return &o, nil
}

// ReadRunStatus read run status of the device
func ReadRunStatus(client *ModbusClient) (*RunStatus, error) {
	r := RunStatus{}

	results, err := client.ReadHoldingRegisters(RUNSTATUS_ADDR, RUNSTATUS_LEN)
	if err != nil {
		return &r, err
	}

	err = binary.Read(bytes.NewReader(results), binary.BigEndian, &r)
	if err != nil {
		return &r, err
	}

	return &r, nil
}

func ReadData(client *ModbusClient) (*MetricalData, error) {
	d := MetricalData{}

	results, err := client.ReadHoldingRegisters(METRICALDATA_ADDR, METRICALDATA_LEN)
	if err != nil {
		return &d, err
	}

	err = binary.Read(bytes.NewReader(results), binary.BigEndian, &d)
	if err != nil {
		return &d, err
	}

	return &d, nil
}

func ReadLogs(client *ModbusClient, addr uint16) (*RecordLogs, error) {
	l := RecordLogs{}

	results, err := client.ReadHoldingRegisters(addr, RECORD_LOG_LEN)
	if err != nil {
		return &l, err
	}

	err = binary.Read(bytes.NewReader(results), binary.BigEndian, &l)
	if err != nil {
		return &l, err
	}

	return &l, nil
}

func ReadProtectParameters(client *ModbusClient) (*ProtectParameters, error) {
	p := ProtectParameters{}

	results, err := client.ReadHoldingRegisters(PRETECTPARAMETERS_ADDR, PRETECTPARAMETERS_LEN)
	if err != nil {
		return &p, err
	}

	err = binary.Read(bytes.NewReader(results), binary.BigEndian, &p)
	if err != nil {
		return &p, err
	}

	return &p, nil
}

func ReadRecord(client *ModbusClient, record uint16) (*Record, error) {
	r := Record{}

	results, err := client.ReadHoldingRegisters(record, RECORD_INFO_LEN)
	if err != nil {
		return &r, err
	}

	err = binary.Read(bytes.NewReader(results), binary.BigEndian, &r)
	if err != nil {
		return &r, err
	}

	return &r, nil
}

func ReadSummary1(client *ModbusClient) (*Summary1, error) {
	s := Summary1{}

	results, err := client.ReadHoldingRegisters(FE_TEMPERATURES_ADDR, FE_TEMPERATURES_LEN)
	if err != nil {
		return &s, err
	}

	err = binary.Read(bytes.NewReader(results), binary.BigEndian, &s)
	if err != nil {
		return &s, err
	}

	return &s, nil
}

func ReadSummary2(client *ModbusClient) (*Summary2, error) {
	s := Summary2{}

	results, err := client.ReadHoldingRegisters(FE_ENERGYPERHOUR_ADDR, FE_ENERGYPERHOUR_LEN)
	if err != nil {
		return &s, err
	}

	err = binary.Read(bytes.NewReader(results), binary.BigEndian, &s)
	if err != nil {
		return &s, err
	}

	return &s, nil
}

func ReadSummary3(client *ModbusClient) (*Summary3, error) {
	s := Summary3{}

	results, err := client.ReadHoldingRegisters(FE_ENERGYPERDAY_ADDR, FE_ENERGYPERDAY_LEN)
	if err != nil {
		return &s, err
	}

	err = binary.Read(bytes.NewReader(results), binary.BigEndian, &s)
	if err != nil {
		return &s, err
	}

	return &s, nil
}

func ReadSummary4(client *ModbusClient) (*Summary4, error) {
	s := Summary4{}

	results, err := client.ReadHoldingRegisters(FE_ENERGYPERMONTH_ADDR, FE_ENERGYPERMONTH_LEN)
	if err != nil {
		return &s, err
	}

	err = binary.Read(bytes.NewReader(results), binary.BigEndian, &s)
	if err != nil {
		return &s, err
	}

	return &s, nil
}

func ReadTimerParameters(client *ModbusClient) (*TimerControlParameter, error) {
	s := TimerControlParameter{}

	results, err := client.ReadHoldingRegisters(REMOTECONTROL_ADDR, REMOTECONTROL_LEN)
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

func SetTimerParameters(client *ModbusClient, jsonpath string) error {
	r := TimerControlParameter{}

	err := GetRemoteCtlSetting(jsonpath, &r)
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
	return client.WriteMultipleRegisters(REMOTECONTROL_ADDR, REMOTECONTROL_LEN, buf.Bytes())
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
