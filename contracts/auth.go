package contracts

type SigninRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupRequest struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	PasswordRepeat string `json:"password_repeat"`
}

type AuthResponse struct {
	Code    int
	Message string
	Token   string
}

type AuthErrorResponse struct {
	Code    int
	Message string
}
