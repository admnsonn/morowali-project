package servicelayanandesa

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Input(c *gin.Context) {
	type Request struct {
		ID      int    `json:"id"`
		DesaID  int    `json:"desa_id"`
		Nama    string `json:"nama"`
		NomorHP string `json:"nomor_hp"`
		Alamat  string `json:"alamat"`
		Layanan string `json:"layanan"`
		Status  string `json:"status"`
		NIK     string `json:"nik"`
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

	var nikcek string

	cek_NIK := `
	select nik from dev.pengguna p where nik = $1
	`
	err = tx.QueryRow(ctx, cek_NIK, input.NIK).Scan(&nikcek)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	if nikcek == "" {
		c.JSON(http.StatusOK, gin.H{"status": false, "message": "NIK tidak terdaftar dalam desa"})
		return
	}

	cek_berita := `
	SELECT MAX(id) AS last_id
	FROM dev.layanan_sementara;	
	`
	err = tx.QueryRow(ctx, cek_berita).Scan(&id_now)

	if err != nil {
		id_now = 0
	}

	id_next = id_now + 1

	query := `
		INSERT INTO dev.layanan_sementara (
			id      ,
			desa_id ,
			nama    ,
			nomor_hp,
			alamat  ,
			layanan ,
			status  
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err = tx.Exec(ctx, query, id_next, input.DesaID, input.Nama, input.NomorHP, input.Alamat,
		input.Layanan, input.Status)
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
