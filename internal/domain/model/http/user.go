package http

type UserRegisterRequest struct {
	UserId string `json:"user_id"`
	Passwd string `json:"password"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

type UserRegisterResponse struct {
	Status       int    `json:"status"`
	ErrorMessage string `json:"error_message"`
}
