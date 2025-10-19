package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Connect() *pgx.Conn {

	pglink := "postgres://a4bhi:a4bhi@localhost:5432/slothcrypt"
	conn, _ := pgx.Connect(context.Background(), pglink)
	return conn

}
