package api

import (
	"crypto/tls"
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

func download(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	return client.Do(req)
}

func downloadSync(url string) ([]byte, error) {
	resp, err := download(url)
	if resp == nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	return data, err
}
