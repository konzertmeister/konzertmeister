package command

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

var (
	Logger *zerolog.Logger

	debug bool

	rootCmd = &cobra.Command{
		Use:   "meister",
		Short: "A tool to bootstrap your directory structures",
		Long: `Konzertmeister is a CLI tool to bootstrap your projects' file and 
		directory structure using git based community and custom stencils.`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "set log level to debug")
}

func initConfig() {
	if debug {
		newLogger := Logger.Level(zerolog.DebugLevel)
		Logger = &newLogger
		Logger.Debug().Msg("Debug log activated")
	}
}
