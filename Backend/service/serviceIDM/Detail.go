package serviceidm

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Detail(c *gin.Context) {
	type Data_kontainer struct {
		ID         int    `json:"id"`
		Tahun      string `json:"tahun"`
		Indikator  string `json:"indikator"`
		Skor       string `json:"skor"`
		Keterangan string `json:"keterangan"`
		Kegiatan   string `json:"kegiatan"`
		Nilai      string `json:"nilai"`
		Pusat      string `json:"pusat"`
		Provinsi   string `json:"provinsi"`
		Kabupaten  string `json:"kabupaten"`
		Desa       string `json:"desa"`
		CSR        string `json:"csr"`
		Lainya     string `json:"lainya"`
	}

	type Request struct {
		IDDesa string `json:"desa_id"`
		Tahun  string `json:"tahun"`
	}

	var input Request

	if c.GetHeader("content-type") == "application/x-www-form-urlencoded" || c.GetHeader("content-type") == "application/x-www-form-urlencoded; charset=utf-8" {

		if err := c.Bind(&input); err != nil {
			fmt.Print("masuk sini")
			return
		}

	} else {

		if err := c.BindJSON(&input); err != nil {
			fmt.Print("masuk sini")
			return
		}

	}

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(context.Background())

	query1 := `
	select 
		id        ,
		tahun     ,
		indikator ,
		skor      ,
		keterangan,
		kegiatan  ,
		nilai     ,
		pusat     ,
		provinsi  ,
		kabupaten ,
		csr       ,
		desa 	  ,
		lainya    
	from dev.idm where tahun = $1 and desa_id = $2

	`

	row, err := tx.Query(ctx, query1, input.Tahun, input.IDDesa)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		err = tx.Commit(ctx)
		if err != nil {
			panic(err.Error())
		}
		return
	}

	defer row.Close()

	var Tampung1 []Data_kontainer

	for row.Next() {
		var ambil Data_kontainer
		err := row.Scan(
			&ambil.ID,
			&ambil.Tahun,
			&ambil.Indikator,
			&ambil.Skor,
			&ambil.Keterangan,
			&ambil.Kegiatan,
			&ambil.Nilai,
			&ambil.Pusat,
			&ambil.Provinsi,
			&ambil.Kabupaten,
			&ambil.Desa,
			&ambil.CSR,
			&ambil.Lainya,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			err = tx.Commit(ctx)
			if err != nil {
				panic(err.Error())
			}
			return
		}

		Tampung1 = append(Tampung1, ambil)

	}

	query2 := `
	select 
		id        ,
		tahun     ,
		indikator ,
		skor      ,
		keterangan,
		kegiatan  ,
		nilai     ,
		pusat     ,
		provinsi  ,
		kabupaten ,
		csr       ,
		desa 	  ,
		lainya    
	from dev.iks where tahun = $1 and desa_id = $2

	`

	row2, err := tx.Query(ctx, query2, input.Tahun, input.IDDesa)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		err = tx.Commit(ctx)
		if err != nil {
			panic(err.Error())
		}
		return
	}

	defer row2.Close()

	var Tampung2 []Data_kontainer

	for row2.Next() {
		var ambil Data_kontainer
		err := row2.Scan(
			&ambil.ID,
			&ambil.Tahun,
			&ambil.Indikator,
			&ambil.Skor,
			&ambil.Keterangan,
			&ambil.Kegiatan,
			&ambil.Nilai,
			&ambil.Pusat,
			&ambil.Provinsi,
			&ambil.Kabupaten,
			&ambil.Desa,
			&ambil.CSR,
			&ambil.Lainya,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			err = tx.Commit(ctx)
			if err != nil {
				panic(err.Error())
			}
			return
		}

		Tampung2 = append(Tampung2, ambil)

	}

	query3 := `
	select 
		id        ,
		tahun     ,
		indikator ,
		skor      ,
		keterangan,
		kegiatan  ,
		nilai     ,
		pusat     ,
		provinsi  ,
		kabupaten ,
		csr       ,
		desa 	  ,
		lainya    
	from dev.ike where tahun = $1 and desa_id = $2

	`

	row3, err := tx.Query(ctx, query3, input.Tahun, input.IDDesa)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		err = tx.Commit(ctx)
		if err != nil {
			panic(err.Error())
		}
		return
	}

	defer row3.Close()

	var Tampung3 []Data_kontainer

	for row3.Next() {
		var ambil Data_kontainer
		err := row3.Scan(
			&ambil.ID,
			&ambil.Tahun,
			&ambil.Indikator,
			&ambil.Skor,
			&ambil.Keterangan,
			&ambil.Kegiatan,
			&ambil.Nilai,
			&ambil.Pusat,
			&ambil.Provinsi,
			&ambil.Kabupaten,
			&ambil.Desa,
			&ambil.CSR,
			&ambil.Lainya,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			err = tx.Commit(ctx)
			if err != nil {
				panic(err.Error())
			}
			return
		}

		Tampung3 = append(Tampung3, ambil)

	}

	if len(Tampung1) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Data tersedia",
			"idm":     Tampung1,
			"iks":     Tampung2,
			"ike":     Tampung3,
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
