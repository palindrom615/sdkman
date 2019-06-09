package utils

import (
	"errors"
	"github.com/fatih/color"
	"os"
)

var (
	ErrNotOnline       = errors.New("sdkman: not available on offline mode")
	ErrNotValidVersion = errors.New("sdkman: not valid version")
	ErrVersionExists   = errors.New("sdkman: already installed version")
	ErrNoCandidate     = errors.New("sdkman: no candidate specified")
	ErrNoArchive       = errors.New("sdkman: archive file not exists")
)

func Check(e error) {
	if e != nil {
		color.Red(e.Error())
		os.Exit(1)
	}
}
