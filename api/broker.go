package api

import (
	"github.com/palindrom615/sdk/utils"
	"io"
)

func GetVersion(api string) ([]byte, error) {
	brokerApi := api + "/broker"
	return requestSync(brokerApi + "/version")
}

func GetDownloadSdkmanVersion(api string, versionType string) ([]byte, error) {
	brokerApi := api + "/broker"
	return requestSync(brokerApi + "/download/sdkman/version/" + versionType)
}

func GetDownload(api string, candidate string, version string) (io.ReadCloser, error, string) {
	brokerApi := api + "/broker"
	return download(brokerApi + "/download/" + candidate + "/" + version + "/" + utils.Platform())
}
