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
		Nama   string `json:"nama_wisata"`
		Alamat string `json:"alamat"`
		Foto   string `json:"foto_wisata"`
		NoTelp string `json:"no_telp"`
		ID     int    `json:"id_wisata"`
	}

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}

	wisata := `
	select 
		a.nama_wisata ,
		a.alamat ,
		a.foto_wisata ,
		a.no_telp ,
		a.id_wisata 
	from wisata a, desa d  
	where a.desa_id  = d.id_desa 
	and d.nama_desa = 'Morowali'
	`

	row, err := tx.Query(ctx, wisata)

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

	fmt.Println(Tampung_wisata)

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
