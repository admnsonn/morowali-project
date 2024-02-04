package servicepemerintahgo

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Setting_visimisi(c *gin.Context) {
	type Visi_desa_kontainer struct {
		VisiDesa string `json:"visi_desa"`
		MisiDesa string `json:"misi_desa"`
		ID       string `json:"id_desa"`
	}

	var input Visi_desa_kontainer

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
	defer tx.Rollback(ctx)

	query := `
			UPDATE dev.desa
			SET visi = $1, misi  = $2
			WHERE id_desa = $3
		`

	_, err = tx.Exec(ctx, query, input.VisiDesa, input.MisiDesa, input.ID)
	if err != nil {
		fmt.Println("MASUK SINI")
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Visi Misi updated successfully"})

}
