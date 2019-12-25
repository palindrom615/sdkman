package api

import (
	"io"
	"strings"
)

var candidatesApi = e.Api + "/candidates"

func GetDefault(candidate string) (string, error) {
	res, err := requestSync(candidatesApi + "/default/" + candidate)
	return string(res), err
}

func GetValidate(candidate string, version string) (bool, error) {
	res, err := requestSync(candidatesApi + "/validate/" + candidate + "/" + version + "/" + e.Platform)
	return string(res) == "valid", err
}

func GetList() (io.ReadCloser, error) {
	return request(candidatesApi + "/list")
}

func GetVersionsList(candidate string, current string, installed []string) (io.ReadCloser, error) {
	return request(candidatesApi + "/" + candidate + "/" + e.Platform + "/versions/list?current=" + current +
		"&installed=" + strings.Join(installed, ","))
}

func GetAll() ([]string, error) {
	res, err := requestSync(candidatesApi + "/all")
	return strings.Split(string(res), ","), err
}

func GetVersionsAll(candidate string) ([]byte, error) {
	return requestSync(candidatesApi + "/" + candidate + "/" + e.Platform + "/versions/all")
}
