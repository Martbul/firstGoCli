package info

import (
	"fmt"
	"log"
	"os"

	"github.com/ricochet2200/go-disk-usage/du"
	"github.com/shirou/gopsutil/v4/process"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
		"github.com/fatih/color" 

)

var diskUsageCmd = &cobra.Command{
	Use:   "diskUsage",
	Short: "Prints disk usage of the current directory",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {

		//Setting default directory
		defaultDirectory := "."
		
		//Overwriting directory of .yaml file is presented
		if dir := viper.GetViper().GetString("cmd.info.diskUsage.defaulDirectory"); dir != "" {
			defaultDirectory = dir
		}
		usage := du.NewDiskUsage(defaultDirectory)
		fmt.Printf("Free disk space: %dGB in directory %s\n", usage.Free() / 1073741824,defaultDirectory)
	},
}


var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Stats of your cpu",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
	
		pid := int32(os.Getpid())

		proc, err := process.NewProcess(pid)
		if err != nil{
			log.Fatalf("Error creating process object: %v", err)
		}

		

		cpuInfo, err := proc.Username()
		if err != nil {
			log.Fatalf("Error getting cou info: %v", err)
		}
		
		color.Green("\nSystem User:")
		fmt.Printf("  %s\n", cpuInfo)
	},
}

func init() {
	InfoCmd.AddCommand(diskUsageCmd)
	InfoCmd.AddCommand(userCmd)

}
