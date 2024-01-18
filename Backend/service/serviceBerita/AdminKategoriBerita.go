package serviceberita

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Kategori_berita(c *gin.Context) {
	type Kategori_berita_kontainer struct {
		BeritaKategori string `json:"berita_kategori"`
		IDBerita       int    `json:"id_kategori_berita"`
	}

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(context.Background())

	query := `
		select 
			berita_kategori,
			id_kategori_berita
		from dev.kategori_berita kb 
	`

	row, err := tx.Query(ctx, query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	defer row.Close()

	var Tampung_Kategori_berita []Kategori_berita_kontainer

	for row.Next() {
		var ambil Kategori_berita_kontainer
		err := row.Scan(
			&ambil.BeritaKategori,
			&ambil.IDBerita,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			return
		}

		Tampung_Kategori_berita = append(Tampung_Kategori_berita, ambil)
	}

	if len(Tampung_Kategori_berita) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Kategori Berita tersedia",
			"data":    Tampung_Kategori_berita,
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return

	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Kategori Berita tidak ada",
			"data":    []Kategori_berita_kontainer{},
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}
}
