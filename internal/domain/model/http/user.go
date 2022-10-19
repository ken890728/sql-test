package http

type UserRegisterRequest struct {
	UserId string `json:"user_id"`
	Passwd string `json:"password"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

type UserRegisterResponse struct {
	Status       int    `json:"status"`
	ErrorMessage string `json:"error_message,omitempty"`
}

type UserLoginRequest struct {
	UserId string `json:"user_id"`
	Passwd string `json:"password"`
}

type UserLoginResponse struct {
	Status       int    `json:"status"`
	ErrorMessage string `json:"error_message,omitempty"`
	Token        string `json:"token,omitempty"`
}

type UserDataRequest struct {
}

type UserDataResponse struct {
	Status       int    `json:"status"`
	ErrorMessage string `json:"error_message,omitempty"`
	UserId       string `json:"user_id,omitempty"`
	Name         string `json:"name,omitempty"`
	Email        string `json:"email,omitempty"`
}

type UserAddBookRequest struct {
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type UserAddBookResponse struct {
	Status       int    `json:"status"`
	ErrorMessage string `json:"error_message,omitempty"`
}

type UserGetBooksRequest struct {
	Isbn   string `json:"isbn,omitempty"`
	Title  string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
}

type BookResponse struct {
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type UserGetBooksResponse struct {
	Status       int            `json:"status"`
	ErrorMessage string         `json:"error_message,omitempty"`
	Books        []BookResponse `json:"books"`
}
