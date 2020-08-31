package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
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
	var cmdUninstall = &cobra.Command{
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
	var cmdUpdate = &cobra.Command{
        Use:     "update [name of packages]",
        Aliases: []string{"u"},
        Short:   "updates the list of packages",
        Long:   "Updating the list of packages in DonyaOS",
        Args:    cobra.MinimumNArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
            leng := strconv.Itoa(len(args))
            fmt.Println("Updating " + leng + " package(s): " + strings.Join(args, " "))
        },
	}

	var rootCmd = &cobra.Command{
	    Use:   "donya",
	    Short: "Donya package manager is a tool for installing different packages.",
    }
	rootCmd.AddCommand(cmdInstall, cmdPurge, cmdSearch, cmdUninstall, cmdUpdate)
	rootCmd.Execute()
}
