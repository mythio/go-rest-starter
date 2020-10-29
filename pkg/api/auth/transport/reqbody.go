package transport

// SignupReq represents incoming req body via HTTP
type SignupReq struct {
	FirstName       string `json:"first_name" binding:"required"`
	LastName        string `json:"last_name" binding:"required"`
	Username        string `json:"username" binding:"required,min=3,alphanum"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"password_confirm" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
}

// SigninReq represents incoming req body via HTTP
type SigninReq struct {
	Username string `json:"username" binding:"min=3,alphanum"`
	Password string `json:"password" binding:"required,min=8"`
	Email    string `json:"email" binding:"email"`
}
