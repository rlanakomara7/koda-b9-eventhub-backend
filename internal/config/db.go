package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// kumpulan koneksi database yg berisi
func InitDB() (*pgxpool.Pool, error) {
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbHost := os.Getenv("DBHOST")
	dbPort := os.Getenv("DBPORT")
	dbName := os.Getenv("DBNAME")
	// Tambahkan ?sslmode=disable di akhir connection string
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)
	return pgxpool.New(context.Background(), connString)
}

func PingDB(db *pgxpool.Pool) error {
	return db.Ping(context.Background())
}
