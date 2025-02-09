package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"rbr.de/ledblink/rpi"
)

var rootCmd = &cobra.Command{
	Use:   "blink",
	Short: "Let a LCD on Raspberry PI blink",
	Long:  `Let a LCD on Raspberry PI blink`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := handleParam(cmd, args); err != nil {
			cmd.Help()
			fmt.Println("\nRoot command has parsing error: ", err)
			os.Exit(1)
		}
	},
}

// Execute the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Execute root cmd has error", err)
		os.Exit(1)
	}
}

// handleParam parameter evaluation
func handleParam(cmd *cobra.Command, args []string) error {
	if prtVersion {
		PrintVersion(os.Stdout)
		return nil
	}

	err := rpi.ApplyBlink(interval, ntimes)
	if err != nil {
		cmd.Help()
		return err
	}

	return nil
}
