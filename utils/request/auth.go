package request

// api/v1/auth/signup
type SignupReq struct {
	Name                 string `json:"name,omitempty"`
	Email                string `json:"email,omitempty"`
	Password             string `json:"password,omitempty"`
	PasswordConfirmation string `json:"password_confirmation,omitempty"`
}
