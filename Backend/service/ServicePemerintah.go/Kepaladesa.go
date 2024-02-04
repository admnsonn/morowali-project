package servicepemerintahgo

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Kepaladesa(c *gin.Context) {
	type Data_kepaladesa struct {
		ID_kepala_desa   int    `json:"id_kepala_desa"`
		Nama             string `json:"nama"`
		Foto_kepala_desa string `json:"foto_kepala_desa"`
		Jabatan_dimulai  string `json:"jabatan_dimulai"`
		Jabatan_berakhir string `json:"jabatan_berakhir"`
		Desa_id          int    `json:"desa_id"`
	}

	id := c.Param("id")

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}

	kepaladesa := `
	select 
		a.id_kepala_desa ,
		a.nama ,
		a.foto_kepala_desa, 
		a.jabatan_dimulai, 
		a.jabatan_berakhir, 
		a.desa_id
	from dev.kepala_desa a, dev.desa d  
	where a.desa_id  = d.id_desa 
	and d.id_desa = $1
	`

	row, err := tx.Query(ctx, kepaladesa, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		err = tx.Commit(ctx)
		if err != nil {
			panic(err.Error())
		}
		return
	}

	defer row.Close()

	var Tampung_kepaladesa []Data_kepaladesa

	for row.Next() {
		var ambil Data_kepaladesa
		err :=
			row.Scan(
				&ambil.ID_kepala_desa,
				&ambil.Nama,
				&ambil.Foto_kepala_desa,
				&ambil.Jabatan_dimulai,
				&ambil.Jabatan_berakhir,
				&ambil.Desa_id,
			)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			err = tx.Commit(ctx)
			if err != nil {
				panic(err.Error())
			}
			return
		}

		Tampung_kepaladesa = append(Tampung_kepaladesa, ambil)
	}

	kepaladesa2 := `
		SELECT 
			a.id_kepala_desa,
			a.nama,
			a.foto_kepala_desa,
			a.jabatan_dimulai,
			a.jabatan_berakhir,
			a.desa_id
		FROM 
			dev.kepala_desa a
		INNER JOIN 
			dev.desa d ON a.desa_id = d.id_desa
		WHERE 
			d.id_desa = $1
		ORDER BY 
			a.jabatan_dimulai DESC 
		LIMIT 1;
	`

	row2, err := tx.Query(ctx, kepaladesa2, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		err = tx.Commit(ctx)
		if err != nil {
			panic(err.Error())
		}
		return
	}

	defer row2.Close()

	var Tampung_kepaladesa2 []Data_kepaladesa

	for row2.Next() {
		var ambil Data_kepaladesa
		err :=
			row2.Scan(
				&ambil.ID_kepala_desa,
				&ambil.Nama,
				&ambil.Foto_kepala_desa,
				&ambil.Jabatan_dimulai,
				&ambil.Jabatan_berakhir,
				&ambil.Desa_id,
			)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			err = tx.Commit(ctx)
			if err != nil {
				panic(err.Error())
			}
			return
		}

		Tampung_kepaladesa2 = append(Tampung_kepaladesa2, ambil)
	}

	if len(Tampung_kepaladesa) != 0 {

		c.JSON(http.StatusOK, gin.H{
			"status":               true,
			"message":              "Kepala Desa tersedia",
			"data":                 Tampung_kepaladesa,
			"kepala_desa_saat_ini": Tampung_kepaladesa2,
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return

	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Kepala desa tidak ada",
			"data":    []Data_kepaladesa{},
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}
}
