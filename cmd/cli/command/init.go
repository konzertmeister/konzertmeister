/*
	meister init github.com/konzertmeister/meister_command ProjectName
*/
package command

import (
	"github.com/spf13/cobra"

	"github.com/konzertmeister/konzertmeister/internal/pkg/discovery"
)

var initCmd = &cobra.Command{
	Use:   "init PACK_URL PROJECT_NAME [DIRECTORY]",
	Short: "Initialize a new project",
	Long:  `Create a new project from the provided stencil pack`,
	Args:  cobra.MinimumNArgs(2),
	RunE:  bootstrapProject,
}

func init() {
	meisterCmd.AddCommand(initCmd)
}

/*
0- Read url
1- Get stencil pack from url
2- Untar stencil pack
3- Parse stencil pack
4- Prompt user for arguments
5- Execute stancil pack
*/
func bootstrapProject(cmd *cobra.Command, args []string) (err error) {
	_logger.Debug().Strs("Arguments", args).Msg("Arguments recieved")

	packUrl := args[0]
	//projectName := args[1]

	discovery.Initialize(packUrl, conf.CacheDir)

	return
}
