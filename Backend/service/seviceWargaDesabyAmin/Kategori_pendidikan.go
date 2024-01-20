package sevicewargadesabyamin

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Kategori_pendidikan(c *gin.Context) {
	type Kategori_pendidikan_kontainer struct {
		Pendidikan string `json:"pendidikan"`
		ID         int    `json:"id_pendidikan"`
		Kategori   string `json:"kategori_pendidikan"`
	}

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(context.Background())

	query := `
		select description, id , category_name  from dev.kategori_pendidikan
	`

	row, err := tx.Query(ctx, query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	defer row.Close()

	var Tampung_pendidikan []Kategori_pendidikan_kontainer

	for row.Next() {
		var ambil Kategori_pendidikan_kontainer
		err := row.Scan(
			&ambil.Pendidikan,
			&ambil.ID,
			&ambil.Kategori,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			return
		}

		Tampung_pendidikan = append(Tampung_pendidikan, ambil)
	}

	if len(Tampung_pendidikan) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Kategori Pendidikan tersedia",
			"data":    Tampung_pendidikan,
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return

	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Kategori Pendidikan tidak ada",
			"data":    []Kategori_pendidikan_kontainer{},
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}
}
