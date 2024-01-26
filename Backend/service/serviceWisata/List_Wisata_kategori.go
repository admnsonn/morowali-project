package servicewisata

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func List_kat_wisata(c *gin.Context) {
	type Kategori_wisata_kontainer struct {
		IDKategori     int    `json:"id_wisata_kategori"`
		WISATAkategori string `json:"wisata_kategori"`
	}

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(context.Background())

	query := `
		select idkategoriwisata, nama_kategori from dev.kategoriwisata
	`

	row, err := tx.Query(ctx, query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	defer row.Close()

	var Tampung_kategori_wisata []Kategori_wisata_kontainer

	for row.Next() {
		var ambil Kategori_wisata_kontainer
		err := row.Scan(
			&ambil.IDKategori,
			&ambil.WISATAkategori,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			return
		}

		Tampung_kategori_wisata = append(Tampung_kategori_wisata, ambil)
	}

	if len(Tampung_kategori_wisata) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Kategori WISATA tersedia",
			"data":    Tampung_kategori_wisata,
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return

	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Kategori WISATA tidak ada",
			"data":    []Kategori_wisata_kontainer{},
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}
}
