package main

import (
	"github.com/spf13/cobra"
	"github.com/yoyofx/yoyogo/cli/yygctl/cmds"
)

var rootCmd = &cobra.Command{
	Use:   "yoyogo",
	Short: "A generator for Cobra based Applications",
	Long: `yoyogo is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a yoyogo application.`,
}

func main() {
	rootCmd.AddCommand(cmds.VersionCmd)
	rootCmd.AddCommand(cmds.RunCmd)
	rootCmd.AddCommand(cmds.BuildCmd)
	_ = rootCmd.Execute()

}
