package conf

import (
	"os"
	"path"
	"runtime"
	"sync"
)

type config struct {
	Dir      string
	Api      string
	Platform string
	Insecure bool
}

var envInstance *config
var once sync.Once

func GetConf() *config {
	once.Do(func() {
		home, _ := os.UserHomeDir()
		platform := runtime.GOOS
		if platform == "windows" {
			platform = "msys_nt-10.0"
		}
		envInstance = &config{
			Dir:      path.Join(home, ".sdkman"),
			Api:      "https://api.sdkman.io/2",
			Platform: platform,
			Insecure: false,
		}
	})
	return envInstance
}
