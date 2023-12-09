package sevicewargadesabyamin

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Warga_desa_by_admin(c *gin.Context) {
	type Data_list_warga_kontainer struct {
		IDPengguna     int    `json:"id_pengguna"`
		NIK            int    `json:"nik"`
		NamaLengkap    string `json:"nama_lengkap"`
		AlamatPengguna string `json:"alamat_pengguna"`
		NoTelp         string `json:"no_telp"`
	}

	type Request struct {
		IDPengguna int    `json:"id_pengguna"`
		RTPengguna string `json:"rt"`
		RWPengguna string `json:"rw"`
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

	query := `
		select 
			id_pengguna          ,
			nik                  ,
			nama_lengkap         ,
			alamat_pengguna      ,
			rt                   ,
			rw                   ,
			kode_pos             ,
			no_telp              
		from dev.pengguna 
		where desa_id = $1
		and role_id = 2
	`

	row, err := tx.Query(ctx, query, input.IDPengguna)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		err = tx.Commit(ctx)
		if err != nil {
			panic(err.Error())
		}
		return
	}

	defer row.Close()

	var Tampung_list_warga []Data_list_warga_kontainer
	var rt, rw, kode_pos string

	for row.Next() {
		var ambil Data_list_warga_kontainer

		err := row.Scan(
			&ambil.IDPengguna,
			&ambil.NIK,
			&ambil.NamaLengkap,
			&ambil.AlamatPengguna,
			&rt,
			&rw,
			&kode_pos,
			&ambil.NoTelp,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			err = tx.Commit(ctx)
			if err != nil {
				panic(err.Error())
			}
			return
		}

		Tampung_list_warga = append(Tampung_list_warga, ambil)
	}

	if len(Tampung_list_warga) > 0 {

		Tampung_list_warga[0].AlamatPengguna = Tampung_list_warga[0].AlamatPengguna + ", RT: " + rt + ", RW: " + rw + ", Kode Post: " + kode_pos

		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Wisata tersedia",
			"data":    Tampung_list_warga,
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	} else {

		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Wisata tidak tersedia !",
			"data":    []Data_list_warga_kontainer{},
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return

	}

}
