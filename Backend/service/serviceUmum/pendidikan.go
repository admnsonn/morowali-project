package serviceumum

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Pendidikan(c *gin.Context) {
	type Pendidikan_kontainer struct {
		TK  int `json:"tk"`
		SD  int `json:"sd"`
		SMP int `json:"smp"`
		SMA int `json:"sma"`
		SMK int `json:"smk"`
		D1  int `json:"d1"`
		D2  int `json:"d2"`
		D3  int `json:"d3"`
		D4  int `json:"d4"`
		S1  int `json:"s1"`
		S2  int `json:"s2"`
		S3  int `json:"s3"`
	}

	type Database struct {
		Nama         string `json:"nama"`
		PendidikanID int    `json:"pendidikan_id"`
		Pendidikan   string `json:"pendidikan"`
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

	pendidikan := `
	Select 
			a.nama_lengkap,
			a.pendidikan_id ,
			e.category_name 
		FROM
			dev.pengguna a, dev.kategori_pendidikan e
			where 
				a.pendidikan_id  = e.id and 
				a.role_id  = 2 and 
				a.desa_id = $1
	`

	row, err := tx.Query(ctx, pendidikan, input.IDDesa)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		err = tx.Commit(ctx)
		if err != nil {
			panic(err.Error())
		}
		return
	}

	defer row.Close()

	var Tampung_pendidikan []Database

	for row.Next() {
		var ambil Database
		row.Scan(
			&ambil.Nama,
			&ambil.PendidikanID,
			&ambil.Pendidikan,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			err = tx.Commit(ctx)
			if err != nil {
				panic(err.Error())
			}
			return
		}

		Tampung_pendidikan = append(Tampung_pendidikan, ambil)

	}

	var Datapendidikan []Pendidikan_kontainer
	var countPendidikan Pendidikan_kontainer

	for _, v := range Tampung_pendidikan {
		switch v.PendidikanID {
		case 1:
			countPendidikan.TK++
		case 2:
			countPendidikan.SD++
		case 3:
			countPendidikan.SMP++
		case 4:
			countPendidikan.SMA++
		case 5:
			countPendidikan.SMK++
		case 6:
			countPendidikan.D1++
		case 7:
			countPendidikan.D2++
		case 8:
			countPendidikan.D3++
		case 9:
			countPendidikan.D4++
		case 10:
			countPendidikan.S1++
		case 11:
			countPendidikan.S2++
		case 12:
			countPendidikan.S3++
		}
	}

	Datapendidikan = append(Datapendidikan, countPendidikan)

	fmt.Println("INI ADALAH DATA PENDIDIKNA", Datapendidikan[0])

	if len(Tampung_pendidikan) != 0 {

		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Pendidikan tersedia",
			"data":    Datapendidikan,
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return

	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Pendidikan tidak ada",
			"data":    []Database{},
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}
}
