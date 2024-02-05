package serviceprodukhukum

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Tambah(c *gin.Context) {
	type PotensiDesa struct {
		Nama     string `json:"nama_produk"`
		Kategori int    `json:"id_kategori"`
		File     string `json:"file"`
		IDDesa   string `json:"id_desa"`
	}

	var input PotensiDesa

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
	SELECT MAX(id) AS last_id
	FROM dev.produk_hukum;
	`
	err = tx.QueryRow(ctx, cek_wisata).Scan(&id_now)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	id_next = id_now + 1

	fmt.Println("INI ID WISATA BARU", id_next)

	query := `
		INSERT INTO dev.produk_hukum (
			id, nama_produk, tanggal_terbit, desa_id, kategori_id, file
		)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err = tx.Exec(ctx, query, id_next, input.Nama, time.Now(), input.IDDesa, input.Kategori, input.File)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Error committing transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Data berhasil ditambahkan"})
}
