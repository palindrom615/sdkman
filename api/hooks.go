package api

import (
	"github.com/palindrom615/sdk/local"
	"github.com/palindrom615/sdk/utils"
	"strconv"
)

func GetAlive(api string) ([]byte, error) {
	return requestSync(api + "/alive")
}

func GetSelfupdate(api string, beta bool) ([]byte, error) {
	return requestSync(api + "/selfupdate?beta=" + strconv.FormatBool(beta))
}

func GetHooks(api string, phase string, sdk local.Sdk) ([]byte, error) {
	return requestSync(api + "/hooks/" + phase + "/" + sdk.Candidate + "/" + sdk.Version + "/" + utils.Platform())
}
