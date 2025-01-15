package dto

type SignUpRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignInResponse struct {
	Token string `json:"token"`
}
