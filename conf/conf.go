package conf

import "sync"

type config struct {
	Dir        string
	Candidates []string
	Api        string
	Platform   string
	Insecure   bool
}

var envInstance *config
var once sync.Once

func GetConf() *config {
	once.Do(func() {
		envInstance = &config{
			Dir:        "C:\\Users\\palin\\.sdkman",
			Candidates: []string{"java", "scala"},
			Api:        "https://api.sdkman.io/2",
			Platform:   "MSYS_NT-10.0",
			Insecure:   false,
		}
	})
	return envInstance
}
