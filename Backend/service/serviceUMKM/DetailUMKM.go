package serviceumkm

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Detailumkm(c *gin.Context) {

	type Data_detail_umkm_kontainer struct {
		ID         int       `json:"id_umkm"`
		Nama       string    `json:"nama_umkm"`
		Konten     string    `json:"konten_umkm"`
		Kategori   string    `json:"kategori_umkm"`
		KategoriID int       `json:"kategori_umkm_id"`
		Foto       string    `json:"foto_umkm"`
		DesaID     int       `json:"desa_id"`
		NoTelp     string    `json:"no_telp_umkm"`
		CreatedAt  time.Time `json:"created_at"`
		CreatedBy  string    `json:"created_by"`
		UpdatedAt  time.Time `json:"update_at"`
		UpdatedBy  string    `json:"updated_by"`
		LinkLokasi string    `json:"link_lokasi"`
		Alamat     string    `json:"alamat"`
	}

	id := c.Param("id")

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}

	var ambil Data_detail_umkm_kontainer
	var Tampung []Data_detail_umkm_kontainer

	query := `
	select 
		id_umkm,
		nama_umkm,
		konten_umkm,
		kategori_umkm,
		kategori_umkm_id,
		foto_umkm,
		desa_id,
		no_telp_umkm,
		created_at,
		created_by,
		update_at,
		updated_by,
		link_lokasi,
		alamat
	from umkm where id_umkm = $1
	
	`

	err = tx.QueryRow(ctx, query, id).Scan(
		&ambil.ID,
		&ambil.Nama,
		&ambil.Konten,
		&ambil.Kategori,
		&ambil.KategoriID,
		&ambil.Foto,
		&ambil.DesaID,
		&ambil.NoTelp,
		&ambil.CreatedAt,
		&ambil.CreatedBy,
		&ambil.UpdatedAt,
		&ambil.UpdatedBy,
		&ambil.LinkLokasi,
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

		c.JSON(http.StatusOK, gin.H{"status": false, "message": "Data tidak ditemukan", "data": []Data_detail_umkm_kontainer{}})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return

	}

}
