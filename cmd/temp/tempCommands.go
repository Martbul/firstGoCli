package temp

import (
	// "fmt"
	

	"github.com/spf13/cobra"
	
)

var temperatureCmd = &cobra.Command{
	Use:   "diskUsage",
	Short: "Prints disk usage of the current directory",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
// 	   // Initialize the ACPI instance
//     acpiInfo, err := acpi.Cooling()
// 	  if err != nil{
// 		fmt.Println(err)
// 	  }
//    fmt.Println(acpiInfo)
// },
	},
}

func init() {
	TempCmd.AddCommand(temperatureCmd)
	
}
