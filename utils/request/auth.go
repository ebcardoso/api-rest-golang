package request

// api/v1/auth/signup
type SignupReq struct {
	Name                 string `json:"name,omitempty" validate:"required,min=5,max=50"`
	Email                string `json:"email,omitempty" validate:"required,email,min=6,max=60"`
	Password             string `json:"password,omitempty" validate:"required,min=6,max=20"`
	PasswordConfirmation string `json:"password_confirmation,omitempty" validate:"required,min=6,max=20"`
}

// api/v1/auth/signin
type SigninReq struct {
	Email    string `json:"email,omitempty" validate:"required,email,min=6,max=60"`
	Password string `json:"password,omitempty" validate:"required,min=6,max=20"`
}

// api/v1/auth/forgot_password_token
type ForgotPasswordReq struct {
	Email string `json:"email,omitempty" validate:"required,email,min=6,max=60"`
}

// api/v1/auth/reset_password_confirm
type ResetPasswordReq struct {
	Email                string `json:"email,omitempty" validate:"required,email,min=6,max=60"`
	Token                string `json:"token,omitempty" validate:"required"`
	Password             string `json:"password,omitempty" validate:"required,min=6,max=20"`
	PasswordConfirmation string `json:"password_confirmation,omitempty" validate:"required,min=6,max=20"`
}
