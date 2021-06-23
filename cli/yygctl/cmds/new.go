package cmds

import (
	"fmt"
	"github.com/spf13/cobra"
)

var l bool
var NewCmd = &cobra.Command{
	Use:   "new",
	Short: "create new yoyogo demo by template",
	Long:  "create new yoyogo demo by template",
	//Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if l {
			fmt.Println("console / webapi / mvc / grpc / xxl-job")
		}
	},
}

func init() {
	NewCmd.Flags().BoolVarP(&l, "templates", "l", false, "list of template")
}
