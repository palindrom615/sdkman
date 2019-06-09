package api

import (
	"io"
	"io/ioutil"
	"strings"
)

var brokerApi = e.Api + "/broker"

func GetVersion() ([]byte, error) {
	return downloadSync(brokerApi + "/version")
}

func GetDownloadSdkmanVersion(versionType string) ([]byte, error) {
	return downloadSync(brokerApi + "/download/sdkman/version/" + versionType)
}

func GetDownload(candidate string, version string) (io.ReadCloser, error) {
	resp, err := download(brokerApi + "/download/" + candidate + "/" + version + "/" + e.Platform)
	if resp == nil {
		empty := ioutil.NopCloser(strings.NewReader(""))
		return empty, err
	}
	return resp.Body, err
}
