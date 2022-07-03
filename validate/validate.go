package validate

import (
	"github.com/palindrom615/sdkman/errors"
	"github.com/palindrom615/sdkman/store"
)

func CheckValidCand(root string, candidate string) error {
	for _, can := range store.GetCandidates(root) {
		if can == candidate {
			return nil
		}
	}
	return errors.ErrNoCand
}
