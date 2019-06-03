package store

import (
	"path"
	"sdkman-cli/conf"
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

func SetCandidates(val []byte) error {
	e := db.Put("candidates/all", val)
	return e
}
