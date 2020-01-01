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

func getBroadcastLatestId(api string) ([]byte, error) {
	broadcastApi := api + "/broadcast"
	return requestSync(broadcastApi + "/latest/id")
}

func getBroadcastLatest(api string) ([]byte, error) {
	broadcastApi := api + "/broadcast"
	return requestSync(broadcastApi + "/latest")
}

func getBroadcastId(api string, id string) ([]byte, error) {
	broadcastApi := api + "/broadcast"
	return requestSync(broadcastApi + "/" + id)
}

func getVersion(api string) ([]byte, error) {
	brokerApi := api + "/broker"
	return requestSync(brokerApi + "/version")
}

func getDownloadSdkmanVersion(api string, versionType string) ([]byte, error) {
	brokerApi := api + "/broker"
	return requestSync(brokerApi + "/download/sdkman/version/" + versionType)
}

func getDownload(api string, sdk Sdk) (io.ReadCloser, error, string) {
	brokerApi := api + "/broker"
	return download(brokerApi + "/download/" + sdk.Candidate + "/" + sdk.Version + "/" + platform())
}

func getDefault(api string, candidate string) (string, error) {
	candidatesApi := api + "/candidates"

	res, err := requestSync(candidatesApi + "/default/" + candidate)
	return string(res), err
}

func getValidate(api string, sdk Sdk) (bool, error) {
	candidatesApi := api + "/candidates"
	url := fmt.Sprintf("%s/validate/%s/%s/%s", candidatesApi, sdk.Candidate, sdk.Version, platform())
	res, err := requestSync(url)
	return string(res) == "valid", err
}

func getList(api string) (io.ReadCloser, error) {
	candidatesApi := api + "/candidates"

	return request(candidatesApi + "/list")
}

func getVersionsList(api string, currentSdk Sdk, installed []Sdk) (io.ReadCloser, error) {
	candidatesApi := api + "/candidates"
	installedVers := ""
	for _, sdk := range installed {
		installedVers += sdk.Version + ","
	}
	url := fmt.Sprintf("%s/%s/%s/versions/list?current=%s&installed=%s", candidatesApi, currentSdk.Candidate, platform(), currentSdk.Version, installedVers)
	return request(url)
}

func getAll(api string) ([]string, error) {
	candidatesApi := api + "/candidates"
	res, err := requestSync(candidatesApi + "/all")
	return strings.Split(string(res), ","), err
}

func getVersionsAll(api string, candidate string) ([]byte, error) {
	candidatesApi := api + "/candidates"
	return requestSync(candidatesApi + "/" + candidate + "/" + platform() + "/versions/all")
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
	if r != nil {
		return r.Body
	} else {
		return ioutil.NopCloser(bytes.NewReader([]byte{}))
	}
}

func request(url string) (io.ReadCloser, error) {
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	return wrapResponseBody(resp), err
}

func download(url string) (io.ReadCloser, error, string) {
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	return wrapResponseBody(resp), err, typeOfResponse(resp)
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
	if contentDisposition != "" {
		_, params, _ := mime.ParseMediaType(contentDisposition)
		filename := strings.Split(params["filename"], ".")
		return filename[len(filename)-1]
	} else {
		return extensionByType(contentType)
	}
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
