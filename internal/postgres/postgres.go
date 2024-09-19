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

func (d *DB) SaveRefreshToken(*models.RefreshToken) error {
    return nil
}