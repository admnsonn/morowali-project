package database

import (
	"log"
	servicepemerintahgo "main/service/ServicePemerintah.go"
	serviceproduksi "main/service/ServiceProduksi"
	serviceberita "main/service/serviceBerita"
	servicekreativitas "main/service/serviceKreativitas"
	servicepegawai "main/service/servicePegawai"
	serviceperaturandesa "main/service/servicePeraturanDesa"
	servicepotensidesa "main/service/servicePotensiDesa"
	serviceprodukhukum "main/service/serviceProdukHukum"
	serviceumkm "main/service/serviceUMKM"
	serviceumum "main/service/serviceUmum"
	servicewilayahdesa "main/service/serviceWilayahDesa"
	servicewisata "main/service/serviceWisata"
	"main/service/servicelogin"
	sevicewargadesabyamin "main/service/seviceWargaDesabyAmin"
	"os"

	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewConnect() *pgxpool.Pool {
	// databaseUrl := "postgres://postgres:123123@localhost:5432/morowali"
	databaseUrl := "postgres://postgres:boyang123@morodb.cmwu6s1vldt3.ap-southeast-1.rds.amazonaws.com:5432/morowali"

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
	servicewisata.InitiateDB(db)
	servicepotensidesa.InitiateDB(db)
	servicekreativitas.InitiateDB(db)
	servicewilayahdesa.InitiateDB(db)
	servicelogin.InitiateDB(db)
	serviceperaturandesa.InitiateDB(db)
	serviceumum.InitiateDB(db)
	sevicewargadesabyamin.InitiateDB(db)
	serviceproduksi.InitiateDB(db)
	servicepemerintahgo.InitiateDB(db)
	servicepegawai.InitiateDB(db)
	serviceprodukhukum.InitiateDB(db)

	return db
}

// KONEKSI DATA BASE
