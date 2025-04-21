package repository

import "github.com/odysseymorphey/SimpleAuth/internal/models"

type Repository interface {
	SaveRefreshToken(refreshToken *models.DBRecord) error
	UpdateRefreshToken(guid string, refreshToken *models.ComparableData) error
	GetDataForCompare(guid string) (*models.ComparableData, error)
	GetUserEmailMock(guid string) (string, error)
	Close()
}
