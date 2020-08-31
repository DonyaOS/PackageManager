package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	var cmdPrint = &cobra.Command{
		Use:   "install [name of packages]",
		Short: "Installing package(s) in DonyaOS",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			leng := strconv.Itoa(len(args))
			fmt.Println("Install " + leng + " package(s): " + strings.Join(args, " "))
		},
	}

	var rootCmd = &cobra.Command{Use: "donya"}
	rootCmd.AddCommand(cmdPrint)
	rootCmd.Execute()
}
