package api

import "strconv"

func GetAlive() ([]byte, error) {
	return download(e.Api + "/alive")
}

func GetSelfupdate(beta bool) ([]byte, error) {
	return download(e.Api + "/selfupdate?beta=" + strconv.FormatBool(beta))
}

func GetHooks(phase string, candidate string, version string) ([]byte, error) {
	return download(e.Api + "/hooks/" + phase + "/" + candidate + "/" + version + "/" + e.Platform)
}
