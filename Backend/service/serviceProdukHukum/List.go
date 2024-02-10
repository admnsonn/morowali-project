package serviceprodukhukum

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func List(c *gin.Context) {
	type Data_kontainer struct {
		ID         int    `json:"id_hukum"`
		Nama       string `json:"nama_produk"`
		Kategori   string `json:"nama_kategori"`
		IDKategori int    `json:"id-kategori"`
		File       string `json:"file"`
		Terbit     string `json:"terbit"`
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
	select a.id, a.nama_produk, TO_CHAR(a.tanggal_terbit, 'DD-MM-YYYY') AS tahun_terbit, b.nama_kategori, b.id , a.file
	from dev.produk_hukum a, dev.kategori_jenis_hukum b, dev.desa c 
	where a.desa_id = c.id_desa
	and a.kategori_id = b.id
	and c.id_desa = $1
	ORDER BY a.id DESC
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
			&ambil.Nama,
			&ambil.Terbit,
			&ambil.Kategori,
			&ambil.IDKategori,
			&ambil.File,
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
