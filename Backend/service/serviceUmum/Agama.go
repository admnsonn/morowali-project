package serviceumum

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Agama(c *gin.Context) {
	type Data_kontainer struct {
		Data1 int `json:"islam"`
		Data2 int `json:"kristen_protestan"`
		Data3 int `json:"kristen_katolik"`
		Data4 int `json:"hindu"`
		Data5 int `json:"budha"`
		Data6 int `json:"konghucu"`
	}

	type Acuan_kontainer struct {
		Nama    string `json:"nama_lengkap"`
		AgamaID int    `json:"agama_id"`
		Agama   string `json:"nama_agama"`
	}

	id := c.Param("id")

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}
	defer tx.Rollback(ctx)

	pendidikan := `
	Select 
			a.nama_lengkap,
			a.agama_id  ,
			e.nama_agama  
		FROM
			dev.pengguna a, dev.agama e
			where 
				a.pendidikan_id  = e.id_agama  and 
				a.role_id  = 2 and 
				a.desa_id = $1
	`

	row, err := tx.Query(ctx, pendidikan, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		err = tx.Commit(ctx)
		if err != nil {
			panic(err.Error())
		}
		return
	}

	defer row.Close()

	var Tampung_pendidikan []Acuan_kontainer

	for row.Next() {
		var ambil Acuan_kontainer
		row.Scan(
			&ambil.Nama,
			&ambil.AgamaID,
			&ambil.Agama,
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

	var Dataagama []Data_kontainer
	var hitung Data_kontainer

	for _, v := range Tampung_pendidikan {
		switch v.AgamaID {
		case 1:
			hitung.Data1++
		case 2:
			hitung.Data2++
		case 3:
			hitung.Data3++
		case 4:
			hitung.Data4++
		case 5:
			hitung.Data5++
		case 6:
			hitung.Data6++
		}
	}

	Dataagama = append(Dataagama, hitung)

	// fmt.Println("INI ADALAH DATA PENDIDIKNA", Dataagama[0])

	if len(Tampung_pendidikan) != 0 {

		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Data tersedia",
			"data":    Dataagama,
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
			"data":    []Acuan_kontainer{},
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}
}
