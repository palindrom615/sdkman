package sdkman_cli

import (
	"github.com/fatih/color"
	"os"
)

func ListSdk (candidate string) {
	if candidate == "" {
		listCandidate()
	}
}

func listCandidate() {
	if os.Getenv("SDKMAN_AVAILABLE") == "false" {
		color.Red("This command is not available while offline")
	} else {
		Pager(SecureCurl("https://api.sdkman.io/2/candidates/list"))
	}
}