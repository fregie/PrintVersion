package version

import (
	"fmt"
)

var (
	Name          = "Unknown"
	Version       = "Unknown"
	BuildTime     = "Unknown"
	GitCommitSHA1 = "Unknown"
	Describe      = ""
)

func PrintVersion() {
	fmt.Printf("----------------------- %s -----------------------\n", Name)
	if Describe != "" {
		fmt.Print("  " + Describe + "\n\n")
	}
	fmt.Printf("  Version:       %s\n", Version)
	fmt.Printf("  BuildTime:     %s\n", BuildTime)
	fmt.Printf("  GitCommitSHA1: %s\n", GitCommitSHA1)
}
