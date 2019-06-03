package api

import (
	"strings"
)

var candidatesApi = e.Api + "/candidates"

func GetDefault(candidate string) []byte {
	return download(candidatesApi + "/default/" + candidate)
}

func GetValidate(candidate string, version string) bool {
	return string(download(candidatesApi+"/validate/"+candidate+"/"+version+"/"+e.Platform)) == "valid"
}

func GetAll() []string {
	return strings.Split(string(download(candidatesApi+"/all")), ",")
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
