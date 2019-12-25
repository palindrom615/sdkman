package store

import (
	"github.com/palindrom615/sdkman-cli/conf"
	"path"
	"strings"
	"sync"

	"github.com/prologic/bitcask"
)

var db *bitcask.Bitcask
var once sync.Once

var keyAll = []byte("candidates/all")

func GetStore() *bitcask.Bitcask {
	once.Do(func() {
		e := conf.GetConf()
		db, _ = bitcask.Open(path.Join(e.Dir, "db"))
	})
	return db
}

func GetCandidates() []string {
	candidates, _ := db.Get(keyAll)
	return strings.Split(string(candidates), ",")
}

func SetCandidates(val []string) error {
	e := db.Put(keyAll, []byte(strings.Join(val, ",")))
	return e
}
