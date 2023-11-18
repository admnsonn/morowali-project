package serviceumkm

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Semuaumkm(c *gin.Context) {
	type Data_umkm struct {
		ID       int    `json:"id_umkm"`
		Nama     string `json:"nama_umkm"`
		Kategori string `json:"kategori_umkm"`
		Foto     string `json:"foto_umkm"`
		NoTelp   string `json:"no_telp_umkm"`
		Alamat   string `json:"alamat"`
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

	umkm := `
	select 
		a.id_umkm , 
		a.nama_umkm , 
		a.kategori_umkm , 
		a.foto_umkm , 
		a.no_telp_umkm , 
		a.alamat 
	from umkm a, desa b 
	where a.desa_id = b.id_desa 
	and b.id_desa = $1
	`

	row, err := tx.Query(ctx, umkm, id_desa)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		err = tx.Commit(ctx)
		if err != nil {
			panic(err.Error())
		}
		return
	}

	defer row.Close()

	var Tampung_umkm []Data_umkm

	for row.Next() {
		var ambil Data_umkm
		row.Scan(
			&ambil.ID,
			&ambil.Nama,
			&ambil.Kategori,
			&ambil.Foto,
			&ambil.NoTelp,
			&ambil.Alamat,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			err = tx.Commit(ctx)
			if err != nil {
				panic(err.Error())
			}
			return
		}

		Tampung_umkm = append(Tampung_umkm, ambil)

	}

	fmt.Println(Tampung_umkm)

	if len(Tampung_umkm) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "umkm tersedia",
			"data":    Tampung_umkm,
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return

	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "umkm tidak ada",
			"data":    []Data_umkm{},
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}

}
