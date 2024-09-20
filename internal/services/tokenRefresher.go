package services

import (
	"fmt"
	"log"

	"github.com/odysseymorphey/SimpleAuth/internal/models"
	"github.com/odysseymorphey/SimpleAuth/internal/postgres"
	"golang.org/x/crypto/bcrypt"
)

func RefreshAccessToken(db *postgres.DB, uInfo *models.UserInfo , tokenPair *models.Pair) (*models.Pair, error) {
	userData, err := db.GetDataForCompare(uInfo.GUID)
	if err != nil {
		log.Println("Error getting data for compare: ", err)
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.TokenHash), []byte(tokenPair.RefreshToken))
	if err != nil {
		log.Println("Error comparing hash: ", err)
		return nil, err
	}

	claims, err := parseToken(tokenPair.AccessToken)
	if err != nil {
		log.Println("Error parsing token: ", err)
		return nil, err
	}

	if claims.PairID != userData.PairID {
		return nil, fmt.Errorf("error: wrong pair ID")
	}

	if claims.IP != userData.UserIP {
		sendEmailNotification();
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

	data := &models.ComparableData{
		TokenHash: hashedToken,
		PairID:    pairID,
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
