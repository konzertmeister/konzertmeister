package command

import (
	"github.com/konzertmeister/konzertmeister/internal/pkg/config"
	"github.com/spf13/cobra"
)

var (
	debug bool

	meisterCmd = &cobra.Command{
		Use:   "meister",
		Short: "A tool to bootstrap your directory structures",
		Long: `Konzertmeister is a CLI tool to bootstrap your projects' file and 
		directory structure using git based community and custom stencil packs.`,
	}
)

func Execute() error {
	return meisterCmd.Execute()
}

func init() {
	cobra.OnInitialize(initMeister)
	meisterCmd.PersistentFlags().BoolVar(&debug, "debug", false, "set log level to debug")

}

func initMeister() {
	if debug {
		setLogLevelDebug()
	}

	config.Init(_logger)
}
