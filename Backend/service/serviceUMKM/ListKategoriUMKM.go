package serviceumkm

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func List_kat_umkm(c *gin.Context) {
	type Kategori_UMKM_kontainer struct {
		UMKMKategori string `json:"umkm_kategori"`
		IDKategori   int    `json:"id_kategori_umkm"`
	}

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(context.Background())

	query := `
		select * from dev.kategori_umkm ku 
	`

	row, err := tx.Query(ctx, query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	defer row.Close()

	var Tampung_Kategori_UMKM []Kategori_UMKM_kontainer

	for row.Next() {
		var ambil Kategori_UMKM_kontainer
		err := row.Scan(
			&ambil.IDKategori,
			&ambil.UMKMKategori,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			return
		}

		Tampung_Kategori_UMKM = append(Tampung_Kategori_UMKM, ambil)
	}

	if len(Tampung_Kategori_UMKM) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Kategori UMKM tersedia",
			"data":    Tampung_Kategori_UMKM,
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return

	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Kategori UMKM tidak ada",
			"data":    []Kategori_UMKM_kontainer{},
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}
}
