package servicekreativitas

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Kreatifitas_desa(c *gin.Context) {
	type Kreatifitas_kontainer struct {
		JudulKreatifitas string `json:"judul_kreatifitas"`
		Deskripsi        string `json:"deskripsi"`
		FotoKreatifitas  string `json:"foto_kreatifitas"`
		TotalPengunjung  string `json:"total_pengunjung"`
	}

	id := c.Param("id")

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}

	query := `
	select 

	a.judul_kreatifitas,
	a.deskripsi        ,
	a.foto_kreatifitas 

	from dev.kreatifitas a, dev.desa b
	where a.desa_id = b.id_desa
	and b.id_desa = $1
	
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

	var Tampung_kreatifitas_desa []Kreatifitas_kontainer

	for row.Next() {
		var ambil Kreatifitas_kontainer

		err := row.Scan(
			&ambil.JudulKreatifitas,
			&ambil.Deskripsi,
			&ambil.FotoKreatifitas,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			err = tx.Commit(ctx)
			if err != nil {
				panic(err.Error())
			}
			return
		}

		Tampung_kreatifitas_desa = append(Tampung_kreatifitas_desa, ambil)
	}

	if len(Tampung_kreatifitas_desa) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Data tersedia",
			"data":    Tampung_kreatifitas_desa,
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return

	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Data tidak ada",
			"data":    []Kreatifitas_kontainer{},
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}
}
