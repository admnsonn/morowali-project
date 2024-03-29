package servicewisata

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Detailwisata(c *gin.Context) {

	type Data_detail_umkm_kontainer struct {
		ID         int       `json:"id_wisata"`
		Nama       string    `json:"nama_wisata"`
		Konten     string    `json:"konten_wisata"`
		Kategori   string    `json:"kategori_wisata"`
		DesaID     int       `json:"desa_id"`
		NoTelp     string    `json:"no_telp"`
		CreatedAt  time.Time `json:"created_at"`
		CreatedBy  string    `json:"created_by"`
		UpdatedAt  time.Time `json:"updated_at"`
		UpdatedBy  string    `json:"updated_by"`
		LinkLokasi string    `json:"link_lokasi"`
		Alamat     string    `json:"alamat"`
		FotoWisata string    `json:"foto_wisata"`
		IDKATEGORI int       `json:"id_kategori"`
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
		a.id_wisata      ,
		a.nama_wisata    ,
		a.konten_wisata  ,
		b.nama_kategori	 ,
		a.desa_id        ,
		a.no_telp        ,
		a.created_at     ,
		a.created_by     ,
		a.update_at      ,
		a.updated_by     ,
		a.link_lokasi    ,
		a.alamat         ,
		a.foto_wisata    ,
		a.idkategoriwisata
	from dev.wisata a, dev.kategoriwisata b where id_wisata = $1 and a.idkategoriwisata = b.idkategoriwisata

	`

	err = tx.QueryRow(ctx, query, id).Scan(
		&ambil.ID,
		&ambil.Nama,
		&ambil.Konten,
		&ambil.Kategori,
		&ambil.DesaID,
		&ambil.NoTelp,
		&ambil.CreatedAt,
		&ambil.CreatedBy,
		&ambil.UpdatedAt,
		&ambil.UpdatedBy,
		&ambil.LinkLokasi,
		&ambil.Alamat,
		&ambil.FotoWisata,
		&ambil.IDKATEGORI,
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
