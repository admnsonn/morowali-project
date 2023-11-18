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
		Deskripsi  string `json:"deskripsi"`
		FotoBerita string `json:"foto_berita"`
		Kategori   string `json:"kategori"`
	}

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}

	var id_desa int

	desa := `
		select id_desa from desa d where nama_desa = 'Morowali'
	`

	err = tx.QueryRow(ctx, desa).Scan(&id_desa)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		err = tx.Commit(ctx)
		if err != nil {
			panic(err.Error())
		}
		return
	}

	fmt.Print(id_desa)

	berita := `
	SELECT 
		a.id_berita,
		a.judul,
		a.deskripsi,
		a.foto_berita,
		b.kategori
	FROM
		berita a, kategori_berita b
	WHERE
		a.kategori_id = b.id_kategori_berita
		AND a.desa_id = $1
	`

	row, err := tx.Query(ctx, berita, id_desa)

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

	fmt.Println(Tampung_berita)

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
