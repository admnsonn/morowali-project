package sevicewargadesabyamin

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Kategori_financial(c *gin.Context) {
	type Kategori_Financial_kontainer struct {
		Financial string `json:"financial"`
		ID        int    `json:"id_financial"`
	}

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(context.Background())

	query := `
		select kategori_financial , id_kategori_financial  from dev.kategori_financial
	`

	row, err := tx.Query(ctx, query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	defer row.Close()

	var Tampung_Financial []Kategori_Financial_kontainer

	for row.Next() {
		var ambil Kategori_Financial_kontainer
		err := row.Scan(
			&ambil.Financial,
			&ambil.ID,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			return
		}

		Tampung_Financial = append(Tampung_Financial, ambil)
	}

	if len(Tampung_Financial) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Kategori Financial tersedia",
			"data":    Tampung_Financial,
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return

	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Kategori Financial tidak ada",
			"data":    []Kategori_Financial_kontainer{},
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}
}
