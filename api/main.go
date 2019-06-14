package api

import (
	"bytes"
	"crypto/tls"
	"github.oom/palindrom615/sdkman-cli/conf"
	"github.oom/palindrom615/sdkman-cli/utils"
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

func download(url string) (io.ReadCloser, error) {
	req, err := http.NewRequest("GET", url, nil)
	utils.Check(err)

	resp, err := client.Do(req)
	if resp == nil {
		empty := ioutil.NopCloser(bytes.NewReader([]byte{}))
		return empty, err
	}
	return resp.Body, err
}

func downloadSync(url string) ([]byte, error) {
	r, err := download(url)
	defer r.Close()
	data, _ := ioutil.ReadAll(r)
	return data, err
}
