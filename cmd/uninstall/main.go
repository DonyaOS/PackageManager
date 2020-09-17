package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// Register registers uninstall command.
func Register(root *cobra.Command) {
	cmd := &cobra.Command{
		Use:     "uninstall [name of packages]",
		Aliases: []string{"r", "remove"},
		Short:   "removes a package",
		Long:    "Removing package(s) in DonyaOS",
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			leng := strconv.Itoa(len(args))
			fmt.Println("Uninstalling " + leng + " package(s): " + strings.Join(args, " "))
		},
	}

	root.AddCommand(cmd)
}
