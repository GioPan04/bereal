package models

type BeRealMemory struct {
	ID        string          `json:"id"`
	Thumbnail *BeRealPhoto    `json:"thumbnail"`
	Primary   *BeRealPhoto    `json:"primary"`
	Secondary *BeRealPhoto    `json:"secondary"`
	IsLate    bool            `json:"isLate"`
	MemoryDay string          `json:"memoryDay"`
	Location  *BeRealLocation `json:"location"`
}

type BeRealMemoriesEndpoint struct {
	Prev                 string          `json:"prev"`
	Next                 string          `json:"next"`
	Data                 []*BeRealMemory `json:"data"`
	MemoriesSynchronized bool            `json:"memoriesSynchronized"`
}
