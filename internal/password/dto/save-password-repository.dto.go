package password

type SavePasswordRepositoryDto struct {
	Description string `json:"email,omitempty"`
	UserId      string `json:"userId"`
	Name        string `json:"name"`
	Platform    string `json:"platform,omitempty"`
	Pass        string `json:"pass"`
	HashedKey   string
}
