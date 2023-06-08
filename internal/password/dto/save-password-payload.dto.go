package password

type SavePasswordPayloadDto struct {
	Description string `json:"email,omitempty"`
	UserId      string `json:"userId"`
	Name        string `json:"name"`
	Platform    string `json:"platform,omitempty"`
	Pass        string `json:"pass"`
}
