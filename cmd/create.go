package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zackdavison/tormalon/internal/service/tormalon/handlers"
)

var createCommand *cobra.Command = &cobra.Command{
	Use:   "create",
	Short: "Used to create projects/repos in scm of different types",
	Run: func(cmd *cobra.Command, args []string) {
		handlers.HandlerInterface()
	},
	Aliases: []string{"c"},
}

func init() {
	rootCmd.AddCommand(createCommand)
}
