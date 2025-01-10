package serializers

type UserLoginRequest struct {
	Username string `json:"username" validate:"required_without=Email,omitempty"`
	Email    string `json:"email" validate:"required_without=Username,omitempty,custom_email"`
	Password string `json:"password" validate:"required"`
}

type UserLoginResponse struct {
	Token    string      `json:"token"`
	UserData interface{} `json:"user_data"`
}
