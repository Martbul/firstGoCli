package cpu

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color" 
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/process"
	"github.com/spf13/cobra"
)

// cpuCmd represents the cpu command
var tSpecs = &cobra.Command{
	Use:   "tSpecs",
	Short: "Display detailed CPU specifications",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
				color.Cyan("Stats of your CPU")  

info, err := cpu.Info()
		if err != nil {
			fmt.Println("Error fetching CPU info:", err)
			return
		}

	
		jsonData, err := json.MarshalIndent(info[0], "", "  ")
		if err != nil {
			fmt.Println("Error converting CPU info to JSON:", err)
			return
		}

		fmt.Println(string(jsonData))	},
}


var runingThreadsCmd = &cobra.Command{
	Use:   "rThreads",
	Short: "Show active threads of the cpu",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
	
		pid := int32(os.Getpid())

			// Get the process by PID
		p, err := process.NewProcess(pid)
		if err != nil {
			log.Fatalf("Error fetching process with PID %d: %v", pid, err)
		}

		// Fetch the number of threads for the process
		numThreads, err := p.NumThreads()
		if err != nil {
			log.Fatalf("Error fetching thread count for PID %d: %v", pid, err)
		}
		color.Green("\nThreads Usage Info:")
		fmt.Printf("  Number of runing threads: %d\n", numThreads)
	},
}



var cpu2Cmd = &cobra.Command{
	Use:   "cpu2",
	Short: "Show DEVELOPMENT",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
	
		pid := int32(os.Getpid())

		proc, err := process.NewProcess(pid)
		if err != nil{
			log.Fatalf("Error creating process object: %v", err)
		}

		fmt.Println(proc)

		cpuInfo, err := proc.CPUPercent()
	
		if err != nil {
			log.Fatalf("Error getting cou info: %v", err)
		}

		fmt.Println(cpuInfo)


	},
}

func init() {
	CpuCmd.AddCommand(tSpecs)
	CpuCmd.AddCommand(runingThreadsCmd)
	CpuCmd.AddCommand(cpu2Cmd)


}
