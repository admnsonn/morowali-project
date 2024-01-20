package sevicewargadesabyamin

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func JK_pengguna(c *gin.Context) {
	type JenisKelamin_kontainer struct {
		JenisKelamin string `json:"jenis_kelamin"`
		ID           int    `json:"id_jenis_kelamin"`
	}

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(context.Background())

	query := `
		select jenis_kelamin , id_jk  from dev.jenis_kelamin
	`

	row, err := tx.Query(ctx, query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	defer row.Close()

	var Tampung_JK []JenisKelamin_kontainer

	for row.Next() {
		var ambil JenisKelamin_kontainer
		err := row.Scan(
			&ambil.JenisKelamin,
			&ambil.ID,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			return
		}

		Tampung_JK = append(Tampung_JK, ambil)
	}

	if len(Tampung_JK) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Kategori Jenis Kelamin tersedia",
			"data":    Tampung_JK,
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return

	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Kategori Jenis Kelamin tidak ada",
			"data":    []JenisKelamin_kontainer{},
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}
}
