package database

import (
	serviceberita "backendpgx7071/service/serviceBerita"
	serviceumkm "backendpgx7071/service/serviceUMKM"
	"log"
	"os"

	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewConnect() *pgxpool.Pool {
	databaseUrl := "postgres://postgres:123123@localhost:5432/morowali"

	config, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		log.Fatalf("Failed to parse database URL: %v", err)
		os.Exit(1)
	}

	config.MaxConns = 10

	db, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
		os.Exit(1)
	}

	log.Println("Koneksi ke database berhasil dibuat")

	serviceberita.InitiateDB(db)
	serviceumkm.InitiateDB(db)

	return db
}
