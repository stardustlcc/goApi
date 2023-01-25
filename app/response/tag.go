package response

type TagInfoResponse struct {
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	State int    `json:"state"`
}

type TagListResponse struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}
