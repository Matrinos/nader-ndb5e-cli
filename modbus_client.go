// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0
// Modified by Q.s. Wang for our project usage

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	MODBUS "github.com/goburrow/modbus"
)

var Logger = log.New(os.Stdout, "ascii: ", log.LstdFlags)

const (
	ProtocolTCP = "modbus-tcp"
	ProtocolRTU = "modbus-rtu"

	Address  = "Address"
	Port     = "Port"
	UnitID   = "UnitID"
	BaudRate = "BaudRate"
	DataBits = "DataBits"
	StopBits = "StopBits"
	// Parity: N - None, O - Odd, E - Even
	Parity = "Parity"

	Timeout     = "Timeout"
	IdleTimeout = "IdleTimeout"
)

// ConnectionInfo is device connection info
type ConnectionInfo struct {
	Protocol string
	Address  string
	Port     int
	BaudRate int
	DataBits int
	StopBits int
	Parity   string
	UnitID   uint8
	// Connect & Read timeout(seconds)
	Timeout int
	// Idle timeout(seconds) to close the connection
	IdleTimeout int
}

// ModbusClient is used for connecting the device and read/write value
type ModbusClient struct {
	// IsModbusTcp is a value indicating the connection type
	IsModbusTcp bool
	// TCPClientHandler is ued for holding device TCP connection
	TCPClientHandler MODBUS.TCPClientHandler
	// TCPClientHandler is ued for holding device RTU connection
	RTUClientHandler MODBUS.RTUClientHandler

	client MODBUS.Client
}

func (c *ModbusClient) OpenConnection() error {
	var err error
	var newClient MODBUS.Client
	if c.IsModbusTcp {
		err = c.TCPClientHandler.Connect()
		newClient = MODBUS.NewClient(&c.TCPClientHandler)
		Logger.Println("Modbus client create TCP connection.")
	} else {
		err = c.RTUClientHandler.Connect()
		newClient = MODBUS.NewClient(&c.RTUClientHandler)
		Logger.Println("Modbus client create RTU connection.")
	}
	c.client = newClient
	return err
}

func (c *ModbusClient) CloseConnection() error {
	var err error
	if c.IsModbusTcp {
		err = c.TCPClientHandler.Close()

	} else {
		err = c.RTUClientHandler.Close()
	}
	return err
}

func (c *ModbusClient) ReadHoldingRegisters(startingAddress uint16, length uint16) ([]byte, error) {
	response, err := c.client.ReadHoldingRegisters(startingAddress, length)
	if err != nil {
		return response, err
	}

	Logger.Println(fmt.Sprintf("Modbus client ReadHoldingRegisters's results %v", response))

	return response, nil
}

func (c *ModbusClient) ReadDiscreteInputs(startingAddress uint16, length uint16) ([]byte, error) {
	response, err := c.client.ReadDiscreteInputs(startingAddress, length)
	if err != nil {
		return response, err
	}

	Logger.Println(fmt.Sprintf("Modbus client ReadDiscreteInputs's results %v", response))

	return response, nil
}

func (c *ModbusClient) ReadInputRegisters(startingAddress uint16, length uint16) ([]byte, error) {
	response, err := c.client.ReadInputRegisters(startingAddress, length)
	if err != nil {
		return response, err
	}

	Logger.Println(fmt.Sprintf("Modbus client ReadInputRegisters's results %v", response))

	return response, nil
}

func (c *ModbusClient) ReadCoils(startingAddress uint16, length uint16) ([]byte, error) {
	response, err := c.client.ReadCoils(startingAddress, length)
	if err != nil {
		return response, err
	}

	Logger.Println(fmt.Sprintf("Modbus client ReadCoils's results %v", response))

	return response, nil
}

func (c *ModbusClient) WriteMultipleCoils(startingAddress uint16, length uint16, value []byte) error {
	result, err := c.client.WriteMultipleCoils(startingAddress, length, value)

	if err != nil {
		return err
	}
	Logger.Println(fmt.Sprintf("Modbus client SetValue successful, results: %v", result))

	return nil
}

func (c *ModbusClient) WriteSingleRegister(startingAddress uint16, value uint16) error {
	result, err := c.client.WriteSingleRegister(startingAddress, value)

	if err != nil {
		return err
	}
	Logger.Println(fmt.Sprintf("Modbus client SetValue successful, results: %v", result))

	return nil
}

func (c *ModbusClient) WriteMultipleRegisters(startingAddress uint16, length uint16, value []byte) error {
	result, err := c.client.WriteMultipleRegisters(startingAddress, length, value)

	if err != nil {
		return err
	}
	Logger.Println(fmt.Sprintf("Modbus client SetValue successful, results: %v", result))

	return nil
}

func NewDeviceClient(connectionInfo *ConnectionInfo) (*ModbusClient, error) {
	client := new(ModbusClient)
	var err error
	if connectionInfo.Protocol == ProtocolTCP {
		client.IsModbusTcp = true
	}
	if client.IsModbusTcp {
		client.TCPClientHandler.Address = fmt.Sprintf("%s:%d", connectionInfo.Address, connectionInfo.Port)
		client.TCPClientHandler.SlaveId = byte(connectionInfo.UnitID)
		client.TCPClientHandler.Timeout = time.Duration(connectionInfo.Timeout) * time.Second
		client.TCPClientHandler.IdleTimeout = time.Duration(connectionInfo.IdleTimeout) * time.Second
		client.TCPClientHandler.Logger = log.New(os.Stdout, "", log.LstdFlags)
	} else {
		serialParams := strings.Split(connectionInfo.Address, ",")
		client.RTUClientHandler.Address = serialParams[0]
		client.RTUClientHandler.SlaveId = byte(connectionInfo.UnitID)
		client.RTUClientHandler.Timeout = time.Duration(connectionInfo.Timeout) * time.Second
		client.RTUClientHandler.IdleTimeout = time.Duration(connectionInfo.IdleTimeout) * time.Second
		client.RTUClientHandler.BaudRate = connectionInfo.BaudRate
		client.RTUClientHandler.DataBits = connectionInfo.DataBits
		client.RTUClientHandler.StopBits = connectionInfo.StopBits
		client.RTUClientHandler.Parity = connectionInfo.Parity
		client.RTUClientHandler.Logger = log.New(os.Stdout, "", log.LstdFlags)
	}
	return client, err
}
