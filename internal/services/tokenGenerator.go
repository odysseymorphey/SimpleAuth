package services

import (
	"log"

	"github.com/odysseymorphey/SimpleAuth/internal/models"
	"github.com/odysseymorphey/SimpleAuth/internal/postgres"

)

func GeneratePair(db *postgres.DB, userID string, userIP string, userAgent string) (*models.Pair, error) {
	pairID, err := GeneratePairID()
	if err != nil {
		log.Println("Error generating pair ID: ", err)
		return nil, err
	}

	accessToken, err := generateAccessToken(userIP, userAgent, pairID)
	if err != nil {
		log.Println("Error generating access token: ", err)
		return nil, err
	}

	refreshToken, err := generateRefreshToken()
	if err != nil {
		log.Println("Error generating refresh token: ", err)
		return nil, err
	}

	pair := &models.Pair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	hashedToken, err := generateBCrypt(refreshToken)
	if err != nil {
		log.Println("Error hashing refresh token: ", err)
		return nil, err
	}

	rToken := &models.DBRecord{
		GUID:      userID,
		UserIP:    userIP,
		TokenHash: hashedToken,
		PairID:    pairID,
	}
	
	db.SaveRefreshToken(rToken)

	return pair, nil
}
