package servicepotensidesa

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Update_Potensi(c *gin.Context) {
	type PotensiDesa struct {
		ID        int    `json:"id_potensi"`
		Judul     string `json:"judul_potensi"`
		Deskripsi string `json:"deskripsi"`
		Foto      string `json:"foto_potensi_desa"`
	}

	var input PotensiDesa

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
		UPDATE dev.potensi_desa
		SET judul_potensi = $1, deskripsi  = $2, foto_potensi_desa = $3 
		WHERE id_potensi = $4
	`

	_, err = tx.Exec(ctx, query, input.Judul, input.Deskripsi, input.Foto, input.ID)
	if err != nil {
		fmt.Println("MASUK SINI")
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "POTENSI updated successfully"})
}
