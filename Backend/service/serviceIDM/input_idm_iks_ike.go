package serviceidm

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Input(c *gin.Context) {
	type IDM struct {
		Tahun      string `json:"tahun"`
		Indikator  string `json:"indikator"`
		Skor       string `json:"skor"`
		Keterangan string `json:"keterangan"`
		Kegiatan   string `json:"kegiatan"`
		Nilai      string `json:"nilai"`
		Pusat      string `json:"pusat"`
		Provinsi   string `json:"provinsi"`
		Kabupaten  string `json:"kabupaten"`
		Desa       string `json:"desa"`
		CSR        string `json:"csr"`
		Lainya     string `json:"lainya"`
	}

	type IKS struct {
		Tahun      string `json:"tahun"`
		Indikator  string `json:"indikator"`
		Skor       string `json:"skor"`
		Keterangan string `json:"keterangan"`
		Kegiatan   string `json:"kegiatan"`
		Nilai      string `json:"nilai"`
		Pusat      string `json:"pusat"`
		Provinsi   string `json:"provinsi"`
		Kabupaten  string `json:"kabupaten"`
		Desa       string `json:"desa"`
		CSR        string `json:"csr"`
		Lainya     string `json:"lainya"`
	}

	type IKE struct {
		Tahun      string `json:"tahun"`
		Indikator  string `json:"indikator"`
		Skor       string `json:"skor"`
		Keterangan string `json:"keterangan"`
		Kegiatan   string `json:"kegiatan"`
		Nilai      string `json:"nilai"`
		Pusat      string `json:"pusat"`
		Provinsi   string `json:"provinsi"`
		Kabupaten  string `json:"kabupaten"`
		Desa       string `json:"desa"`
		CSR        string `json:"csr"`
		Lainya     string `json:"lainya"`
	}

	type Data struct {
		IDMinput []IDM `json:"data_idm"`
		IKSinput []IKS `json:"data_iks"`
		IKEinput []IKE `json:"data_ike"`
	}

	type Input_data struct {
		DesaID string `json:"desa_id"`
		Data   []Data `json:"data"`
	}

	var input Input_data

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

	year := input.Data[0].IDMinput[0].Tahun

	if len(input.Data) != 0 {

		for _, v := range input.Data[0].IDMinput {

			var id_next, id_now int

			cek_berita := `
				SELECT MAX(id) AS last_id
				FROM dev.idm;	
				`
			err = tx.QueryRow(ctx, cek_berita).Scan(&id_now)

			if err != nil {
				id_now = 0
			}

			id_next = id_now + 1

			query := `
					INSERT INTO dev.idm (
						id,
						desa_id, 
						tahun, 
						indikator, 
						skor, 
						keterangan, 
						kegiatan,
						nilai,
						pusat,
						provinsi,
						kabupaten,
						desa,
						csr,
						lainya
					)
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
				`

			_, err = tx.Exec(ctx, query, id_next, input.DesaID, year, v.Indikator,
				v.Skor, v.Keterangan, v.Kegiatan, v.Nilai, v.Pusat, v.Provinsi, v.Kabupaten,
				v.Desa, v.CSR, v.Lainya)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
				return
			}
		}

		for _, v := range input.Data[0].IKSinput {
			var id_next, id_now int

			cek_berita := `
				SELECT MAX(id) AS last_id
				FROM dev.iks;	
				`
			err = tx.QueryRow(ctx, cek_berita).Scan(&id_now)

			if err != nil {
				id_now = 0
			}

			id_next = id_now + 1

			query := `
					INSERT INTO dev.iks (
						id,
						desa_id, 
						tahun, 
						indikator, 
						skor, 
						keterangan, 
						kegiatan,
						nilai,
						pusat,
						provinsi,
						kabupaten,
						desa,
						csr,
						lainya
					)
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
				`

			_, err = tx.Exec(ctx, query, id_next, input.DesaID, year, v.Indikator,
				v.Skor, v.Keterangan, v.Kegiatan, v.Nilai, v.Pusat, v.Provinsi, v.Kabupaten,
				v.Desa, v.CSR, v.Lainya)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
				return
			}
		}

		for _, v := range input.Data[0].IKEinput {
			var id_next, id_now int

			cek_berita := `
				SELECT MAX(id) AS last_id
				FROM dev.ike;	
				`
			err = tx.QueryRow(ctx, cek_berita).Scan(&id_now)

			if err != nil {
				id_now = 0
			}

			id_next = id_now + 1

			query := `
					INSERT INTO dev.ike (
						id,
						desa_id, 
						tahun, 
						indikator, 
						skor, 
						keterangan, 
						kegiatan,
						nilai,
						pusat,
						provinsi,
						kabupaten,
						desa,
						csr,
						lainya
					)
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
				`

			_, err = tx.Exec(ctx, query, id_next, input.DesaID, year, v.Indikator,
				v.Skor, v.Keterangan, v.Kegiatan, v.Nilai, v.Pusat, v.Provinsi, v.Kabupaten,
				v.Desa, v.CSR, v.Lainya)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
				return
			}
		}

	}

	err = tx.Commit(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Error committing transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Upload Berhasil"})
}
