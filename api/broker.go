package api

var brokerApi = e.Api + "/broker"

func GetVersion() ([]byte, error) {
	return download(brokerApi + "/version")
}

func getDownloadSdkmanVersion(versionType string) ([]byte, error) {
	return download(brokerApi + "/download/sdkman/version/" + versionType)
}

func getDownload(candidate string, version string) ([]byte, error) {
	return download(brokerApi + "/download/" + candidate + "/" + version + "/" + e.Platform)
}
