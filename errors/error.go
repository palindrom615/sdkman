package errors

import (
	"errors"
)

var (
	// ErrNotOnline is caused by registry is not online
	ErrNotOnline = errors.New("sdkman: not available on offline mode")
	//ErrNoVer is caused if the version does not exists in registry
	ErrNoVer = errors.New("sdkman: not valid version")
	// ErrNoCand is caused if the candidate does not exists in registry
	ErrNoCand = errors.New("sdkman: no valid candidate")
	// ErrVerNotIns is caused when the version is not installed
	ErrVerNotIns = errors.New("sdkman: specified version not installed")
	// ErrArcNotIns is caused when the archive file does not exists in local directory
	ErrArcNotIns = errors.New("sdkman: archive file not exists")
	// ErrNoCurrCands is caused when no sdk is used now
	ErrNoCurrCands = errors.New("sdkman: no candidates are in use")
	// ErrVerExists is caused when trying to install already installed version
	ErrVerExists = errors.New("sdkman: already installed version")
	// ErrVerInsFail is caused when installing version is failed
	ErrVerInsFail = errors.New("sdkman: installation failed")
)

// ErrNoCurrSdk is caused when no sdk of the candidate is in use
func ErrNoCurrSdk(cand string) error {
	return errors.New("sdkman: not using any version of " + cand)
}
