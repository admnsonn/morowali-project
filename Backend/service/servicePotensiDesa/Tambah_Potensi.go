package servicepotensidesa

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Tambah_potensi(c *gin.Context) {
	type PotensiDesa struct {
		Judul     string `json:"judul_potensi"`
		Deskripsi string `json:"deskripsi"`
		Foto      string `json:"foto_potensi_desa"`
		IDDesa    int    `json:"desa_id"`
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
	SELECT MAX(id_potensi) AS last_id
	FROM dev.potensi_desa;
	`
	err = tx.QueryRow(ctx, cek_wisata).Scan(&id_now)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	id_next = id_now + 1

	fmt.Println("INI ID WISATA BARU", id_next)

	query := `
		INSERT INTO dev.potensi_desa (
			id_potensi , judul_potensi , deskripsi , foto_potensi_desa , desa_id ,
			created_at,
			created_by,
			update_at,
			updated_by
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err = tx.Exec(ctx, query, id_next, input.Judul, input.Deskripsi, input.Foto, input.IDDesa,
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

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "POTENSI berhasil ditambahkan"})

}
