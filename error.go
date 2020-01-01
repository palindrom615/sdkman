package sdkman

import (
	"errors"
)

var (
	ErrNotOnline   = errors.New("sdkman: not available on offline mode")
	ErrNoVer       = errors.New("sdkman: not valid version")
	ErrNoCand      = errors.New("sdkman: no valid candidate")
	ErrVerNotIns   = errors.New("sdkman: specified version not installed")
	ErrArcNotIns   = errors.New("sdkman: archive file not exists")
	ErrNoCurrCands = errors.New("sdkman: no candidates are in use")
	ErrVerExists   = errors.New("sdkman: already installed version")
	ErrVerInsFail  = errors.New("sdkman: installation failed")
)

func ErrNoCurrSdk(cand string) error {
	return errors.New("sdkman: not using any version of " + cand)
}
