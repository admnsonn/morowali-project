package sevicewargadesabyamin

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Get_detail_warga_by_admin(c *gin.Context) {

	type Data_Detail struct {
		NamaLengkap       string `json:"nama_lengkap"`
		Umur              string `json:"umur"`
		NIK               string `json:"nik"`
		KK                string `json:"kk"`
		TempatLahir       string `json:"tempat_lahir"`
		TanggalLahir      string `json:"tanggal_lahir"`
		JenisKelamin      string `json:"jenis_kelamin"`
		NamaAgama         string `json:"nama_agama"`
		StatusPerkawinan  string `json:"status_perkawinan"`
		Profesi           string `json:"profesi"`
		NoTelp            string `json:"no_telp"`
		TingkatPendidikan string `json:"tingkat_pendidikan"`
		FotoPengguna      string `json:"foto_pengguna"`
		AlamatPengguna    string `json:"alamat_pengguna"`
		RT                string `json:"rt"`
		RW                string `json:"rw"`
		NamaDesa          string `json:"nama_desa"`
		NamaKecamatan     string `json:"nama_kecamatan"`
	}

	id := c.Param("id")

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}

	query := `
	select 
		a.nama_lengkap ,
		CAST(DATE_PART('year', AGE(NOW(), a.tanggal_lahir)) AS VARCHAR) AS umur,
		a.nik,
		a.kk ,
		a.tempat_lahir ,
		TO_CHAR(a.tanggal_lahir, 'YYYY-MM-DD') AS tahun_lahir,
		b.jenis_kelamin ,
		d.nama_agama ,
		e.status_perkawinan ,
		a.profesi ,
		a.no_telp ,
		c.tingkat_pendidikan ,
		a.foto_pengguna ,
		a.alamat_pengguna ,
		a.rt ,
		a.rw ,
		f.nama_desa ,
		g.nama_kecamatan 
	from dev.pengguna a, dev.jenis_kelamin b, dev.pendidikan c, dev.agama d, dev.status_perkawinan e, dev.desa f, dev.kecamatan g
	where a.id_pengguna = c.pengguna_id 
	and a.jk_id = b.id_jk 
	and a.agama_id  = d.id_agama  
	and a.status_perkawinan_id  = e.id_status  
	and a.desa_id = f.id_desa 
	and f.kecamatan_id  = g.id_kecamatan 
	and a.id_pengguna = $1
	`

	var Tampung_detail_warga []Data_Detail
	var ambil Data_Detail

	err = tx.QueryRow(ctx, query, id).Scan(
		&ambil.NamaLengkap,
		&ambil.Umur,
		&ambil.NIK,
		&ambil.KK,
		&ambil.TempatLahir,
		&ambil.TanggalLahir,
		&ambil.JenisKelamin,
		&ambil.NamaAgama,
		&ambil.StatusPerkawinan,
		&ambil.Profesi,
		&ambil.NoTelp,
		&ambil.TingkatPendidikan,
		&ambil.FotoPengguna,
		&ambil.AlamatPengguna,
		&ambil.RT,
		&ambil.RW,
		&ambil.NamaDesa,
		&ambil.NamaKecamatan,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		err = tx.Commit(ctx)
		if err != nil {
			panic(err.Error())
		}
		return
	}

	Tampung_detail_warga = append(Tampung_detail_warga, ambil)

	if len(Tampung_detail_warga) != 0 {

		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Data ditemukan",
			"data":    Tampung_detail_warga,
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return

	} else {

		c.JSON(http.StatusOK, gin.H{"status": false, "message": "Data tidak ditemukan", "data": []Data_Detail{}})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return

	}
}
