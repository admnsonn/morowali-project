package serviceperaturandesa

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Peraturan_desa(c *gin.Context) {

	type Peraturan_desa_kontainer struct {
		PeraturanDesa string `json:"peraturan_desa"`
	}

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}

	query_peraturan_desa := `
		select a.peraturan_desa 
		from dev.peraturan_desa a, dev.desa d  
		where a.desa_id = d.id_desa 
		and d.id_desa = 1
	`

	row, err := tx.Query(ctx, query_peraturan_desa)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		err = tx.Commit(ctx)
		if err != nil {
			panic(err.Error())
		}
		return
	}

	defer row.Close()

	var Tampung_Peraturan []Peraturan_desa_kontainer

	for row.Next() {
		var ambil Peraturan_desa_kontainer

		err := row.Scan(
			&ambil.PeraturanDesa,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			err = tx.Commit(ctx)
			if err != nil {
				panic(err.Error())
			}
			return
		}

		Tampung_Peraturan = append(Tampung_Peraturan, ambil)
	}

	if len(Tampung_Peraturan) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Peraturan Desa Ada!",
			"data":    Tampung_Peraturan,
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Peraturan Desa Tidak Ada!",
			"data":    []Peraturan_desa_kontainer{},
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}
}
