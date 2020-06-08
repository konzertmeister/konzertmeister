package command

import (
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Return the current version of Konzertmeister",
	Run: func(cmd *cobra.Command, args []string) {
		_logger.Info().Msg("Konzertmeister v1.0")
	},
}

func init() {
	meisterCmd.AddCommand(versionCmd)
}
