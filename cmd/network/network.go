package network

import (
	"github.com/spf13/cobra"
)

// memoryCmd represents the memory command
var NetworkCmd = &cobra.Command{
	Use:   "network",
	Short: "Flex on your friends with fast internet speed",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
cmd.Help()	},
}

func init() {

}
