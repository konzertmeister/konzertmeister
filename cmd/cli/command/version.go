package command

import (
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Return the current version of Konzertmeister",
	Run: func(cmd *cobra.Command, args []string) {
		_logger.Info().Msg("Konzertmeister v1.0")
		_logger.Info().Msg("---")
		_logger.Info().Msgf("Current configuration:\n%v", conf)
	},
}

func init() {
	meisterCmd.AddCommand(versionCmd)
}
