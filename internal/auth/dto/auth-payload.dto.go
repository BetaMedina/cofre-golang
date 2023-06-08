package auth

type AuthPayloadDto struct {
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}
