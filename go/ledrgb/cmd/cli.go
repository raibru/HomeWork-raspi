package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"rbr.de/ledrgb/rpi"
)

var (
	ledRed     bool
	ledGreen   bool
	ledBlue    bool
	prtVersion bool
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&ledRed, "red", "r", false, "enable led in red")
	rootCmd.PersistentFlags().BoolVarP(&ledGreen, "green", "g", false, "enable led in green")
	rootCmd.PersistentFlags().BoolVarP(&ledBlue, "blue", "b", false, "enable led in blue")
	rootCmd.PersistentFlags().BoolVarP(&prtVersion, "version", "v", false, "display rgbled version")
}

var rootCmd = &cobra.Command{
	Use:   "rgbled",
	Short: "Set a RGB-LED on Raspberry PI to red, green or blue light",
	Long:  "Set a RGB-LED on Raspberry PI to red, green or blue light",
	Run: func(cmd *cobra.Command, args []string) {
		if err := handleParam(cmd, args); err != nil {
			cmd.Help()
			fmt.Println("\nCommand has parsing error: ", err)
			os.Exit(1)
		}
	},
}

// Execute the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Execute cmd has error", err)
		os.Exit(1)
	}
}

// handleParam parameter evaluation
func handleParam(cmd *cobra.Command, args []string) error {
	if prtVersion {
		PrintVersion(os.Stdout)
		return nil
	}

	var r, g, b uint32

	r = 0
	if ledRed {
		r = 1024
	}

	g = 0
	if ledGreen {
		g = 1024
	}

	b = 0
	if ledBlue {
		b = 1024
	}

	if err := rpi.ApplyLedRgb(r, g, b); err != nil {
		cmd.Help()
		return err
	}

	return nil
}
