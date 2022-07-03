package script

import (
	"fmt"
	"github.com/palindrom615/sdkman/sdk"
	"github.com/yargevad/filepathx"
	"path/filepath"
	"runtime"
	"strings"
)

func getPathsHomes(sdks []sdk.Sdk) ([]string, []envVar) {
	paths := []string{}
	homes := []envVar{}
	for _, s := range sdks {
		candHome := filepath.Join(s.SdkHome, "candidates", s.Candidate, "current")

		binPath := filepath.Join(candHome, "bin")
		homePath := candHome
		if s.Candidate == "java" && runtime.GOOS == "darwin" {
			// fix for macOS
			matches, _ := filepathx.Glob(fmt.Sprintf("%s/**/javac", homePath))
			binPath = filepath.Join(matches[0], "..")
			homePath = filepath.Join(binPath, "..")
		}
		paths = append(paths, binPath)
		homes = append(homes, envVar{fmt.Sprintf("%s_HOME", strings.ToUpper(s.Candidate)), homePath})
	}
	return paths, homes
}
