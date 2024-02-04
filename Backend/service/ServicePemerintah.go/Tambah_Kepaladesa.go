package servicepemerintahgo

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Tambah_kepalaDesa(c *gin.Context) {
	type Kepala_Desa struct {
		Nama   string `json:"nama"`
		Foto   string `json:"foto_kepala_desa"`
		Start  string `json:"jabatan_dimulai"`
		End    string `json:"jabatan_berakhir"`
		DesaID string `json:"desa_id"`
	}

	var input Kepala_Desa

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

	cek_berita := `
	SELECT MAX(id_kepala_desa) AS last_id
	FROM dev.kepala_desa;
	`
	err = tx.QueryRow(ctx, cek_berita).Scan(&id_now)

	if err != nil {
		id_next = id_now + 1
	} else {
		id_next = id_now + 1
	}

	query := `
		INSERT INTO dev.kepala_desa (
			id_kepala_desa , nama , foto_kepala_desa , jabatan_dimulai , jabatan_berakhir , desa_id
		)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err = tx.Exec(ctx, query, id_next, input.Nama, input.Foto, input.Start, input.End, input.DesaID)
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
