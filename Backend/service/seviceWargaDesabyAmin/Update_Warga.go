package sevicewargadesabyamin

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Update_data_warga(c *gin.Context) {
	type Request struct {
		Fullname        string `json:"nama_lengkap"`
		NIK             string `json:"nik"`
		KK              string `json:"kk"`
		TempatLahir     string `json:"tempat_lahir"`
		TanggalLahir    string `json:"tanggal_lahir"`
		JenisKelamin    string `json:"jenis_kelamin_id"`
		Alamat          string `json:"alamat_pengguna"`
		Agama           string `json:"agama_id"`
		StatusPenikahan string `json:"status_perkawinan_id"`
		Profesi         string `json:"profesi"`
		NoTelp          string `json:"no_telp"`
		PendidikanLast  string `json:"pendidikan_terakhir_id"`
		FotoProfile     string `json:"foto_profile"`
		RW              string `json:"rw"`
		RT              string `json:"rt"`
		Financial       string `json:"kategori_financial_id"`
		KodePos         string `json:"kode_pos"`
		Kewarganegaraan string `json:"kewarganegaraan"`
		ID              string `json:"id_pengguna"`
	}

	var input Request
	// var id_pendidikan int

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

	query_update_pendidikan := `
	UPDATE dev.pengguna SET
		nama_lengkap = $1,
		nik = $2,
		kk = $3,
		tempat_lahir = $4,
		tanggal_lahir = $5,
		jk_id = $6,
		alamat_pengguna = $7,
		agama_id = $8,
		status_perkawinan_id = $9,
		profesi = $10,
		no_telp = $11,
		pendidikan_id = $12,
		foto_pengguna = $13,
		rw = $14,
		rt = $15,
		kategori_financial_id = $16,
		kode_pos = $17,
		kewarganegaraan = $18
	WHERE
    	id_pengguna = $19;
	`

	_, err = tx.Exec(ctx, query_update_pendidikan,
		input.Fullname,
		input.NIK,
		input.KK,
		input.TempatLahir,
		input.TanggalLahir,
		input.JenisKelamin,
		input.Alamat,
		input.Agama,
		input.StatusPenikahan,
		input.Profesi,
		input.NoTelp,
		input.PendidikanLast,
		input.FotoProfile,
		input.RW,
		input.RT,
		input.Financial,
		input.KodePos,
		input.Kewarganegaraan,
		input.ID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		err = tx.Commit(ctx)
		if err != nil {
			panic(err.Error())
		}
		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Data Gagal Di Update",
		})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Data Berhasil Di Update",
	})

}
