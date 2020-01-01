package sdkman

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"strconv"
	"strings"
)

func getBroadcastLatestID(api string) ([]byte, error) {
	broadcastAPI := api + "/broadcast"
	return requestSync(broadcastAPI + "/latest/id")
}

func getBroadcastLatest(api string) ([]byte, error) {
	broadcastAPI := api + "/broadcast"
	return requestSync(broadcastAPI + "/latest")
}

func getBroadcastID(api string, id string) ([]byte, error) {
	broadcastAPI := api + "/broadcast"
	return requestSync(broadcastAPI + "/" + id)
}

func getVersion(api string) ([]byte, error) {
	brokerAPI := api + "/broker"
	return requestSync(brokerAPI + "/version")
}

func getDownloadSdkmanVersion(api string, versionType string) ([]byte, error) {
	brokerAPI := api + "/broker"
	return requestSync(brokerAPI + "/download/sdkman/version/" + versionType)
}

func getDownload(api string, sdk Sdk) (io.ReadCloser, string, error) {
	brokerAPI := api + "/broker"
	return download(brokerAPI + "/download/" + sdk.Candidate + "/" + sdk.Version + "/" + platform())
}

func getDefault(api string, candidate string) (string, error) {
	candidatesAPI := api + "/candidates"

	res, err := requestSync(candidatesAPI + "/default/" + candidate)
	return string(res), err
}

func getValidate(api string, sdk Sdk) (bool, error) {
	candidatesAPI := api + "/candidates"
	url := fmt.Sprintf("%s/validate/%s/%s/%s", candidatesAPI, sdk.Candidate, sdk.Version, platform())
	res, err := requestSync(url)
	return string(res) == "valid", err
}

func getList(api string) (io.ReadCloser, error) {
	candidatesAPI := api + "/candidates"

	return request(candidatesAPI + "/list")
}

func getVersionsList(api string, currentSdk Sdk, installed []Sdk) (io.ReadCloser, error) {
	candidatesAPI := api + "/candidates"
	installedVers := ""
	for _, sdk := range installed {
		installedVers += sdk.Version + ","
	}
	url := fmt.Sprintf("%s/%s/%s/versions/list?current=%s&installed=%s", candidatesAPI, currentSdk.Candidate, platform(), currentSdk.Version, installedVers)
	return request(url)
}

func getAll(api string) ([]string, error) {
	candidatesAPI := api + "/candidates"
	res, err := requestSync(candidatesAPI + "/all")
	return strings.Split(string(res), ","), err
}

func getVersionsAll(api string, candidate string) ([]byte, error) {
	candidatesAPI := api + "/candidates"
	return requestSync(candidatesAPI + "/" + candidate + "/" + platform() + "/versions/all")
}

func getAlive(api string) ([]byte, error) {
	return requestSync(api + "/alive")
}

func getSelfupdate(api string, beta bool) ([]byte, error) {
	return requestSync(api + "/selfupdate?beta=" + strconv.FormatBool(beta))
}

func getHooks(api string, phase string, sdk Sdk) ([]byte, error) {
	return requestSync(api + "/hooks/" + phase + "/" + sdk.Candidate + "/" + sdk.Version + "/" + platform())
}

func wrapResponseBody(r *http.Response) io.ReadCloser {
	if r == nil {
		return ioutil.NopCloser(bytes.NewReader([]byte{}))

	}
	return r.Body
}

func request(url string) (io.ReadCloser, error) {
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	return wrapResponseBody(resp), err
}

func download(url string) (io.ReadCloser, string, error) {
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	return wrapResponseBody(resp), typeOfResponse(resp), err
}

func requestSync(url string) ([]byte, error) {
	r, err := request(url)
	defer r.Close()
	data, _ := ioutil.ReadAll(r)
	return data, err
}

func typeOfResponse(r *http.Response) string {
	if r == nil {
		return ""
	}
	contentType, contentDisposition := r.Header.Get("Content-Type"), r.Header.Get("Content-Disposition")
	if contentDisposition == "" {
		return extensionByType(contentType)
	}
	_, params, _ := mime.ParseMediaType(contentDisposition)
	filename := strings.Split(params["filename"], ".")
	return filename[len(filename)-1]
}

func extensionByType(contentType string) string {
	if contentType == "application/zip" {
		return ".zip"
	} else if contentType == "application/gzip" {
		return ".tar.gz"
	} else if contentType == "application/x-bzip" {
		return ".tar.bz"
	} else if contentType == "application/x-bzip2" {
		return ".tar.bz2"
	} else if contentType == "application/x-rar-compressed" {
		return ".rar"
	} else if contentType == "application/x-tar" {
		return ".tar"
	} else if contentType == "application/x-7z-compressed" {
		return ".7z"
	}
	return ""
}
