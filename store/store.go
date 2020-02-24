package store

import (
	"github.com/prologic/bitcask"
	"path"
	"strings"
)

var keyAll = []byte("candidates/all")

func getStore(dir string) *bitcask.Bitcask {
	db, _ := bitcask.Open(path.Join(dir, "db"))
	return db
}

func GetCandidates(dir string) []string {
	db := getStore(dir)
	defer db.Close()
	candidates, _ := db.Get(keyAll)
	return strings.Split(string(candidates), ",")
}

func SetCandidates(dir string, val []string) error {
	db := getStore(dir)
	defer db.Close()
	return db.Put(keyAll, []byte(strings.Join(val, ",")))
}
