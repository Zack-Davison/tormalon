package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tormalon",
	Short: "A CLI tool for creating and pushing application structures to scm tool",
	Long:  `A cloud native approach for creating startup application project structure, to speed up the pace of development of new products`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Default action when invoked with no subcommands
		return cmd.Help()
	},
}

func Execute() error {
	return rootCmd.Execute()
}
