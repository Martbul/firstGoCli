package cpu

import (
	"bytes"
	"strings"

	"github.com/fatih/color" 
	"github.com/spf13/cobra"
)

// cpuCmd represents the cpu command
var CpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "Find every important cpu stat",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {


		commandColor := color.New(color.FgGreen, color.Bold) 
		flagColor := color.New(color.FgYellow)                

		
		buf := new(bytes.Buffer)
		cmd.SetOut(buf)
		cmd.Help()

		helpText := buf.String()

		
		lines := strings.Split(helpText, "\n")
		for _, line := range lines {
			
			if strings.HasPrefix(line, "  ") && strings.Contains(line, " ") {
		
				if strings.Contains(line, "completion") || strings.Contains(line, "cpu") || strings.Contains(line, "CPU") || strings.Contains(line, "Show")|| strings.Contains(line, "Display")||
					strings.Contains(line, "help") || strings.Contains(line, "info") ||
					strings.Contains(line, "memory") || strings.Contains(line, "network") {
					commandColor.Println(line)
				} else {
					flagColor.Println(line)
				}
			} else {
				color.White(line) 
			}
		}
		},
}


func init() {
	

}
