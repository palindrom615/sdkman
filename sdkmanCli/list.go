package sdkmanCli

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
)

func List(e *Env, candidate string) {
	if candidate == "" {
		listCandidates(e)
	}
	listVersions(e, candidate)
}

func listCandidates(e *Env) {
	if os.Getenv("SDKMAN_AVAILABLE") == "false" {
		color.Red("This command is not available while offline")
	} else {
		Pager(SecureCurl(e.CandidatesApi + "/list"))
	}
}

func listVersions(e *Env, candidate string) {
	ins, _ := installed(e, candidate)
	curr, _ := currentVersion(e, candidate)
	fmt.Println(SecureCurl(e.CandidatesApi + "/" + candidate + "/" +
		e.Platform + "/versions/list?current=" + curr + "&installed=" + ins))
}

func installed(e *Env, candidate string) (string, error) {
	res := ""
	vers, err := ioutil.ReadDir(e.CandidateDir + string(os.PathSeparator) + candidate)
	for _, ver := range vers {
		res += "," + ver.Name()
	}
	return res, err
}
