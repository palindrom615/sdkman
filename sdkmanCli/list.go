package sdkmanCli

import (
	"github.com/fatih/color"
	"os"
)

func List(candidate string, e *Env) {
	if candidate == "" {
		listCandidates(e)
	}
	listVersions(candidate)
}

func listCandidates(e *Env) {
	if os.Getenv("SDKMAN_AVAILABLE") == "false" {
		color.Red("This command is not available while offline")
	} else {
		Pager(SecureCurl(e.CandidatesApi + "/list"))
	}
}

func listVersions(candidate string) {

}
