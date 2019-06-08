package api

import (
	"strings"
)

var candidatesApi = e.Api + "/candidates"

func GetDefault(candidate string) (string, error) {
	res, err := downloadSync(candidatesApi + "/default/" + candidate)
	return string(res), err
}

func GetValidate(candidate string, version string) (bool, error) {
	res, err := downloadSync(candidatesApi + "/validate/" + candidate + "/" + version + "/" + e.Platform)
	return string(res) == "valid", err
}

func GetAll() ([]string, error) {
	res, err := downloadSync(candidatesApi + "/all")
	return strings.Split(string(res), ","), err
}

func GetList() (string, error) {
	res, err := downloadSync(candidatesApi + "/list")
	return string(res), err
}

func GetVersionsAll(candidate string) ([]byte, error) {
	return downloadSync(candidatesApi + "/" + candidate + "/" + e.Platform + "/versions/all")
}

func GetVersionsList(candidate string, current string, installed []string) (string, error) {
	res, err := downloadSync(candidatesApi + "/" + candidate + "/" + e.Platform + "/versions/list?current=" + current +
		"&installed=" + strings.Join(installed, ","))
	return string(res), err
}
