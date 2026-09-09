package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version string = "1.0.0"

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Print Tormalon CLI version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Current Tormalon Version: %s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCommand)
}
