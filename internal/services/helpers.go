package services

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/odysseymorphey/SimpleAuth/internal/models"
	"golang.org/x/crypto/bcrypt"
)

var (
	secretKey = []byte("Barbara_with_big_titties")
)

func GeneratePairID() (string, error) {
	t := make([]byte, 16)
	if _, err := rand.Read(t); err != nil {
		return "", err
	}

	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", t[0:4], t[4:6], t[6:8], t[8:10], t[10:])

	return uuid, nil
}

func generateBCrypt(token string) (string, error) {
	hashedToken, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedToken), nil
}

func generateAccessToken(userIP string, userAgent string, pairID string) (string, error) {
	claims := &models.CustomClaims{
		IP:        userIP,
		UserAgent: userAgent,
		PairID:    pairID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	return token.SignedString(secretKey)
}

func generateRefreshToken() (string, error) {
	token := make([]byte, 32)
	if _, err := rand.Read(token); err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(token), nil
}