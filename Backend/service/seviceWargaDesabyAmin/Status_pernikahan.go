package sevicewargadesabyamin

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Pernikahan(c *gin.Context) {
	type Kategori_Pernikahan_kontainer struct {
		ID                int    `json:"id"`
		Status_perkawinan string `json:"pernikahan"`
	}

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(context.Background())

	query := `
		select status_perkawinan , id_status  from dev.status_perkawinan
	`

	row, err := tx.Query(ctx, query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	defer row.Close()

	var Tampung_pernikahan []Kategori_Pernikahan_kontainer

	for row.Next() {
		var ambil Kategori_Pernikahan_kontainer
		err := row.Scan(
			&ambil.Status_perkawinan,
			&ambil.ID,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			return
		}

		Tampung_pernikahan = append(Tampung_pernikahan, ambil)
	}

	if len(Tampung_pernikahan) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Kategori Pernikahan tersedia",
			"data":    Tampung_pernikahan,
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return

	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Kategori Pernikahan tidak ada",
			"data":    []Kategori_Pernikahan_kontainer{},
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}
}
