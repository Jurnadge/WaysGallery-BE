package authdto

type AuthRequest struct {
	Email    string `json:"email" validate:"required" form:"email"`
	Password string `json:"password" validate:"required" form:"password"`
	Fullname string `json:"fullname" validate:"required" form:"fullname"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required" form:"email"`
	Password string `json:"password" validate:"required" form:"password"`
}

type LoginResponse struct {
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Password string `json:"-"`
	Token    string `json:"token"`
}
