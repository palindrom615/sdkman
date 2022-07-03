package custom_errors

import (
	"errors"
)

var (
	// ErrNotOnline is caused by registry is not online
	ErrNotOnline = errors.New("not available on offline mode")
	//ErrNoVer is caused if the version does not exists in registry
	ErrNoVer = errors.New("not valid version")
	// ErrNoCand is caused if the candidate does not exists in registry
	ErrNoCand = errors.New("no valid candidate")
	// ErrVerNotIns is caused when the version is not installed
	ErrVerNotIns = errors.New("specified version not installed")
	// ErrArcNotIns is caused when the archive file does not exists in local directory
	ErrArcNotIns = errors.New("archive file not exists")
	// ErrNoCurrCands is caused when no sdk is used now
	ErrNoCurrCands = errors.New("no candidates are in use")
	// ErrVerExists is caused when trying to install already installed version
	ErrVerExists = errors.New("already installed version")
	// ErrVerInsFail is caused when installing version is failed
	ErrVerInsFail = errors.New("installation failed")
)

// ErrNoCurrSdk is caused when no sdk of the candidate is in use
func ErrNoCurrSdk(cand string) error {
	return errors.New("not using any version of " + cand)
}
