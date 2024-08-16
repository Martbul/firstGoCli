package temp

import (
	"bytes"
	"strings"

	"github.com/fatih/color" // Import color package
	"github.com/spf13/cobra"
)

// cpuCmd represents the cpu command
var TempCmd = &cobra.Command{
	Use:   "temp",
	Short: "Find every important cpu stat",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
// Define colors
		headingColor := color.New(color.FgCyan, color.Bold)    // Light Cyan for headings
		commandColor := color.New(color.FgGreen, color.Bold)  // Light Green for commands
		flagColor := color.New(color.FgYellow)                // Light Yellow for flags

		// Custom help output
		buf := new(bytes.Buffer)
		cmd.SetOut(buf)
		cmd.Help()

		helpText := buf.String()

		// Print the modified help text with colors
		lines := strings.Split(helpText, "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "Usage:") {
				headingColor.Println(line)
			} else if strings.HasPrefix(line, "Available Commands:") {
				headingColor.Println(line)
			} else if strings.HasPrefix(line, "Flags:") {
				headingColor.Println(line)
			} else if strings.HasPrefix(line, "  ") && strings.Contains(line, " ") {
				// Detect command lines
				if strings.Contains(line, "completion") || strings.Contains(line, "cpu") ||
					strings.Contains(line, "help") || strings.Contains(line, "info") ||
					strings.Contains(line, "memory") || strings.Contains(line, "network") {
					commandColor.Println(line)
				} else {
					flagColor.Println(line)
				}
			} else {
				color.White(line) // Default color for other text
			}
		}
		},
}


func init() {
	

}
