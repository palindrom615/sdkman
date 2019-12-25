package api

var broadcastApi = e.Api + "/broadcast"

func GetBroadcastLatestId() ([]byte, error) {
	return requestSync(broadcastApi + "/latest/id")
}

func GetBroadcastLatest() ([]byte, error) {
	return requestSync(broadcastApi + "/latest")
}

func GetBroadcastId(id string) ([]byte, error) {
	return requestSync(broadcastApi + "/" + id)
}
