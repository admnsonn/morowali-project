package servicevisi

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Visi_desa(c *gin.Context) {

	type Visi_desa_kontainer struct {
		VisiDesa string `json:"visi_desa"`
	}

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}

	query_visi_desa := `
		select a.visi_desa 
		from dev.visi_desa a, dev.desa d  
		where a.desa_id = d.id_desa 
		and d.id_desa = 1
	`

	row, err := tx.Query(ctx, query_visi_desa)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		err = tx.Commit(ctx)
		if err != nil {
			panic(err.Error())
		}
		return
	}

	defer row.Close()

	var Tampung_Visi []Visi_desa_kontainer

	for row.Next() {
		var ambil Visi_desa_kontainer

		err := row.Scan(
			&ambil.VisiDesa,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			err = tx.Commit(ctx)
			if err != nil {
				panic(err.Error())
			}
			return
		}

		Tampung_Visi = append(Tampung_Visi, ambil)
	}

	if len(Tampung_Visi) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Visi Desa Ada!",
			"data":    Tampung_Visi,
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Visi Desa Tidak Ada!",
			"data":    []Visi_desa_kontainer{},
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}
}
