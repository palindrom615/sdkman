package api

var broadcastApi = e.Api + "/broadcast"

func GetBroadcastLatestId() []byte {
	return download(broadcastApi + "/latest/id")
}

func GetBroadcastLatest() []byte {
	return download(broadcastApi + "/latest")
}

func GetBroadcastId(id string) []byte {
	return download(broadcastApi + "/" + id)
}
