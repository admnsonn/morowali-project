package servicepegawai

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Tambah(c *gin.Context) {
	type Request struct {
		Nama    string `json:"nama"`
		Jabatan string `json:"jabatan"`
		Foto    string `json:"foto_potensi_desa"`
		IDDesa  int    `json:"desa_id"`
	}

	var input Request

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
	SELECT MAX(id_pegawai) AS last_id
	FROM dev.pegawai;
	`
	err = tx.QueryRow(ctx, cek_wisata).Scan(&id_now)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	id_next = id_now + 1

	fmt.Println("INI ID: ", id_next)

	query := `
		INSERT INTO dev.pegawai (
			id_pegawai , nama, jabatan, foto , id_desa
		)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err = tx.Exec(ctx, query, id_next, input.Nama, input.Jabatan, input.Foto, input.IDDesa)
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
