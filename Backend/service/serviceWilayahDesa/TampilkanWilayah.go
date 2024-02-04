package servicewilayahdesa

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Wilayah_produksi(c *gin.Context) {
	type Wilayah_Perkebunan_container struct {
		LuasWilayah  string `json:"luas_wilayah"`
		NamaProduksi string `json:"nama_produksi"`
	}

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}

	query := `
		select a.luas_wilayah, a.nama_produksi 
		from dev.produksi_desa a, dev.desa d 
		where a.desa_id = d.id_desa 
		and d.id_desa = $1
	`

	row, err := tx.Query(ctx, query)

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

		total_luas_wilayah := 0

		for _, v := range Tampung_data {
			replacedString := strings.Replace(v.LuasWilayah, " hektar", "", -1)
			angka, err := strconv.Atoi(replacedString)
			if err != nil {
				c.JSON(403, gin.H{
					"status":  false,
					"message": "Data Error!",
					"data":    []Wilayah_Perkebunan_container{},
				})
				err = tx.Commit(ctx)
				if err != nil {
					fmt.Println(err.Error())
				}
				return
			}

			total_luas_wilayah = total_luas_wilayah + angka

		}

		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Data Luas Perkebunan Ada!",
			"data":    total_luas_wilayah,
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

func Total_warga(c *gin.Context) {
	type Wilayah_Perkebunan_container struct {
		TotalWarga int `json:"total_warga"`
	}

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}

	query := `
		SELECT COUNT(*) AS total_warga
		FROM dev.pengguna
		WHERE desa_id  = 1 and role_pengguna = 'Warga';
	`

	row, err := tx.Query(ctx, query)

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
			&ambil.TotalWarga,
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
			"status":  false,
			"message": "Data Luas Perkebunan Tidak Ada!",
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

func Luas_desa(c *gin.Context) {
	type Wilayah_Perkebunan_container struct {
		LuasDesa string `json:"luas_wilayah"`
		Longi    string `json:"longitude"`
		Lati     string `json:"latitude"`
	}

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}

	query := `
		select luas_wilayah, longitude, latitude from dev.desa where id_desa = 1
	`

	row, err := tx.Query(ctx, query)

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
			&ambil.LuasDesa,
			&ambil.Longi,
			&ambil.Lati,
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
			"status":  false,
			"message": "Data Luas Desa Tidak Ada!",
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
			"message": "Data Luas Desa Tidak Ada!",
			"data":    []Wilayah_Perkebunan_container{},
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}
}
