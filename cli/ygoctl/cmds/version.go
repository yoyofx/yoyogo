package cmds

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yoyofx/yoyogo"
	"github.com/yoyofx/yoyogo/abstractions/platform/consolecolors"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of yoyogo",
	Long:  `All software has versions. This is yoyogo's`,
	Run: func(cmd *cobra.Command, args []string) {
		logo := yoyogo.Logo
		fmt.Println(consolecolors.Blue(string(logo)))
		fmt.Println(" ")
		fmt.Printf("%s   (version:  %s)", consolecolors.Green(":: YoyoGo ::"), consolecolors.Blue(yoyogo.Version))

		fmt.Print(consolecolors.Blue(`
light and fast , dependency injection based micro-service framework written in Go.
`))

		fmt.Println(" ")
	},
}
