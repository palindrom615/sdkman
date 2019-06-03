package api

import "strconv"

func GetAlive() []byte {
	return download(e.Api + "/alive")
}

func GetSelfupdate(beta bool) []byte {
	return download(e.Api + "/selfupdate?beta=" + strconv.FormatBool(beta))
}

func GetHooks(phase string, candidate string, version string) []byte {
	return download(e.Api + "/hooks/" + phase + "/" + candidate + "/" + version + "/" + e.Platform)
}
