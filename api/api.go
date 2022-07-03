package api

import (
	"bytes"
	"fmt"
	"github.com/palindrom615/sdkman/util"
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

func GetDownload(api string, candidate string, version string) (io.ReadCloser, string, error) {
	brokerAPI := api + "/broker"
	downloadUrl := brokerAPI + "/download/" + candidate + "/" + version + "/" + util.Platform()
	println(candidate + "@" + version + ": download from " + downloadUrl)
	return download(downloadUrl)
}

func GetDefault(api string, candidate string) (string, error) {
	candidatesAPI := api + "/candidates"

	res, err := requestSync(candidatesAPI + "/default/" + candidate)
	return string(res), err
}

func GetValidate(api string, candidate string, version string) (bool, error) {
	candidatesAPI := api + "/candidates"
	url := fmt.Sprintf("%s/validate/%s/%s/%s", candidatesAPI, candidate, version, util.Platform())
	res, err := requestSync(url)
	return string(res) == "valid", err
}

func GetList(api string) (io.ReadCloser, error) {
	candidatesAPI := api + "/candidates"

	return request(candidatesAPI + "/list")
}

func GetVersionsList(api string, candidate string, version string, installed []string) (io.ReadCloser, error) {
	candidatesAPI := api + "/candidates"
	installedVersions := strings.Join(installed, ",")
	url := fmt.Sprintf("%s/%s/%s/versions/list?current=%s&installed=%s", candidatesAPI, candidate, util.Platform(), version, installedVersions)
	return request(url)
}

func GetAll(api string) ([]string, error) {
	candidatesAPI := api + "/candidates"
	res, err := requestSync(candidatesAPI + "/all")
	return strings.Split(string(res), ","), err
}

func getVersionsAll(api string, candidate string) ([]byte, error) {
	candidatesAPI := api + "/candidates"
	return requestSync(candidatesAPI + "/" + candidate + "/" + util.Platform() + "/versions/all")
}

func getAlive(api string) ([]byte, error) {
	return requestSync(api + "/alive")
}

func getSelfupdate(api string, beta bool) ([]byte, error) {
	return requestSync(api + "/selfupdate?beta=" + strconv.FormatBool(beta))
}

func getHooks(api string, phase string, candidate string, version string) ([]byte, error) {
	return requestSync(api + "/hooks/" + phase + "/" + candidate + "/" + version + "/" + util.Platform())
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
	ext := filename[len(filename)-1]
	if ext == "gz" || ext == "bz2" || ext == "xz" || ext == "lz4" || ext == "sz" {
		ext = filename[len(filename)-2] + "." + ext
	}
	return ext
}

func extensionByType(contentType string) string {
	if contentType == "application/zip" {
		return "zip"
	} else if contentType == "application/gzip" {
		return "tar.gz"
	} else if contentType == "application/x-bzip" {
		return "tar.bz"
	} else if contentType == "application/x-bzip2" {
		return "tar.bz2"
	} else if contentType == "application/x-rar-compressed" {
		return "rar"
	} else if contentType == "application/x-tar" {
		return "tar"
	} else if contentType == "application/x-7z-compressed" {
		return "7z"
	}
	return ""
}
