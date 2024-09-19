package services

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"time"

	"github.com/odysseymorphey/SimpleAuth/internal/models"
	"golang.org/x/crypto/bcrypt"

	jwt "github.com/golang-jwt/jwt/v5"
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
    token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"ip": userIP,
		"userAgent": userAgent,
		"pairID": pairID,
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

func GeneratePair(userID string, userIP string, userAgent string) (*models.Tokens, *models.RefreshToken, error) {
	pairID, err := GeneratePairID()
	if err != nil {
		log.Println("Error generating pair ID: ", err)
		return nil, nil, err
	}

	accessToken, err := generateAccessToken(userIP, userAgent, pairID)
	if err != nil {
		log.Println("Error generating access token: ", err)
		return nil, nil, err
	}

	refreshToken, err := generateRefreshToken()
	if err != nil {
		log.Println("Error generating refresh token: ", err)
		return nil, nil, err
	}

	pair := &models.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	hashedToken, err := generateBCrypt(refreshToken)
	if err != nil {
		log.Println("Error hashing refresh token: ", err)
		return nil, nil, err
	}

	rToken := &models.RefreshToken{
		GUID: userID,
		UserIP: userIP,
		Hash: hashedToken,
		PairID: pairID,
	}

	return pair, rToken, nil
}