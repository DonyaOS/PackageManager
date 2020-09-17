package main

import (
	"fmt"

	"github.com/DonyaOS/PackageManager/cmd"
)

// nolint: gocheckglobals
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	fmt.Printf("gosimac %s, commit %s, built at %s\n", version, commit, date)

	cmd.Execute()
}
