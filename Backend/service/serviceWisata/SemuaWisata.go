package servicewisata

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Semuawisata(c *gin.Context) {
	type Data_wisata struct {
		Nama       string `json:"nama_wisata"`
		Alamat     string `json:"alamat"`
		Foto       string `json:"foto_wisata"`
		NoTelp     string `json:"no_telp"`
		ID         int    `json:"id_wisata"`
		IDkategori int    `json:"id_kategori"`
		Kategori   string `json:"kategori"`
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

	wisata := `
	select 	a.nama_wisata ,
			a.alamat ,
			a.foto_wisata ,
			a.no_telp ,
			a.id_wisata ,
			a.idkategoriwisata ,
			c.nama_kategori 
		from dev.wisata a, dev.desa b  , DEV.kategoriwisata C
		where a.desa_id  = b.id_desa  and a.idkategoriwisata = c.idkategoriwisata 
	and b.id_desa = $1
	`

	row, err := tx.Query(ctx, wisata, input.IDDesa)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		err = tx.Commit(ctx)
		if err != nil {
			panic(err.Error())
		}
		return
	}

	defer row.Close()

	var Tampung_wisata []Data_wisata

	for row.Next() {
		var ambil Data_wisata
		row.Scan(
			&ambil.Nama,
			&ambil.Alamat,
			&ambil.Foto,
			&ambil.NoTelp,
			&ambil.ID,
			&ambil.IDkategori,
			&ambil.Kategori,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			err = tx.Commit(ctx)
			if err != nil {
				panic(err.Error())
			}
			return
		}

		Tampung_wisata = append(Tampung_wisata, ambil)

	}

	if len(Tampung_wisata) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Wisata tersedia",
			"data":    Tampung_wisata,
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return

	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "wisata tidak ada",
			"data":    []Data_wisata{},
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}

}
