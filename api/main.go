package api

import (
	"bytes"
	"crypto/tls"
	"io"
	"io/ioutil"
	"net/http"
	"sdkman-cli/conf"
	"sdkman-cli/utils"
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
