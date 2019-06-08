package utils

import (
	"errors"
	"github.com/fatih/color"
	"os"
)

var (
	ErrNotOnline       = errors.New("not available on offline mode")
	ErrNotValidVersion = errors.New("not valid version")
	ErrVersionExists   = errors.New("already installed version")
	ErrNoCandidate     = errors.New("no candidate specified")
)

func ThrowError(e error) {
	color.Red(e.Error())
	os.Exit(1)
}
