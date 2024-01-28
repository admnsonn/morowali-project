package sevicewargadesabyamin

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Create_warga(c *gin.Context) {
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
		DesaID          string `json:"id_desa"`
		Financial       string `json:"kategori_financial_id"`
		KodePos         string `json:"kode_pos"`
		Kewarganegaraan string `json:"kewarganegaraan"`
	}

	var input Request

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

	var hitung2 int

	cek_nik := `
		SELECT COUNT(*)
		FROM dev.pengguna
		WHERE nik = $1;
		`
	err = tx.QueryRow(ctx, cek_nik, input.NIK).Scan(&hitung2)
	if err != nil {
		log.Fatal(err)
	}

	if hitung2 != 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "NIK Sudah Terdaftar !",
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}

	password := "desabahomoleo123"

	hasher := md5.New()
	hasher.Write([]byte(password))
	hashedInputPassword := hex.EncodeToString(hasher.Sum(nil))
	fmt.Println(hashedInputPassword)

	fmt.Println("ini hasher : ", hasher)
	fmt.Println("ini hashedInputPassword : ", hashedInputPassword)

	query := `
		INSERT INTO dev.pengguna (
			nama_lengkap,
			nik,
			kk,
			tempat_lahir,
			tanggal_lahir,
			jk_id ,
			alamat_pengguna,
			agama_id ,
			status_perkawinan_id ,
			profesi,
			no_telp,
			pendidikan_id ,
			foto_pengguna ,
			rw,
			rt,
			desa_id,
			kategori_financial_id ,
			role_id,
			kode_pos ,
			kewarganegaraan,
			created_at ,
			created_by ,
			update_at ,
			updated_by ,
			password
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25
		)
		RETURNING id_pengguna;
		`

	var newID int

	err = tx.QueryRow(ctx, query,
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
		input.DesaID,
		input.Financial,
		"2",
		input.KodePos,
		input.Kewarganegaraan,
		time.Now(),
		"Admin Sistem",
		time.Now(),
		"Admin Sistem",
		hashedInputPassword,
	).Scan(&newID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Query Insert Bermasalah",
		})
		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Data Gagal Ditambahkan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Data Berhasil Di Tambahkan",
	})

}
