package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

type DB struct {
    conn *pgx.Conn
}

func NewConnection() (*DB, error) {
    conn, err := pgx.Connect(context.Background(), "postgres://postgres:mysecretpassword@localhost:5432/postgres")
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

func (d *DB) SaveRefreshToken(token string) error {
    _, err := d.conn.Exec(context.Background(), "INSERT INTO refresh_tokens (token) VALUES ($1)", token)
    if err != nil {
        return err
    }
	
    return nil
}