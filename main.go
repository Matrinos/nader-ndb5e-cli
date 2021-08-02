package main

import (
	"encoding/json"
	"log"
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
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "operation",
				Aliases: []string{"op"},
				Usage:   "Operation parameter",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "runstatus",
				Aliases: []string{"rs"},
				Usage:   "Operation parameter",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "protectparameters",
				Aliases: []string{"pp"},
				Usage:   "Protect parameters",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "summary",
				Aliases: []string{"su"},
				Usage:   "Summary data",
				Action: func(c *cli.Context) error {
					return nil
				},
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
		log.Fatal(err)
	}
}

func turnOff(c *cli.Context) error {
	client, err := openConnection(c)
	defer client.CloseConnection()
	if err != nil {
		log.Fatal(err)
		return err
	}
	return SwitchBreaker(client, false)
}

func turnOn(c *cli.Context) error {
	client, err := openConnection(c)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return SwitchBreaker(client, true)
}

func readData(c *cli.Context) error {
	client, err := openConnection(c)
	defer client.CloseConnection()
	if err != nil {
		log.Fatal(err)
		return err
	}

	data, err := ReadData(client)
	if err != nil {
		log.Fatal(err)
		return err
	}

	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Println(string(jsonData))

	return nil
}

func openConnection(c *cli.Context) (*ModbusClient, error) {
	slaveID := c.Int("slave")
	address := c.String("address")

	return ConnectSlave(address, uint8(slaveID))

}
