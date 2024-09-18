package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type DB struct {
    conn *pgx.Conn
}

func NewConnection() (*DB, error) {
    conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/postgres")
	if err != nil {
		return nil, err
	}

	return &DB{
        conn: conn,
    }, nil
}

func (d *DB) Close() {
    d.conn.Close(context.Background())
}