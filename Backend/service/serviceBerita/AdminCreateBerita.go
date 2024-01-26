package serviceberita

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Tulis_Berita(c *gin.Context) {
	type Request struct {
		Judul      string `json:"judul"`
		SubJudul   string `json:"sub_judul"`
		Deskripsi  string `json:"deskripsi"`
		FotoBerita string `json:"foto_berita"`
		DesaID     string `json:"desa_id"`
		KategoriID int    `json:"kategori_id"`
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

	cek_berita := `
	SELECT MAX(id_berita) AS last_id
	FROM dev.berita;	
	`
	err = tx.QueryRow(ctx, cek_berita).Scan(&id_now)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	id_next = id_now + 1

	query := `
		INSERT INTO dev.berita (
			id_berita,
			judul, 
			sub_judul, 
			deskripsi, 
			foto_berita, 
			kategori_id, 
			desa_id,
			created_at,
			created_by,
			update_at,
			updated_by
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`

	_, err = tx.Exec(ctx, query, id_next, input.Judul, input.SubJudul, input.Deskripsi, input.FotoBerita, input.KategoriID, input.DesaID,
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

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Berita berhasil ditambahkan"})

}
