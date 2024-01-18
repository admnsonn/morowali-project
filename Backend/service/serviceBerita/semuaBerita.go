package serviceberita

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func SemuaBerita(c *gin.Context) {
	type Data_berita struct {
		ID         int    `json:"id_berita"`
		Judul      string `json:"judul"`
		SubJudul   string `json:"sub_judul"`
		Deskripsi  string `json:"deskripsi"`
		FotoBerita string `json:"foto_berita"`
		Kategori   string `json:"kategori"`
	}

	type Request struct {
		IDDesa string `json:"id_desa"`
	}

	var input Request

	if c.GetHeader("content-type") == "application/x-www-form-urlencoded" || c.GetHeader("content-type") == "application/x-www-form-urlencoded; charset=utf-8" {

		if err := c.Bind(&input); err != nil {
			return
		}

	} else {

		if err := c.BindJSON(&input); err != nil {
			return
		}

	}

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(context.Background())

	var id_desa int

	fmt.Print(id_desa)

	berita := `
	SELECT 
		a.id_berita,
		a.judul,
		a.sub_judul,
		a.deskripsi,
		a.foto_berita,
		b.kategori
	FROM
		dev.berita a, dev.kategori_berita b
	WHERE
		a.kategori_id = b.id_kategori_berita
		AND a.desa_id = $1
	`

	row, err := tx.Query(ctx, berita, input.IDDesa)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		err = tx.Commit(ctx)
		if err != nil {
			panic(err.Error())
		}
		return
	}

	defer row.Close()

	var Tampung_berita []Data_berita

	for row.Next() {
		var ambil Data_berita
		row.Scan(
			&ambil.ID,
			&ambil.Judul,
			&ambil.SubJudul,
			&ambil.Deskripsi,
			&ambil.FotoBerita,
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

		Tampung_berita = append(Tampung_berita, ambil)

	}

	// fmt.Println(Tampung_berita)

	if len(Tampung_berita) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Berita tersedia",
			"data":    Tampung_berita,
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return

	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Berita tidak ada",
			"data":    []Data_berita{},
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}

}
