package serviceumkm

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Update_umkm(c *gin.Context) {
	type Data_umkm_kontainer struct {
		ID           int    `json:"id_umkm"`
		Nama         string `json:"nama_umkm"`
		Konten       string `json:"konten_umkm"`
		Kategori     int    `json:"kategori_umkm_id"`
		Foto         string `json:"foto_umkm"`
		NomorTelepon string `json:"no_telp_umkm"`
	}

	var input Data_umkm_kontainer
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid input"})
		return
	}

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}
	defer tx.Rollback(ctx)

	query := `
		UPDATE dev.umkm
		SET nama_umkm = $1, konten_umkm  = $2, kategori_umkm_id = $3 , foto_umkm = $4 , no_telp_umkm = $5
		WHERE id_umkm = $6
	`

	_, err = tx.Exec(ctx, query, input.Nama, input.Konten, input.Kategori, input.Foto, input.NomorTelepon, input.ID)
	if err != nil {
		fmt.Println("MASUK SINI")
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Berita updated successfully"})

}
