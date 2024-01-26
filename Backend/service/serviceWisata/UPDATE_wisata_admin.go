package servicewisata

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Update_Wisata(c *gin.Context) {
	type Wisata struct {
		ID           string `json:"id_wisata"`
		Nama         string `json:"nama_wisata"`
		Konten       string `json:"konten_wisata"`
		NomorTelepon string `json:"no_telp"`
		LinkLokasi   string `json:"link_lokasi"`
		Alamat       string `json:"alamat"`
		FotoWisata   string `json:"foto_wisata"`
		IDKategori   string `json:"idkategoriwisata"`
	}

	var input Wisata

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
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}
	defer tx.Rollback(ctx)

	query := `
		UPDATE dev.wisata
		SET nama_wisata = $1, konten_wisata  = $2, no_telp = $3 , link_lokasi = $4 , alamat = $5, foto_wisata = $6, idkategoriwisata = $7
		WHERE id_wisata = $8
	`

	_, err = tx.Exec(ctx, query, input.Nama, input.Konten, input.NomorTelepon, input.LinkLokasi, input.Alamat, input.FotoWisata, input.IDKategori, input.ID)
	if err != nil {
		fmt.Println("MASUK SINI")
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "WISATA updated successfully"})
}
