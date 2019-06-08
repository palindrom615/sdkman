package api

var broadcastApi = e.Api + "/broadcast"

func GetBroadcastLatestId() ([]byte, error) {
	return download(broadcastApi + "/latest/id")
}

func GetBroadcastLatest() ([]byte, error) {
	return download(broadcastApi + "/latest")
}

func GetBroadcastId(id string) ([]byte, error) {
	return download(broadcastApi + "/" + id)
}
