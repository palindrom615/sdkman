package conf

import (
	"os"
	"path"
	"sync"
)

type config struct {
	Dir        string
	Api        string
	Platform   string
	Insecure   bool
}

var envInstance *config
var once sync.Once

func GetConf() *config {
	once.Do(func() {
		home, _ := os.UserHomeDir()
		envInstance = &config{
			Dir:        path.Join(home, ".sdkman"),
			Api:        "https://api.sdkman.io/2",
			Platform:   "MSYS_NT-10.0",
			Insecure:   false,
		}
	})
	return envInstance
}
