package servicepotensidesa

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Detailpotensi_desa(c *gin.Context) {

	type Data_detail_potensi_desa_kontainer struct {
		ID           int       `json:"id_potensi"`
		JudulPotensi string    `json:"judul_potensi"`
		Deskripsi    string    `json:"deskripsi"`
		FotoPotensi  string    `json:"foto_potensi_desa"`
		DesaID       int       `json:"desa_id"`
		CreatedAt    time.Time `json:"created_at"`
		CreatedBy    string    `json:"created_by"`
		UpdatedAt    time.Time `json:"update_at"`
		UpdatedBy    string    `json:"updated_by"`
		SubJudul     string    `json:"sub_judul"`
	}

	id := c.Param("id")

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}

	var ambil Data_detail_potensi_desa_kontainer
	var Tampung []Data_detail_potensi_desa_kontainer

	query := `
	select 
		id_potensi       ,
		judul_potensi    ,
		deskripsi        ,
		foto_potensi_desa,
		desa_id          ,
		created_at       ,
		created_by       ,
		update_at        ,
		updated_by       ,
		sub_judul        
	from potensi_desa where id_potensi = $1
	
	`

	err = tx.QueryRow(ctx, query, id).Scan(
		&ambil.ID,
		&ambil.JudulPotensi,
		&ambil.Deskripsi,
		&ambil.FotoPotensi,
		&ambil.DesaID,
		&ambil.CreatedAt,
		&ambil.CreatedBy,
		&ambil.UpdatedAt,
		&ambil.UpdatedBy,
		&ambil.SubJudul,
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

		c.JSON(http.StatusOK, gin.H{"status": false, "message": "Data tidak ditemukan", "data": []Data_detail_potensi_desa_kontainer{}})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return

	}

}
