package sdkmanCli

import (
	"github.com/fatih/color"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func List(e *Env, candidate string) {
	if candidate == "" {
		if os.Getenv("SDKMAN_AVAILABLE") == "false" {
			color.Red("This command is not available while offline")
		} else {
			Pager(string(download(listCandidatesApi(e))))
		}
	} else {
		ins, _ := installed(e, candidate)
		curr, _ := currentVersion(e, candidate)
		Pager(string(download(listCandidateApi(e, candidate, curr, ins))))
	}
}

func listCandidatesApi(e *Env) string {
	return e.CandidatesApi + "/list"
}

func listCandidateApi(e *Env, candidate string, current string, installed []string) string {
	return e.CandidatesApi + "/" + candidate + "/" + e.Platform + "/versions/list?" +
		"current=" + current + "&installed=" + strings.Join(installed, ",")
}
