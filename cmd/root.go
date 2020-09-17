package cmd

import (
	"fmt"
	"os"

	"github.com/DonyaOS/PackageManager/cmd/install"
	"github.com/DonyaOS/PackageManager/cmd/purge"
	"github.com/DonyaOS/PackageManager/cmd/search"
	"github.com/spf13/cobra"
)

// ExitFailure status code.
const ExitFailure = 1

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	root := &cobra.Command{
		Use:   "donya",
		Short: "Donya Package Manager/System",
	}

	install.Register(root)
	purge.Register(root)
	search.Register(root)

	if err := root.Execute(); err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(ExitFailure)
	}
}
