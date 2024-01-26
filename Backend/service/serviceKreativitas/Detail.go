package servicekreativitas

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Detail(c *gin.Context) {
	type Kreatifitas struct {
		ID               int       `json:"id_kreatifitas"`
		JudulKreatifitas string    `json:"judul_kreatifitas"`
		Deskripsi        string    `json:"deskripsi"`
		FotoKreatifitas  string    `json:"foto_kreatifitas"`
		DesaID           int       `json:"desa_id"`
		CreatedAt        time.Time `json:"created_at"`
		CreatedBy        string    `json:"created_by"`
		UpdatedAt        time.Time `json:"updated_at"`
		UpdatedBy        string    `json:"updated_by"`
		Pengunjung       int       `json:"pengunjung"`
	}

	id := c.Param("id")

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}

	var ambil Kreatifitas
	var Tampung []Kreatifitas

	query := `
	select 
		id_kreatifitas   ,
		judul_kreatifitas,
		deskripsi        ,
		foto_kreatifitas ,
		desa_id          ,
		created_at       ,
		created_by       ,
		update_at        ,
		updated_by       ,
		pengunjung       
	from dev.kreatifitas k where id_kreatifitas = $1
	
	`

	err = tx.QueryRow(ctx, query, id).Scan(
		&ambil.ID,
		&ambil.JudulKreatifitas,
		&ambil.Deskripsi,
		&ambil.FotoKreatifitas,
		&ambil.DesaID,
		&ambil.CreatedAt,
		&ambil.CreatedBy,
		&ambil.UpdatedAt,
		&ambil.UpdatedBy,
		&ambil.Pengunjung,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		err = tx.Commit(ctx)
		if err != nil {
			panic(err.Error())
		}
		return
	}

	Tampung = append(Tampung, ambil)

	if len(Tampung) != 0 {

		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Data ditemukan",
			"data":    Tampung,
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return

	} else {

		c.JSON(http.StatusOK, gin.H{"status": false, "message": "Data tidak ditemukan", "data": []Kreatifitas{}})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return

	}
}
