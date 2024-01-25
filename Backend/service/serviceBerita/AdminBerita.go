package serviceberita

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func UpdateBerita(c *gin.Context) {
	type Data_berita_admin struct {
		ID         string `json:"id_berita"`
		Judul      string `json:"judul"`
		Deskripsi  string `json:"deskripsi"`
		FotoBerita string `json:"foto_berita"`
		Kategori   string `json:"kategori_id"`
		SubJudul   string `json:"sub_judul"`
	}

	var inputBerita Data_berita_admin
	if c.GetHeader("content-type") == "application/x-www-form-urlencoded" || c.GetHeader("content-type") == "application/x-www-form-urlencoded; charset=utf-8" {

		if err := c.Bind(&inputBerita); err != nil {
			return
		}

	} else {

		if err := c.BindJSON(&inputBerita); err != nil {
			return
		}

	}

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}
	defer tx.Rollback(ctx)

	query := `
		UPDATE dev.berita
		SET judul = $1, deskripsi  = $2, sub_judul = $3 , foto_berita = $4 , kategori_id = $5
		WHERE id_berita = $6
	`

	_, err = tx.Exec(ctx, query, inputBerita.Judul, inputBerita.Deskripsi, inputBerita.SubJudul, inputBerita.FotoBerita, inputBerita.Kategori, inputBerita.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Berita updated successfully"})

}
