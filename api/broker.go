package api

import (
	"io"
)

var brokerApi = e.Api + "/broker"

func GetVersion() ([]byte, error) {
	return requestSync(brokerApi + "/version")
}

func GetDownloadSdkmanVersion(versionType string) ([]byte, error) {
	return requestSync(brokerApi + "/download/sdkman/version/" + versionType)
}

func GetDownload(candidate string, version string) (io.ReadCloser, error, string) {
	return download(brokerApi + "/download/" + candidate + "/" + version + "/" + e.Platform)
}
