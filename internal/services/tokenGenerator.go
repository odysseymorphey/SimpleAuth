package services

import (
	"log"

	"github.com/odysseymorphey/SimpleAuth/internal/models"
	"github.com/odysseymorphey/SimpleAuth/internal/postgres"

)

func GeneratePair(db *postgres.DB, uInfo *models.UserInfo) (*models.Pair, error) {
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

	rToken := &models.DBRecord{
		GUID:      uInfo.GUID,
		UserIP:    uInfo.UserIP,
		TokenHash: hashedToken,
		PairID:    pairID,
	}
	
	db.SaveRefreshToken(rToken)
	
	pair := &models.Pair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return pair, nil
}
