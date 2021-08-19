package store

import (
	"io/ioutil"
	"path"
	"strings"
)

func GetCandidates(dir string) []string {
	candidates, err := ioutil.ReadFile(path.Join(dir, "candidates"))
	if err != nil {
		return []string{}
	}
	return strings.Split(string(candidates), ",")
}

func SetCandidates(dir string, val []string) error {
	return ioutil.WriteFile(path.Join(dir, "candidates"), []byte(strings.Join(val, ",")), 0666)
}
