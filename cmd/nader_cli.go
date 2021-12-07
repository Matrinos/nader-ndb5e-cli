package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Matrinos/nader-ndb5e-cli/models"
	"github.com/urfave/cli/v2"
)

var Logger = log.New(os.Stdout, "ascii: ", log.LstdFlags)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "slave",
				Aliases: []string{"s"},
				Value:   100,
				Usage:   "Slave ID",
			},
			&cli.StringFlag{
				Name:     "address",
				Aliases:  []string{"a"},
				Value:    "",
				Usage:    "/dev/usb....",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "protocol",
				Aliases:  []string{"p"},
				Value:    "modbus-tcp",
				Usage:    "modbus-tcp/modbus-modbus",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "port",
				Aliases:  []string{"r"},
				Value:    "502",
				Usage:    "502",
				Required: false,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "turnon",
				Aliases: []string{"to"},
				Usage:   "Turn on",
				Action:  turnOn,
			},
			{
				Name:    "turnoff",
				Aliases: []string{"tf"},
				Usage:   "Turn off",
				Action:  turnOff,
			},
			{
				Name:    "powerdata",
				Aliases: []string{"pd"},
				Usage:   "Read power data",
				Action:  readData,
			},
			{
				Name:    "product",
				Aliases: []string{"po"},
				Usage:   "Read product information",
				Action:  readProduct,
			},
			{
				Name:    "operation",
				Aliases: []string{"op"},
				Usage:   "Operation parameter",
				Action:  readOpParameters,
			},
			{
				Name:    "operation",
				Aliases: []string{"sp"},
				Usage:   "Operation parameter",
				Action:  setOpParameters,
			},
			{
				Name:    "runstatus",
				Aliases: []string{"rs"},
				Usage:   "Operation parameter",
				Action:  readRunStatus,
			},
			{
				Name:    "protectparameters",
				Aliases: []string{"pp"},
				Usage:   "Protect parameters",
				Action:  readProtectParameters,
			},
			{
				Name:    "setrecordnum",
				Aliases: []string{"sr"},
				Usage:   "setRecordNumber --recordtype=type, Record type --recordnum Record No.",
				Action:  setRecordNumber,
				Flags: []cli.Flag{
					&cli.IntFlag{Name: "recordtype", Usage: "--recordtype"},
					&cli.IntFlag{Name: "recordnum", Usage: "--recordnum"},
				},
			},
			{
				Name:    "readrecord",
				Aliases: []string{"rr"},
				Usage:   "record data",
				Action:  readRecord,
				Flags: []cli.Flag{
					&cli.IntFlag{Name: "recordtype", Usage: "--recordtype"},
				},
			},
			{
				Name:    "timerparameters",
				Aliases: []string{"tp"},
				Usage:   "Timer parameters",
				Action:  readTimerParameters,
			},
			{
				Name:    "settimerparameters",
				Aliases: []string{"st"},
				Usage:   "Json file path --jsonpath=file",
				Action:  setTimerParameters,
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "jsonpath", Usage: "--jsonpath"},
				},
			},
			{
				Name:    "telecommand",
				Aliases: []string{"tc"},
				Usage:   "telecommand",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "summary1",
				Aliases: []string{"su1"},
				Usage:   "Summary data",
				Action:  readSummary1,
			},
			{
				Name:    "summary2",
				Aliases: []string{"su2"},
				Usage:   "Summary data",
				Action:  readSummary2,
			},
			{
				Name:    "summary3",
				Aliases: []string{"su3"},
				Usage:   "Summary data",
				Action:  readSummary3,
			},
			{
				Name:    "summary4",
				Aliases: []string{"su4"},
				Usage:   "Summary data",
				Action:  readSummary4,
			},
			{
				Name:    "logs",
				Aliases: []string{"lg"},
				Usage:   "readLogs --logtype=type --logindex=index",
				Action:  readLogs,
				Flags: []cli.Flag{
					&cli.IntFlag{Name: "logtype", Usage: "--logtype"},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		Logger.Fatal(err)
	}
}

/*
func main() {

	set := flag.NewFlagSet("flag", 0)
	set.String("address", "com3", "test")
	set.Int("slave", 100, "test")

	ctx := cli.NewContext(nil, set, nil)

	//readProduct(ctx)
	//readOpParameters(ctx)
	//readRunStatus(ctx)
	//readData(ctx)
	//readProtectParameters(ctx)
	//readSummary1(ctx)
	//readSummary2(ctx)
	//readSummary3(ctx)
	//readSummary4(ctx)
	setTimerToSwitch(ctx)
	readRemoteCmd(ctx)
}
*/
func turnOff(c *cli.Context) error {
	client, err := openConnection(c)
	defer client.CloseConnection()
	if err != nil {
		Logger.Fatal(err)
		return err
	}
	return SwitchBreaker(client, false)
}

func turnOn(c *cli.Context) error {
	Logger.Println("Turning on")
	client, err := openConnection(c)
	defer client.CloseConnection()
	if err != nil {
		Logger.Fatal(err)
		return err
	}
	return SwitchBreaker(client, true)
}

func setTimerParameters(c *cli.Context) error {
	Logger.Println("set Timers to control switch")
	client, err := openConnection(c)
	defer client.CloseConnection()
	if err != nil {
		Logger.Fatal(err)
		return err
	}

	jsonpath := c.String("jsonpath")

	return SetTimerParameters(client, jsonpath)
}

func readData(c *cli.Context) error {
	client, err := openConnection(c)
	defer client.CloseConnection()
	if err != nil {
		Logger.Fatal(err)
		return err
	}

	data, err := ReadData(client)
	if err != nil {
		Logger.Fatal(err)
		return err
	}

	return outputData(data)
}

func readLogs(c *cli.Context) error {
	err := error(nil)
	var GroupNum uint16 = 0
	logType := uint16(c.Int("logtype"))

	if logType == models.FAULT_TYPE {
		GroupNum = (models.MAX_FAULTRECORDLOG_NUM / models.LOGSINGROUP_NUM)
	} else if logType == models.ALARM_TYPE {
		GroupNum = (models.MAX_ALARMRECORDLOG_NUM / models.LOGSINGROUP_NUM)
	} else if logType == models.SWITCH_TYPE {
		GroupNum = (models.MAX_SWITCHRECORDLOG_NUM / models.LOGSINGROUP_NUM)
	}

	for GroupID := uint16(0); GroupID < GroupNum; GroupID++ {
		readLogGroup(c, GroupID)
	}

	return err
}

func readLogGroup(c *cli.Context, GroupIndex uint16) error {
	client, err := openConnection(c)
	defer client.CloseConnection()
	if err != nil {
		Logger.Fatal(err)
		return err
	}

	var addr uint16 = 0
	logType := uint16(c.Int("logtype"))

	for index := uint16(0); index < models.LOGSINGROUP_NUM; index++ {
		logIndex := GroupIndex*models.LOGSINGROUP_NUM + index
		if logType == models.FAULT_TYPE {
			addr = (models.FAULTRECORDLOG_ADDR + logIndex*models.RECORD_LOG_LEN)
		} else if logType == models.ALARM_TYPE {
			addr = (models.ALARMRECORDLOG_ADDR + logIndex)
		} else if logType == models.SWITCH_TYPE {
			addr = (models.SWITCHRECORDLOG_ADDR + logIndex)
		} else {
			return err
		}
		data, err := ReadLogs(client, addr)
		if err != nil {
			Logger.Fatal(err)
			return err
		}

		data.LogNo = logIndex
		data.LogType = logType

		err = outputData(data)
		time.Sleep(2000)
	}

	return err
}

func readProtectParameters(c *cli.Context) error {
	client, err := openConnection(c)
	defer client.CloseConnection()
	if err != nil {
		Logger.Fatal(err)
		return err
	}

	data, err := ReadProtectParameters(client)
	if err != nil {
		Logger.Fatal(err)
		return err
	}

	return outputData(data)
}

func setRecordNumber(c *cli.Context) error {
	client, err := openConnection(c)
	defer client.CloseConnection()
	if err != nil {
		Logger.Fatal(err)
		return err
	}

	recordtype := c.Int("recordtype")
	recordnum := uint16(c.Int("recordnum"))
	if recordtype == models.FAULT_TYPE {
		return SetRecordNo(client, models.FAULTRECORD_NUM_ADDR, recordnum)
	} else if recordtype == models.ALARM_TYPE {
		return SetRecordNo(client, models.ALARMRECORD_NUM_ADDR, recordnum)
	} else if recordtype == models.SWITCH_TYPE {
		return SetRecordNo(client, models.SWITCHRECORD_NUM_ADDR, recordnum)
	}

	return err
}

func readRecord(c *cli.Context) error {
	client, err := openConnection(c)
	defer client.CloseConnection()
	if err != nil {
		Logger.Fatal(err)
		return err
	}

	recordtype := c.Int("recordtype")
	var addr uint16 = 0

	if recordtype == models.FAULT_TYPE {
		addr = models.FAULTRECORD_ADDR
	} else if recordtype == models.ALARM_TYPE {
		addr = models.ALARMRECORD_ADDR
	} else if recordtype == models.SWITCH_TYPE {
		addr = models.SWITCHRECORD_ADDR
	} else {
		return err
	}

	data, err := ReadRecord(client, addr)
	if err != nil {
		Logger.Fatal(err)
		return err
	}

	return outputData(data)
}

func readSummary1(c *cli.Context) error {
	client, err := openConnection(c)
	defer client.CloseConnection()
	if err != nil {
		Logger.Fatal(err)
		return err
	}

	data, err := ReadSummary1(client)
	if err != nil {
		Logger.Fatal(err)
		return err
	}

	return outputData(data)
}

func readSummary2(c *cli.Context) error {
	client, err := openConnection(c)
	defer client.CloseConnection()
	if err != nil {
		Logger.Fatal(err)
		return err
	}

	data, err := ReadSummary2(client)
	if err != nil {
		Logger.Fatal(err)
		return err
	}

	return outputData(data)
}

func readSummary3(c *cli.Context) error {
	client, err := openConnection(c)
	defer client.CloseConnection()
	if err != nil {
		Logger.Fatal(err)
		return err
	}

	data, err := ReadSummary3(client)
	if err != nil {
		Logger.Fatal(err)
		return err
	}

	return outputData(data)
}

func readSummary4(c *cli.Context) error {
	client, err := openConnection(c)
	defer client.CloseConnection()
	if err != nil {
		Logger.Fatal(err)
		return err
	}

	data, err := ReadSummary4(client)
	if err != nil {
		Logger.Fatal(err)
		return err
	}

	return outputData(data)
}

func readProduct(c *cli.Context) error {
	client, err := openConnection(c)
	defer client.CloseConnection()
	if err != nil {
		Logger.Fatal(err)
		return err
	}

	data, err := ReadProduct(client)
	if err != nil {
		Logger.Fatal(err)
		return err
	}

	return outputData(data)
}

func readOpParameters(c *cli.Context) error {
	client, err := openConnection(c)
	defer client.CloseConnection()
	if err != nil {
		Logger.Fatal(err)
		return err
	}

	data, err := ReadOpParameters(client)
	if err != nil {
		Logger.Fatal(err)
		return err
	}

	return outputData(data)
}

func setOpParameters(c *cli.Context) error {
	client, err := openConnection(c)
	defer client.CloseConnection()
	if err != nil {
		Logger.Fatal(err)
		return err
	}

	err = SetOpParameters(client)
	if err != nil {
		Logger.Fatal(err)
		return err
	}

	return err
}
func readRunStatus(c *cli.Context) error {
	client, err := openConnection(c)
	defer client.CloseConnection()
	if err != nil {
		Logger.Fatal(err)
		return err
	}

	data, err := ReadRunStatus(client)
	if err != nil {
		Logger.Fatal(err)
		return err
	}

	return outputData(data)
}

func readTimerParameters(c *cli.Context) error {
	client, err := openConnection(c)
	defer client.CloseConnection()
	if err != nil {
		Logger.Fatal(err)
		return err
	}

	data, err := ReadTimerParameters(client)
	if err != nil {
		Logger.Fatal(err)
		return err
	}

	return outputData(data)
}

func outputData(data models.JsonMarshal) error {
	jsonData, err := data.ToJson()
	if err != nil {
		Logger.Fatal(err)
		return err
	}

	Logger.Println(string(jsonData))
	return nil
}

func openConnection(c *cli.Context) (*ModbusClient, error) {

	slaveID := c.Int("slave")
	address := c.String("address")
	protocol := c.String("protocol")
	port := c.Int("port")

	Logger.Println(fmt.Sprintf("Opening connection to address:%s,slave:%d, port:%d", address, slaveID, port))

	return ConnectSlave(address, uint8(slaveID), protocol, uint8(port))

}
