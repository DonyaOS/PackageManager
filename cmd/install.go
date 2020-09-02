package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var cmdInstall = &cobra.Command{
	Use:     "install [name of packages]",
	Aliases: []string{"i"},
	Short:   "installs a package",
	Long:    "Installing package(s) in DonyaOS",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		leng := strconv.Itoa(len(args))
		fmt.Println("Installing " + leng + " package(s): " + strings.Join(args, " "))
	},
}

func init() {
	rootCmd.AddCommand(cmdInstall)
}
