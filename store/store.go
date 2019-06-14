package store

import (
	"github.oom/palindrom615/sdkman-cli/conf"
	"path"
	"strings"
	"sync"

	"github.com/prologic/bitcask"
)

var db *bitcask.Bitcask
var once sync.Once

func GetStore() *bitcask.Bitcask {
	once.Do(func() {
		e := conf.GetConf()
		db, _ = bitcask.Open(path.Join(e.Dir, "db"))
	})
	return db
}

func GetCandidates() []string {
	candidates, _ := db.Get("candidates/all")
	return strings.Split(string(candidates), ",")
}

func SetCandidates(val []string) error {
	e := db.Put("candidates/all", []byte(strings.Join(val, ",")))
	return e
}
