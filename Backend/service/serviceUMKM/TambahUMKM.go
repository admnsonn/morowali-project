package serviceumkm

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Tambah_UMKM(c *gin.Context) {
	type UMKM struct {
		Nama         string `json:"nama_umkm"`
		Konten       string `json:"konten_umkm"`
		KategoriID   int    `json:"kategori_umkm_id"`
		Foto         string `json:"foto_umkm"`
		DesaID       int    `json:"desa_id"`
		NomorTelepon string `json:"no_telp_umkm"`
		Alamat       string `json:"alamat_umkm"`
	}

	var input UMKM

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
	defer tx.Rollback(context.Background())

	var id_next, id_now int

	cek_berita := `
	SELECT MAX(id_umkm) AS last_id
	FROM dev.umkm;
	`
	err = tx.QueryRow(ctx, cek_berita).Scan(&id_now)

	if err != nil {
		id_next = id_now + 1
	} else {
		id_next = id_now + 1
	}

	query := `
		INSERT INTO dev.umkm (
			id_umkm ,nama_umkm , konten_umkm , kategori_umkm_id , foto_umkm , desa_id , no_telp_umkm , alamat,
			created_at,
			created_by,
			update_at,
			updated_by
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`

	_, err = tx.Exec(ctx, query, id_next, input.Nama, input.Konten, input.KategoriID, input.Foto, input.DesaID, input.NomorTelepon, input.Alamat,
		time.Now(), "admin", time.Now(), "admin")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Error committing transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Berita berhasil ditambahkan"})
}
