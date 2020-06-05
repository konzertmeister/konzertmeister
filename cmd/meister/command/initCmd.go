package command

import (
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new project",
	Long:  `Create a new project from the provided stencil`,
	Run: func(cmd *cobra.Command, args []string) {
		Logger.Info().Msg("Test!")
		Logger.Debug().Msg("DEBUG!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
