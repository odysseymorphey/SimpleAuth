package models

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshToken struct {
	GUID   string
	UserIP string
	Hash   string
	PairID string
}
