package api

import (
	"bytes"
	"crypto/tls"
	"github.com/palindrom615/sdk/conf"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"strings"
)

var (
	e      = conf.GetConf()
	client = &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: e.Insecure},
	}}
)

func wrapResponseBody(r *http.Response) io.ReadCloser {
	if r != nil {
		return r.Body
	} else {
		return ioutil.NopCloser(bytes.NewReader([]byte{}))
	}
}

func request(url string) (io.ReadCloser, error) {
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	return wrapResponseBody(resp), err
}

func download(url string) (io.ReadCloser, error, string) {
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
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
		exts, _ := mime.ExtensionsByType(contentType)
		return exts[0]
	}
}
