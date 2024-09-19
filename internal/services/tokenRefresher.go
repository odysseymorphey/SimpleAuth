package services

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/odysseymorphey/SimpleAuth/internal/models"
	"github.com/odysseymorphey/SimpleAuth/internal/postgres"
	"golang.org/x/crypto/bcrypt"
)

func RefreshAccessToken(db *postgres.DB, guid string, tokenPair *models.Pair) (string, error) {
	data, err := db.GetDataForCompare(guid)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(data.TokenHash), []byte(tokenPair.RefreshToken))
	if err != nil {
		return "", err
	}

	token, err := jwt.ParseWithClaims(tokenPair.AccessToken, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if token.Claims.(*models.CustomClaims).PairID != data.PairID {
		return "", err
	}



	return "", nil
}
