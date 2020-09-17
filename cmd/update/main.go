package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// Register registers update command.
func Register(root *cobra.Command) {
	cmd := &cobra.Command{
		Use:     "update [name of packages]",
		Aliases: []string{"u"},
		Short:   "updates the list of packages",
		Long:    "Updating the list of packages in DonyaOS",
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			leng := strconv.Itoa(len(args))
			fmt.Println("Updating " + leng + " package(s): " + strings.Join(args, " "))
		},
	}

	root.AddCommand(cmd)
}
