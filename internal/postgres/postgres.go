package postgres

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/odysseymorphey/SimpleAuth/internal/models"
)

type DB struct {
	conn *pgx.Conn
}

func NewConnection() (*DB, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		return nil, err
	}
	log.Println("Postgres connected to localhost:5432/postgres")

	return &DB{
		conn: conn,
	}, nil
}

func (d *DB) Close() {
	d.conn.Close(context.Background())
}

func (d *DB) SaveRefreshToken(refreshToken *models.DBRecord) error {
	query := `
        INSERT INTO refresh_tokens (guid, user_ip, hashed_token, pair_id)
        VALUES ($1, $2, $3, $4)
    `
	_, err := d.conn.Exec(context.Background(), query, refreshToken.GUID, refreshToken.UserIP, refreshToken.TokenHash, refreshToken.PairID)
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) UpdateRefreshToken(guid string, refreshToken *models.ComparableData) error {
	query := `
        UPDATE refresh_tokens
        SET hashed_token = $1, pair_id = $2 WHERE guid = $3
    `
	_, err := d.conn.Exec(context.Background(), query, refreshToken.TokenHash, refreshToken.PairID, guid)
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) GetDataForCompare(guid string) (*models.ComparableData, error) {
	query := `
        SELECT hashed_token, pair_id, user_ip FROM refresh_tokens WHERE guid = $1
    `
	row := d.conn.QueryRow(context.Background(), query, guid)

	var data models.ComparableData
	err := row.Scan(&data.TokenHash, &data.PairID, &data.UserIP)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (d *DB) GetUserEmailMock(guid string) (string, error) {
	return "gorillamango@gmail.com", nil
}
