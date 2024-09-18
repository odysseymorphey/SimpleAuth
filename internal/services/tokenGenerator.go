package services

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"time"

	"github.com/odysseymorphey/SimpleAuth/internal/models"

	jwt "github.com/golang-jwt/jwt/v5"
)

var (
	secretKey = []byte("Barbara_with_big_titties")
)

func generateAccessToken(userIP string, userAgent string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"ip": userIP,
		"userAgent": userAgent,
		"pairPass": pairID,  //TODO: Create a pairPass Generator
		"exp": jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
	})
	
    return token.SignedString(secretKey)
}

func generateRefreshToken() (string, error) {
	token := make([]byte, 32)
	if _, err := rand.Read(token); err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(token), nil
}

func GeneratePair(userIP string, userAgent string) (interface{}, error) {
	accessToken, err := generateAccessToken(userIP, userAgent)
	if err != nil {
		log.Println("Error generating access token: ", err)
		return "", err
	}

	refreshToken, err := generateRefreshToken()
	if err != nil {
		log.Println("Error generating refresh token: ", err)
		return "", err
	}

	pair := &models.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return pair, nil
}