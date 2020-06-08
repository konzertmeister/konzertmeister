/*
	meister init github.com/konzertmeister/meister_command ProjectName
*/
package command

import (
	"github.com/konzertmeister/konzertmeister/pkg/stencil"
	"github.com/spf13/cobra"
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
	//stencil.NewStencilPack(args[0])
	_logger.Debug().Strs("Arguments", args).Msg("Arguments recieved")
	return
}

func parseStencil(path string) (st *stencil.Stencil, err error) {
	st, err = stencil.NewStencil(path)
	if err != nil {
		_logger.Err(err).Msgf("Failed to parse stencil `%v`: %v", path, err)
		return
	}

	return
}
