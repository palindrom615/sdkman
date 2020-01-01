package api

import (
	"fmt"
	"github.com/palindrom615/sdk/local"
	"github.com/palindrom615/sdk/utils"
	"io"
	"strings"
)

func GetDefault(api string, candidate string) (string, error) {
	candidatesApi := api + "/candidates"

	res, err := requestSync(candidatesApi + "/default/" + candidate)
	return string(res), err
}

func GetValidate(api string, sdk local.Sdk) (bool, error) {
	candidatesApi := api + "/candidates"
	url := fmt.Sprintf("%s/validate/%s/%s/%s", candidatesApi, sdk.Candidate, sdk.Version, utils.Platform())
	res, err := requestSync(url)
	return string(res) == "valid", err
}

func GetList(api string) (io.ReadCloser, error) {
	candidatesApi := api + "/candidates"

	return request(candidatesApi + "/list")
}

func GetVersionsList(api string, candidate string, current string, installed []string) (io.ReadCloser, error) {
	candidatesApi := api + "/candidates"
	return request(candidatesApi + "/" + candidate + "/" + utils.Platform() + "/versions/list?current=" + current +
		"&installed=" + strings.Join(installed, ","))
}

func GetAll(api string) ([]string, error) {
	candidatesApi := api + "/candidates"
	res, err := requestSync(candidatesApi + "/all")
	return strings.Split(string(res), ","), err
}

func GetVersionsAll(api string, candidate string) ([]byte, error) {
	candidatesApi := api + "/candidates"
	return requestSync(candidatesApi + "/" + candidate + "/" + utils.Platform() + "/versions/all")
}
