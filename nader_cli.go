package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

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
				Name:    "record",
				Aliases: []string{"re"},
				Usage:   "Record information",
				Action:  readRecord,
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
				Usage:   "Full logs",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		Logger.Fatal(err)
	}
}

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

func readRecord(c *cli.Context) error {
	client, err := openConnection(c)
	defer client.CloseConnection()
	if err != nil {
		Logger.Fatal(err)
		return err
	}

	data, err := ReadRecord(client)
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

func outputData(data JsonMarshal) error {
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

	Logger.Println(fmt.Sprintf("Opening connection to address:%s,slave:%d", address, slaveID))

	return ConnectSlave(address, uint8(slaveID))

}
