package serviceumum

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func GaleriFoto(c *gin.Context) {
	type Data_kontainer struct {
		Foto string `json:"foto"`
		Nama string `json:"nama"`
	}

	id := c.Param("id")

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}
	defer tx.Rollback(ctx)

	query1 := `
	select foto_berita , judul  from dev.berita where desa_id = $1
	`

	row, err := tx.Query(ctx, query1, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		err = tx.Commit(ctx)
		if err != nil {
			panic(err.Error())
		}
		return
	}

	defer row.Close()

	var Tampung []Data_kontainer

	for row.Next() {
		var ambil Data_kontainer
		err := row.Scan(
			&ambil.Nama,
			&ambil.Foto,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			err = tx.Commit(ctx)
			if err != nil {
				panic(err.Error())
			}
			return
		}

		Tampung = append(Tampung, ambil)
	}

	query2 := `
	select foto_potensi_desa , judul_potensi  from dev.potensi_desa where desa_id = $1
	`

	row2, err := tx.Query(ctx, query2, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		err = tx.Commit(ctx)
		if err != nil {
			panic(err.Error())
		}
		return
	}

	defer row2.Close()

	for row2.Next() {
		var ambil Data_kontainer
		err := row2.Scan(
			&ambil.Nama,
			&ambil.Foto,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			err = tx.Commit(ctx)
			if err != nil {
				panic(err.Error())
			}
			return
		}

		Tampung = append(Tampung, ambil)
	}

	query3 := `
	select foto_kreatifitas , judul_kreatifitas  from dev.kreatifitas where desa_id = $1
	`

	row3, err := tx.Query(ctx, query3, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		err = tx.Commit(ctx)
		if err != nil {
			panic(err.Error())
		}
		return
	}

	defer row3.Close()

	for row3.Next() {
		var ambil Data_kontainer
		err := row3.Scan(
			&ambil.Nama,
			&ambil.Foto,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			err = tx.Commit(ctx)
			if err != nil {
				panic(err.Error())
			}
			return
		}

		Tampung = append(Tampung, ambil)
	}

	if len(Tampung) != 0 {

		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Data tersedia",
			"data":    Tampung,
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
			"data":    []Data_kontainer{},
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}

}
