package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	initConfig()

	client, err := initConnection()

	if err != nil {
		log.Fatal(err)
		return
	}

	// 	var cmdPrint = &cobra.Command{
	// 		Use:   "print [string to print]",
	// 		Short: "Print anything to the screen",
	// 		Long: `print is for printing anything back to the screen.
	// For many years people have printed back to the screen.`,
	// 		Args: cobra.MinimumNArgs(1),
	// 		Run: func(cmd *cobra.Command, args []string) {
	// 			fmt.Println("Print: " + strings.Join(args, " "))
	// 		},
	// 	}

	var cmdMonitorPowerUsage = &cobra.Command{
		Use:   "echo [string to echo]",
		Short: "Echo anything to the screen",
		Long: `echo is for echoing anything back.
	Echo works a lot like print, except it has a child command.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ReadData(client)
		},
	}

	// 	var cmdTimes = &cobra.Command{
	// 		Use:   "times [string to echo]",
	// 		Short: "Echo anything to the screen more times",
	// 		Long: `echo things multiple times back to the user by providing
	// a count and a string.`,
	// 		Args: cobra.MinimumNArgs(1),
	// 		Run: func(cmd *cobra.Command, args []string) {
	// 			for i := 0; i < echoTimes; i++ {
	// 				fmt.Println("Echo: " + strings.Join(args, " "))
	// 			}
	// 		},
	// 	}

	cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdPrint, cmdMonitorPowerUsage)
	cmdMonitorPowerUsage.AddCommand(cmdTimes)
	rootCmd.Execute()
}

func initConfig() {
	viper.SetConfigName("config")       // name of config file (without extension)
	viper.SetConfigType("yaml")         // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/nader/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.nader") // call multiple times to add many search paths
	viper.AddConfigPath(".")
	viper.SetDefault(Protocol, ProtocolTCP)
	viper.SetDefault(UnitID, 100)
	viper.SetDefault(Address, "/dev/ttyUSB0")
	// viper.SetDefault(BaudRate, 9600)
	// viper.SetDefault(DataBits, 8)
	// viper.SetDefault(StopBits, 1)
	// viper.SetDefault(Parity, "N")
	// viper.SetDefault(Timeout, 5)
	// viper.SetDefault(IdleTimeout, 60)

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

}

func initConnection() (*ModbusClient, error) {
	protocol := viper.Get(Protocol)
	if protocol == ProtocolRTU {
		address := viper.GetString(Address)
		unitID := uint8(viper.GetUint(UnitID))

		return ConnectSlave(address, unitID)
	}

	//TODO: handle unknow case
	return nil, nil
}
