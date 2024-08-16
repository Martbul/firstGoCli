package memory

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color" 
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/process"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// cpuCmd represents the cpu command
var stats = &cobra.Command{
	Use:   "stats",
	Short: "Find every spec your cpu has",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		memoryUsage, err := mem.VirtualMemory()
		convJsonMemUsage, err := json.MarshalIndent(memoryUsage, "", " ")
		if err != nil {
			fmt.Println("Error converting Memory data to JSON:", err)
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(convJsonMemUsage))

	},
}
var usageCmd = &cobra.Command{
	Use:   "usage",
	Short: "Find every spec your cpu has",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Get the current process's PID
		pid := int32(os.Getpid())

		// Create a new Process object
		proc, err := process.NewProcess(pid)
		if err != nil {
			log.Fatalf("Error creating process object: %v", err)
		}

		// Get memory percentage usage
		// memPercent, err := proc.MemoryPercent()
		// if err != nil {
		// log.Fatalf("Error getting memory usage percentage: %v", err)
		// }

		// Get detailed memory info (RSS, VMS, etc.)
		memInfo, err := proc.MemoryInfo()
		if err != nil {
			log.Fatalf("Error getting memory info: %v", err)
		}

		// Get the total system memory
		vmStat, err := mem.VirtualMemory()
		if err != nil {
			log.Fatalf("Error getting system memory info: %v", err)
		}
		color.Green("\nMemory Usage Info:")
		fmt.Printf("  RSS (physical memory): %.2f MB\n", float64(memInfo.RSS)/1024/1024)
		fmt.Printf("  VMS (virtual memory): %.2f MB\n", float64(memInfo.VMS)/1024/1024)

		// Print total system memory and used memory
		fmt.Printf("  Total available memory: %v bytes (%.2f GB)\n", vmStat.Total, float64(vmStat.Total)/1024/1024/1024)
		fmt.Printf("  Memory usage: %.2f%%\n", vmStat.UsedPercent)
	},
}

// var usageCmd = &cobra.Command{
// 	Use:   "usage [pid]",
// 	Short: "Display memory usage of a specific process by PID",
// 	Long:  "Shows the memory usage as a percentage for the process identified by its PID.",
// 	Args:  cobra.ExactArgs(1),
// 	Run: func(cmd *cobra.Command, args []string) {
// 		// Convert PID argument to integer
// 		pid, err := strconv.Atoi(args[0])
// 		if err != nil {
// 			log.Fatalf("Invalid PID provided: %v", err)
// 		}

// 		// Get the process by PID
// 		p, err := process.NewProcess(int32(pid))
// 		if err != nil {
// 			log.Fatalf("Error fetching process with PID %d: %v", pid, err)
// 		}

// 		// Fetch the memory percentage usage for the process
// 		memoryPercent, err := p.MemoryPercent()
// 		if err != nil {
// 			log.Fatalf("Error fetching memory percentage for PID %d: %v", pid, err)
// 		}

// 		fmt.Printf("Memory Usage for PID %d: %.2f%%\n", pid, memoryPercent)
// 	},
// }

func init() {
	MemoryCmd.AddCommand(stats)
	MemoryCmd.AddCommand(usageCmd)

}
