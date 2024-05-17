package cmd

import (
	"github.com/ldd27/go-starter-kit/pkg/version"
	"github.com/spf13/cobra"
)

func init() {
	RootCMD.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "print version",
	Long:    "print version about application",
	Run: func(cmd *cobra.Command, args []string) {
		version.Print()
	},
}
