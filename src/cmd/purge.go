package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var cmdPurge = &cobra.Command{
	Use:     "purge [name of packages]",
	Aliases: []string{"p"},
	Short:   "removes all orphaned packages",
	Long:    "Removing all orphaned package(s) in DonyaOS",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		leng := strconv.Itoa(len(args))
		fmt.Println("Purging " + leng + " package(s): " + strings.Join(args, " "))
	},
}

func init() {
	rootCmd.AddCommand(cmdPurge)
}
