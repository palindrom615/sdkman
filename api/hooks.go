package api

import "strconv"

func GetAlive() ([]byte, error) {
	return requestSync(e.Api + "/alive")
}

func GetSelfupdate(beta bool) ([]byte, error) {
	return requestSync(e.Api + "/selfupdate?beta=" + strconv.FormatBool(beta))
}

func GetHooks(phase string, candidate string, version string) ([]byte, error) {
	return requestSync(e.Api + "/hooks/" + phase + "/" + candidate + "/" + version + "/" + e.Platform)
}
