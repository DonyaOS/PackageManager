package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var cmdSearch = &cobra.Command{
	Use:     "search [name of packages]",
	Aliases: []string{"s"},
	Short:   "searches for a package",
	Long:    "Searching for a package(s) in DonyaOS",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		leng := strconv.Itoa(len(args))
		fmt.Println("Searching for " + leng + " package(s): " + strings.Join(args, " "))
	},
}

func init() {
	rootCmd.AddCommand(cmdSearch)
}
