package api

import (
	"io"
)

var brokerApi = e.Api + "/broker"

func GetVersion() ([]byte, error) {
	return downloadSync(brokerApi + "/version")
}

func GetDownloadSdkmanVersion(versionType string) ([]byte, error) {
	return downloadSync(brokerApi + "/download/sdkman/version/" + versionType)
}

func GetDownload(candidate string, version string) (io.ReadCloser, error) {
	return download(brokerApi + "/download/" + candidate + "/" + version + "/" + e.Platform)
}
