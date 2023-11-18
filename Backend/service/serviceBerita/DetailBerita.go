package serviceberita

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func DetailBerita(c *gin.Context) {

	type Data_detail_berita_kontainer struct {
		ID         int       `json:"id_berita"`
		Judul      string    `json:"judul"`
		SubJudul   string    `json:"sub_judul"`
		Deskripsi  string    `json:"deskripsi"`
		DesaID     int       `json:"desa_id"`
		FotoBerita string    `json:"foto_berita"`
		KategoriID int       `json:"kategori_id"`
		CreatedAt  time.Time `json:"created_at"`
		CreatedBy  string    `json:"created_by"`
		UpdateAt   time.Time `json:"update_at"`
		UpdatedBy  string    `json:"updated_by"`
	}

	id := c.Param("id")

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}

	var ambil Data_detail_berita_kontainer
	var Tampung []Data_detail_berita_kontainer

	query := `
	select 
		id_berita,
		judul,
		sub_judul,
		deskripsi,
		desa_id,
		foto_berita,
		kategori_id,
		created_at,
		created_by,
		update_at,
		updated_by
	from berita where id_berita = $1
	
	`

	err = tx.QueryRow(ctx, query, id).Scan(
		&ambil.ID,
		&ambil.Judul,
		&ambil.SubJudul,
		&ambil.Deskripsi,
		&ambil.DesaID,
		&ambil.FotoBerita,
		&ambil.KategoriID,
		&ambil.CreatedAt,
		&ambil.CreatedBy,
		&ambil.UpdateAt,
		&ambil.UpdatedBy,
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

		c.JSON(http.StatusOK, gin.H{"status": false, "message": "Data tidak ditemukan", "data": []Data_detail_berita_kontainer{}})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return

	}

}
