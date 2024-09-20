package services

import (
	"log"

	"github.com/golang-jwt/jwt/v5"
	"github.com/odysseymorphey/SimpleAuth/internal/models"
	"github.com/odysseymorphey/SimpleAuth/internal/postgres"
	"golang.org/x/crypto/bcrypt"
)

func RefreshAccessToken(db *postgres.DB, uInfo *models.UserInfo , tokenPair *models.Pair) (*models.Pair, error) {
	data, err := db.GetDataForCompare(uInfo.GUID)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(data.TokenHash), []byte(tokenPair.RefreshToken))
	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseWithClaims(tokenPair.AccessToken, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if token.Claims.(*models.CustomClaims).PairID != data.PairID {
		return nil, err
	}

	pairID, err := GeneratePairID()
	if err != nil {
		log.Println("Error generating pair ID: ", err)
		return nil, err
	}

	accessToken, err := generateAccessToken(uInfo.UserIP, uInfo.UserAgent, pairID)
	if err != nil {
		log.Println("Error generating access token: ", err)
		return nil, err
	}

	refreshToken, err := generateRefreshToken()
	if err != nil {
		log.Println("Error generating refresh token: ", err)
		return nil, err
	}

	hashedToken, err := generateBCrypt(refreshToken)
	if err != nil {
		log.Println("Error hashing refresh token: ", err)
		return nil, err
	}

	data = &models.ComparableData{
		TokenHash: hashedToken,
		PairID:    data.PairID,
	}

	err = db.UpdateRefreshToken(uInfo.GUID, data)
	if err != nil {
		return nil, err
	}

	pair := &models.Pair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return pair, nil
}
