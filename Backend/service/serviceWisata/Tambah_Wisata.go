package servicewisata

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Tambah_wisata(c *gin.Context) {

	type Wisata struct {
		Nama         string `json:"nama_wisata"`
		Konten       string `json:"konten_wisata"`
		NomorTelepon string `json:"no_telp"`
		DesaID       int    `json:"desa_id"`
		LinkLokasi   string `json:"link_lokasi"`
		Alamat       string `json:"alamat"`
		FotoWisata   string `json:"foto_wisata"`
		IDKategori   int    `json:"idkategoriwisata"`
	}

	var input Wisata

	if c.GetHeader("content-type") == "application/x-www-form-urlencoded" || c.GetHeader("content-type") == "application/x-www-form-urlencoded; charset=utf-8" {

		if err := c.Bind(&input); err != nil {
			return
		}

	} else {

		if err := c.BindJSON(&input); err != nil {
			return
		}

	}

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(context.Background())

	var id_next, id_now int

	cek_wisata := `
	SELECT MAX(id_wisata) AS last_id
	FROM dev.wisata;
	`
	err = tx.QueryRow(ctx, cek_wisata).Scan(&id_now)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	id_next = id_now + 1

	fmt.Println("INI ID WISATA BARU", id_next)

	query := `
		INSERT INTO dev.wisata (
			id_wisata , nama_wisata , konten_wisata , no_telp , desa_id , link_lokasi , alamat , foto_wisata, idkategoriwisata,
			created_at,
			created_by,
			update_at,
			updated_by
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`

	_, err = tx.Exec(ctx, query, id_next, input.Nama, input.Konten, input.NomorTelepon, input.DesaID, input.LinkLokasi, input.Alamat, input.FotoWisata, input.IDKategori,
		time.Now(), "admin", time.Now(), "admin")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Error committing transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "WISATA berhasil ditambahkan"})

}
