package servicepemerintahgo

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Sambutan(c *gin.Context) {
	type Visi_desa_kontainer struct {
		Sambutan_kepdes string `json:"sambutan"`
	}

	id := c.Param("id")

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(ctx)

	query := `
		select sambutan_kepdes from dev.desa where id_desa = $1
	`

	row, err := tx.Query(ctx, query, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		err = tx.Commit(ctx)
		if err != nil {
			panic(err.Error())
		}
		return
	}

	defer row.Close()

	var Tampung_Sambutan []Visi_desa_kontainer

	for row.Next() {
		var ambil Visi_desa_kontainer

		err := row.Scan(
			&ambil.Sambutan_kepdes,
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
			"message": "Data Ada!",
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
			"message": "Data Tidak Ada!",
			"data":    []Visi_desa_kontainer{},
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}
}
