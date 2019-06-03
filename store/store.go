package store

import (
	"path"
	"sdkman-cli/conf"
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
