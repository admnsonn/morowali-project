package servicewilayahdesa

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Wilayah_desa(c *gin.Context) {
	type Wilayah_desa struct {
		Longitude      string `json:"longitude"`
		Latitude       string `json:"latitude"`
		Status_wilayah string `json:"status_wilayah"`
		LuasWilayah    string `json:"luas_wilayah"`
		ID             string `json:"id_desa"`
	}

	id := c.Param("id")

	var ambil Wilayah_desa
	var Tampung []Wilayah_desa

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(ctx)

	query := `
		select luas_wilayah , longitude , latitude , status_wilayah  from dev.desa where id_desa = $1
	`

	err = tx.QueryRow(ctx, query, id).Scan(
		&ambil.LuasWilayah,
		&ambil.Longitude,
		&ambil.Latitude,
		&ambil.Status_wilayah,
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

		if Tampung[0].Status_wilayah == "" {
			Tampung[0].Status_wilayah = "Tambah"
		}

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

		c.JSON(http.StatusOK, gin.H{"status": false, "message": "Data tidak ditemukan", "data": []Wilayah_desa{}})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return

	}
}
