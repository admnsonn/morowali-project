package servicepemerintahgo

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Visi_Misi(c *gin.Context) {

	type Visi_desa_kontainer struct {
		VisiDesa string `json:"visi_desa"`
		MisiDesa string `json:"misi_desa"`
	}

	id := c.Param("id")

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(ctx)

	query_visi_desa := `
		select visi, misi  from dev.desa where id_desa = $1
	`

	row, err := tx.Query(ctx, query_visi_desa, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		err = tx.Commit(ctx)
		if err != nil {
			panic(err.Error())
		}
		return
	}

	defer row.Close()

	var Tampung_Visi_Misi []Visi_desa_kontainer

	for row.Next() {
		var ambil Visi_desa_kontainer

		err := row.Scan(
			&ambil.VisiDesa,
			&ambil.MisiDesa,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			err = tx.Commit(ctx)
			if err != nil {
				panic(err.Error())
			}
			return
		}

		Tampung_Visi_Misi = append(Tampung_Visi_Misi, ambil)
	}

	if len(Tampung_Visi_Misi) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Visi Misi Desa Ada!",
			"data":    Tampung_Visi_Misi,
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Visi Misi Desa Tidak Ada!",
			"data":    []Visi_desa_kontainer{},
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}
}
