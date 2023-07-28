package usersdto

type UpdateUserRequest struct {
	Avatar   string `json:"image" form:"image"`
	Art      string `json:"art" form:"art"`
	Fullname string `json:"fullname" form:"fullname"`
	Greeting string `json:"greeting" form:"greeting"`
}
