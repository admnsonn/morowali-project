package servicelayanandesa

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func List(c *gin.Context) {
	type Data_kontainer struct {
		ID      int    `json:"id"`
		Nama    string `json:"nama"`
		NomorHP string `json:"nomor_hp"`
		Alamat  string `json:"alamat"`
		Layanan string `json:"layanan"`
		Status  string `json:"status"`
		NIK     string `json:"nik"`
	}

	type Request struct {
		IDDesa string `json:"id_desa"`
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

	umkm := `
	select 
	a.id      ,
	a.nik     ,
	a.nama    ,
	a.nomor_hp,
	a.alamat  ,
	a.layanan ,
	a.status  
	from dev.layanan_sementara a, dev.desa b
	where a.desa_id = b.id_desa
	and b.id_desa = $1
	`

	row, err := tx.Query(ctx, umkm, input.IDDesa)

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
			&ambil.ID,
			&ambil.NIK,
			&ambil.Nama,
			&ambil.NomorHP,
			&ambil.Alamat,
			&ambil.Layanan,
			&ambil.Status,
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

	// fmt.Println(Tampung)

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
