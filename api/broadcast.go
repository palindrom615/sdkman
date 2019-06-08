package api

var broadcastApi = e.Api + "/broadcast"

func GetBroadcastLatestId() ([]byte, error) {
	return downloadSync(broadcastApi + "/latest/id")
}

func GetBroadcastLatest() ([]byte, error) {
	return downloadSync(broadcastApi + "/latest")
}

func GetBroadcastId(id string) ([]byte, error) {
	return downloadSync(broadcastApi + "/" + id)
}
