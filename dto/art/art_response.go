package artdto

type DataResponse struct {
	ID    string `json:"id"`
	Image string `json:"image"`
}

type ArtsResponse struct {
	Arts interface{} `json:"arts"`
}
