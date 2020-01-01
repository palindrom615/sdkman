package sdkman

import (
	"fmt"
	"os"
	"strings"

	"github.com/scylladb/go-set/strset"
)

type envVar struct {
	name string
	val  string
}

func evalBash(paths []string, envVars []envVar) {
	fmt.Println("export PATH=" + strings.Join(paths, ":") + ":$PATH")
	for _, v := range envVars {
		fmt.Println("export " + v.name + "=" + v.val)
	}
}

func evalFish(paths []string, envVars []envVar) {
	fmt.Println("set -x PATH " + strings.Join(paths, " ") + " $PATH")
	for _, v := range envVars {
		fmt.Println("set -x " + v.name + " " + v.val)
	}
}

func evalPosh(paths []string, envVars []envVar) {
	fmt.Printf("$env:Path = \"%s;\" + $env:Path;", strings.Join(paths, ";"))
	for _, v := range envVars {
		fmt.Printf("$env:%s = \"%s\";", v.name, v.val)
	}
	fmt.Println()
}

func evalWindows(paths []string, envVars []envVar) {
	path := strset.New(strings.Split(os.Getenv("Path"), ";")...)
	path.Merge(strset.New(paths...))
	fmt.Printf("[Environment]::SetEnvironmentVariable(\"Path\", \"%s\", [System.EnvironmentVariableTarget]::User);", strings.Join(path.List(), ";"))
	for _, v := range envVars {
		fmt.Printf("[Environment]::SetEnvironmentVariable(\"%s\", \"%s\", [System.EnvironmentVariableTarget]::User);", v.name, v.val)
	}
	fmt.Println()
}
