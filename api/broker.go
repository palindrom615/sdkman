package api

var brokerApi = e.Api + "/broker"

func GetVersion() []byte {
	return download(brokerApi + "/version")
}

func getDownloadSdkmanVersion(versionType string) []byte {
	return download(brokerApi + "/download/sdkman/version/" + versionType)
}

func getDownload(candidate string, version string) []byte {
	return download(brokerApi + "/download/" + candidate + "/" + version + "/" + e.Platform)
}
