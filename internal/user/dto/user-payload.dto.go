package user

type UserPayloadDto struct {
	Email    string `json:"email"`
	Nickname string `json:"nickName"`
	Password string `json:"password,omitempty"`
}
