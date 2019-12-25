package api

import (
	"bytes"
	"crypto/tls"
	"github.com/palindrom615/sdkman-cli/conf"
	"github.com/palindrom615/sdkman-cli/utils"
	"io"
	"io/ioutil"
	"net/http"
)

var (
	e      = conf.GetConf()
	client = &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: e.Insecure},
	}}
)

func request(url string) (io.ReadCloser, error) {
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if resp == nil {
		empty := ioutil.NopCloser(bytes.NewReader([]byte{}))
		return empty, err
	}
	return resp.Body, err
}

func download(url string) (io.ReadCloser, error, string) {
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if resp == nil {
		empty := ioutil.NopCloser(bytes.NewReader([]byte{}))
		return empty, err, ""
	}
	return resp.Body, err, utils.TypeOfResponse(resp.Header)
}

func requestSync(url string) ([]byte, error) {
	r, err := request(url)
	if r != nil {
		defer r.Close()
	}
	data, _ := ioutil.ReadAll(r)
	return data, err
}
