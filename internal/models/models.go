package models

import "github.com/golang-jwt/jwt/v5"

type UserInfo struct {
	GUID    string
	UserIP    string
	UserAgent string
}

type Pair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type CustomClaims struct {
	IP        string `json:"ip"`
	UserAgent string `json:"userAgent"`
	PairID    string `json:"pairID"`
	jwt.RegisteredClaims
}

type DBRecord struct {
	GUID      string
	UserIP    string
	TokenHash string
	PairID    string
}

type ComparableData struct {
	TokenHash string
	PairID    string
}
