package request

// api/v1/auth/signup
type SignupReq struct {
	Name                 string `json:"name,omitempty"`
	Email                string `json:"email,omitempty"`
	Password             string `json:"password,omitempty"`
	PasswordConfirmation string `json:"password_confirmation,omitempty"`
}

// api/v1/auth/signin
type SigninReq struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

// api/v1/auth/forgot_password_token
type ForgotPasswordReq struct {
	Email string `json:"email,omitempty"`
}

// api/v1/auth/reset_password_confirm
type ResetPasswordReq struct {
	Email                string `json:"email,omitempty"`
	Token                string `json:"token,omitempty"`
	Password             string `json:"password,omitempty"`
	PasswordConfirmation string `json:"password_confirmation,omitempty"`
}
