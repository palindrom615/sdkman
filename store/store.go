package store

import (
	"github.com/palindrom615/sdk/conf"
	"github.com/prologic/bitcask"
	"path"
	"strings"
)

var keyAll = []byte("candidates/all")

func getStore() *bitcask.Bitcask {
	e := conf.GetConf()
	db, _ := bitcask.Open(path.Join(e.Dir, "db"))
	return db
}

func GetCandidates() []string {
	db := getStore()
	defer db.Close()
	candidates, _ := db.Get(keyAll)
	return strings.Split(string(candidates), ",")
}

func SetCandidates(val []string) error {
	db := getStore()
	defer db.Close()
	return db.Put(keyAll, []byte(strings.Join(val, ",")))
}
