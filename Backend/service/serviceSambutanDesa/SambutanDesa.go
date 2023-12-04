package servicesambutandesa

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Sambutan_desa(c *gin.Context) {

	type Sambutan_desa_kontainer struct {
		SambutanDesa string `json:"sambutan_desa"`
	}

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}

	query_sambutan_desa := `
		select a.sambutan_desa 
		from dev.sambutan_desa a, dev.desa d  
		where a.desa_id = d.id_desa 
		and d.id_desa = 1
	`

	row, err := tx.Query(ctx, query_sambutan_desa)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		err = tx.Commit(ctx)
		if err != nil {
			panic(err.Error())
		}
		return
	}

	defer row.Close()

	var Tampung_Sambutan []Sambutan_desa_kontainer

	for row.Next() {
		var ambil Sambutan_desa_kontainer

		err := row.Scan(
			&ambil.SambutanDesa,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			err = tx.Commit(ctx)
			if err != nil {
				panic(err.Error())
			}
			return
		}

		Tampung_Sambutan = append(Tampung_Sambutan, ambil)
	}

	if len(Tampung_Sambutan) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Sambutan Desa Ada!",
			"data":    Tampung_Sambutan,
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Sambutan Desa Tidak Ada!",
			"data":    []Sambutan_desa_kontainer{},
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}
}
