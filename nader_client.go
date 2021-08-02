package main

import (
	"bytes"
	"encoding/binary"
)

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

func ReadData(client *ModbusClient) (*Data, error) {
	d := Data{}

	results, err := client.ReadHoldingRegisters(0x260, 80)
	if err != nil {
		return &d, err
	}

	err = binary.Read(bytes.NewReader(results), binary.BigEndian, &d)
	if err != nil {
		return &d, err
	}

	return &d, nil
}

func SwitchBreaker(client *ModbusClient, is_on bool) error {
	if is_on {
		client.WriteSingleRegister(0x0400, 0xff)
	}

	return client.WriteSingleRegister(0x0400, 0xff00)
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
