package api

import (
	"github.com/palindrom615/sdk/utils"
	"strconv"
)

func GetAlive(api string) ([]byte, error) {
	return requestSync(api + "/alive")
}

func GetSelfupdate(api string, beta bool) ([]byte, error) {
	return requestSync(api + "/selfupdate?beta=" + strconv.FormatBool(beta))
}

func GetHooks(api string, phase string, candidate string, version string) ([]byte, error) {
	return requestSync(api + "/hooks/" + phase + "/" + candidate + "/" + version + "/" + utils.Platform())
}
