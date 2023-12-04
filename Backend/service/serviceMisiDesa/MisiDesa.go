package servicemisidesa

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Misi_desa(c *gin.Context) {

	type Misi_desa_kontainer struct {
		MisiDesa string `json:"misi_desa"`
	}

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}

	query_misi_desa := `
		select a.misi_desa 
		from dev.misi_desa a, dev.desa d  
		where a.desa_id = d.id_desa 
		and d.id_desa = 1
	`

	row, err := tx.Query(ctx, query_misi_desa)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		err = tx.Commit(ctx)
		if err != nil {
			panic(err.Error())
		}
		return
	}

	defer row.Close()

	var Tampung_Misi []Misi_desa_kontainer

	for row.Next() {
		var ambil Misi_desa_kontainer

		err := row.Scan(
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

		Tampung_Misi = append(Tampung_Misi, ambil)
	}

	if len(Tampung_Misi) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Misi Desa Ada!",
			"data":    Tampung_Misi,
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Misi Desa Tidak Ada!",
			"data":    []Misi_desa_kontainer{},
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}
}
