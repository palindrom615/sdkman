package store

import (
	"fmt"
	"github.com/palindrom615/sdkman/api"
	"github.com/palindrom615/sdkman/errors"
	"github.com/palindrom615/sdkman/pkgs"
	"io/ioutil"
	"path"
	"strings"
)

type Store struct {
	SdkHome string
}

func (store Store) GetCandidates() []string {
	candidates, err := ioutil.ReadFile(path.Join(store.SdkHome, "candidates.txt"))
	if err != nil {
		return []string{}
	}
	return strings.Split(string(candidates), ",")
}

func (store Store) SetCandidates(val []string) error {
	return ioutil.WriteFile(path.Join(store.SdkHome, "candidates.txt"), []byte(strings.Join(val, ",")), 0666)
}

func (store Store) Update(registry string) error {
	freshCsv, netErr := api.GetAll(registry)
	if netErr != nil {
		return errors.ErrNotOnline
	}
	fresh := pkgs.NewStrSet(freshCsv...)
	cachedCsv := store.GetCandidates()
	cached := pkgs.NewStrSet(cachedCsv...)

	added := fresh.Difference(cached)
	obsoleted := cached.Difference(fresh)

	if added.Size() != 0 {
		fmt.Println("Adding new candidates: " + strings.Join(added.List(), ", "))
	}
	if obsoleted.Size() != 0 {
		fmt.Println("Removing obsolete candidates: " + strings.Join(obsoleted.List(), ", "))
	}
	return store.SetCandidates(freshCsv)
}

func (store Store) HasCandidate(candidate string) bool {
	for _, can := range store.GetCandidates() {
		if can == candidate {
			return true
		}
	}
	return false
}
