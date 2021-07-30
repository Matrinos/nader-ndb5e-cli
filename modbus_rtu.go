package main

import (
	"time"

	"github.com/goburrow/modbus"
)

// results, err := client.ReadDiscreteInputs(15, 2)
// 9600, 8, "N", 1
// port : "/dev/ttyUSB0"
func ConnectSlave(port string, slaveId byte) (modbus.Client, error) {
	// Modbus RTU/ASCII
	handler := modbus.NewRTUClientHandler(port)
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "N"
	handler.StopBits = 1
	handler.SlaveId = slaveId
	handler.Timeout = 5 * time.Second

	err := handler.Connect()
	defer handler.Close()
	if err != nil {
		return nil, err
	}

	client := modbus.NewClient(handler)

	return client, nil
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
