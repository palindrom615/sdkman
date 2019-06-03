package api

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"sdkman-cli/conf"
)

var e = conf.GetConf()

func download(url string) []byte {
	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: e.Insecure},
	}}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	return data
}
