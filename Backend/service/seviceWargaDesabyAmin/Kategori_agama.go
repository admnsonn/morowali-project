package sevicewargadesabyamin

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Kategori_agama(c *gin.Context) {
	type Kategori_agama_kontainer struct {
		Agama string `json:"Agama"`
		ID    int    `json:"id_agama"`
	}

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())

	}
	defer tx.Rollback(context.Background())

	query := `
	select id_agama , nama_agama  from dev.agama a 
	`

	row, err := tx.Query(ctx, query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	defer row.Close()

	var Tampung_agama []Kategori_agama_kontainer

	for row.Next() {
		var ambil Kategori_agama_kontainer
		err := row.Scan(

			&ambil.ID,
			&ambil.Agama,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			return
		}

		Tampung_agama = append(Tampung_agama, ambil)
	}

	if len(Tampung_agama) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Kategori Agama tersedia",
			"data":    Tampung_agama,
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return

	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Kategori Agama tidak ada",
			"data":    []Kategori_agama_kontainer{},
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}
}
