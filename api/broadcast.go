package api

func GetBroadcastLatestId(api string) ([]byte, error) {
	broadcastApi := api + "/broadcast"
	return requestSync(broadcastApi + "/latest/id")
}

func GetBroadcastLatest(api string) ([]byte, error) {
	broadcastApi := api + "/broadcast"
	return requestSync(broadcastApi + "/latest")
}

func GetBroadcastId(api string, id string) ([]byte, error) {
	broadcastApi := api + "/broadcast"
	return requestSync(broadcastApi + "/" + id)
}
