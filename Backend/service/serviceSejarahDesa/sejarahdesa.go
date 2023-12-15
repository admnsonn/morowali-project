package servicesejarahdesa

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Sejarahdesa(c *gin.Context) {
	type Data_sejarahdesa struct {
		ID_Desa   int    `json:"id_sejarah_desa"`
		Deskripsi string `json:"deskripsi"`
		Desa_ID   int    `json:"desa_id"`
	}

	id := c.Param("id")

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}

	sejarahdesa := `
	select 
		a.id_sejarah_desa ,
		a.deskripsi ,
		a.desa_id 
	from dev.sejarah_desa a, dev.desa d  
	where a.desa_id  = d.id_desa 
	and d.id_desa = $1
	`

	row, err := tx.Query(ctx, sejarahdesa, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		err = tx.Commit(ctx)
		if err != nil {
			panic(err.Error())
		}
		return
	}

	defer row.Close()

	var Tampung_sejarahdesa []Data_sejarahdesa

	for row.Next() {
		var ambil Data_sejarahdesa
		row.Scan(
			&ambil.ID_Desa,
			&ambil.Deskripsi,
			&ambil.Desa_ID,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			err = tx.Commit(ctx)
			if err != nil {
				panic(err.Error())
			}
			return
		}

		Tampung_sejarahdesa = append(Tampung_sejarahdesa, ambil)

	}

	fmt.Println(Tampung_sejarahdesa)

	if len(Tampung_sejarahdesa) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Wisata tersedia",
			"data":    Tampung_sejarahdesa,
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return

	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Sejarah desa tidak ada",
			"data":    []Data_sejarahdesa{},
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}
}
