package utils

import (
	"errors"
	"github.com/fatih/color"
	"github.com/palindrom615/sdkman-cli/store"
	"os"
)

var (
	ErrNotOnline       = errors.New("sdkman: not available on offline mode")
	ErrNotValidVersion = errors.New("sdkman: not valid version")
	ErrVersionExists   = errors.New("sdkman: already installed version")
	ErrNoCandidate     = errors.New("sdkman: no valid candidate")
	ErrNoVersion       = errors.New("sdkman: specified version not installed")
	ErrNoArchive       = errors.New("sdkman: archive file not exists")
)

func Check(e error) {
	if e != nil {
		color.Red(e.Error())
		os.Exit(1)
	}
}

func IsCandidateValid(candidate string) bool {
	for _, can := range store.GetCandidates() {
		if can == candidate {
			return true
		}
	}
	return false
}
