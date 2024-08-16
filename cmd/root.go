package cmd

import (
	"bytes"

	"os"
	"strings"

	"github.com/common-nighthawk/go-figure" 
	"github.com/fatih/color"               

	"github.com/martbul/toolbox/cmd/cpu"
	"github.com/martbul/toolbox/cmd/info"
	"github.com/martbul/toolbox/cmd/memory"
	"github.com/martbul/toolbox/cmd/network"
	"github.com/martbul/toolbox/cmd/temp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "toolbox",
	Short: "",
	Long:  "", 
	Run: func(cmd *cobra.Command, args []string) {
		buf := new(bytes.Buffer)
		cmd.SetOut(buf)
		cmd.Help()

		helpText := buf.String()

		asciiAndHeadingColor := color.New(color.FgCyan, color.Bold)
		subHeadingColor := color.New(color.FgHiBlue)
		// descriptionColor := color.New(color.FgGreen)
		commandColor := color.New(color.FgGreen)
		flagColor := color.New(color.FgYellow)
		defaultColor := color.New(color.FgWhite)

		asciiArt := figure.NewFigure("MonitorX", "standard", true).String()
		asciiAndHeadingColor.Println(asciiArt)

		subHeadingColor.Println("A simple application to track all important parts of your machine\n")

		lines := strings.Split(helpText, "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "Usage:") {
				defaultColor.Println(line)
			} else if strings.HasPrefix(line, "Available Commands:") {
				defaultColor.Println(line)
			} else if strings.HasPrefix(line, "Flags:") {
				defaultColor.Println(line)
			} else if strings.HasPrefix(line, "  ") && strings.Contains(line, " ") {
				if strings.Contains(line, "completion") || strings.Contains(line, "cpu") ||
					strings.Contains(line, "help") || strings.Contains(line, "info") ||
					strings.Contains(line, "memory") || strings.Contains(line, "network") {
					commandColor.Println(line)
				} else if strings.HasPrefix(line, "  ") && (strings.Contains(line, "--") || strings.Contains(line, "-")) {
					flagColor.Println(line)
				} else {
					commandColor.Println(line)
				}
			} else {
				defaultColor.Println(line)
			}
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		color.Red("Error: %v", err) 
		os.Exit(1)
	}
}

func setDefaults() {
	viper.SetDefault("port", "8080")
}

func init() {
	cobra.OnInitialize(initConfig)

	setDefaults()

	if err := viper.WriteConfigAs("toolbox.backup.yaml"); err != nil {
		color.Red("Error writing config backup: %v", err) 
	}

	// Add subcommands
	rootCmd.AddCommand(cpu.CpuCmd)
	rootCmd.AddCommand(info.InfoCmd)
	rootCmd.AddCommand(network.NetworkCmd)
	rootCmd.AddCommand(memory.MemoryCmd)
	rootCmd.AddCommand(temp.TempCmd)

	// Global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.toolbox.yaml)")

	// Local flags
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".toolbox")
	}

	viper.AutomaticEnv() // Read in environment variables that match

	if err := viper.ReadInConfig(); err == nil {
		color.New(color.FgCyan).Printf("Using config file: %s\n", viper.ConfigFileUsed())
	}
}
