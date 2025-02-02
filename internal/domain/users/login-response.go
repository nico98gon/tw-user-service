package users

type LoginResponse struct {
	Token string `json:"token, omitempty"`
}