package serviceproduksi

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Detail(c *gin.Context) {
	type Wilayah_Perkebunan_container struct {
		LuasWilayah   string `json:"luas_wilayah"`
		NamaProduksi  string `json:"nama_produksi"`
		TotalProduksi int    `json:"total_produksi"`
		ID            int    `json:"id_produksi"`
	}

	id := c.Param("id")

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}

	query := `
		select a.luas_wilayah, a.nama_produksi , a.jumlah, a.id
		from dev.produksi_desa a, dev.desa d 
		where a.desa_id = d.id_desa 
		and d.id_desa = $1
	`

	row, err := tx.Query(ctx, query, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		err = tx.Commit(ctx)
		if err != nil {
			panic(err.Error())
		}
		return
	}

	defer row.Close()

	var Tampung_data []Wilayah_Perkebunan_container

	for row.Next() {
		var ambil Wilayah_Perkebunan_container
		err := row.Scan(
			&ambil.LuasWilayah,
			&ambil.NamaProduksi,
			&ambil.TotalProduksi,
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

		Tampung_data = append(Tampung_data, ambil)
	}

	if len(Tampung_data) != 0 {

		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Data Luas Perkebunan Ada!",
			"data":    Tampung_data,
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Data Luas Perkebunan Tidak Ada!",
			"data":    []Wilayah_Perkebunan_container{},
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}
}
