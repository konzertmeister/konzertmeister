package command

import (
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Return the current version of Konzertmeister",
	Run: func(cmd *cobra.Command, args []string) {
		Logger.Info().Msg("Konzertmeister v1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
