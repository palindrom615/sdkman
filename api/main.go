package api

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sdkman-cli/conf"
)

var (
	e      = conf.GetConf()
	client = &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: e.Insecure},
	}}
)

func download(url string) (io.ReadCloser, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, netErr := client.Do(req)
	if netErr != nil {
		resp.Body.Close()
	}
	return resp.Body, netErr
}

func downloadSync(url string) ([]byte, error) {
	body, err := download(url)
	defer body.Close()
	data, _ := ioutil.ReadAll(body)
	return data, err
}
