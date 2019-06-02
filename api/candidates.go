package api

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var candidatesApi = e.Api + "/candidates"

func GetDefault(candidate string) []byte {
	return download(candidatesApi + "s/default/" + candidate)
}

func GetValidate(candidate string, version string) []byte {
	return download(candidatesApi + "/validate/" + candidate + "/" + version + "/" + e.Platform)
}

func GetAll() []byte {
	return download(candidatesApi + "/all")
}

func GetList() []byte {
	return download(candidatesApi + "/list")
}

func GetVersionsAll(candidate string) []byte {
	return download(candidatesApi + "/" + candidate + "/" + e.Platform + "/versions/all")
}

func GetVersionsList(candidate string, current string, installed []string) []byte {
	return download(candidatesApi + "/" + candidate + "/" + e.Platform + "/versions/list?current=" + current +
		"&installed=" + strings.Join(installed, ","))
}

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
