package sevicewargadesabyamin

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Warga_desa_by_admin(c *gin.Context) {
	type Data_list_warga_kontainer struct {
		IDPengguna     int    `json:"id_pengguna"`
		NIK            string `json:"nik"`
		AlamatPengguna string `json:"alamat_pengguna"`
		NamaLengkap    string `json:"nama_lengkap"`
		NoTelp         string `json:"no_telp"`
		KK             string `json:"kk"`
		JenisKelamin   string `json:"jenis_kelamin"`
		Umur           string `json:"umur"`
	}

	type Request struct {
		IDdesa   string `json:"id_desa"`
		NIK      string `json:"nik"`
		KK       string `json:"kk"`
		Umur     string `json:"umur"`
		Fullname string `json:"nama_lengkap"`
	}

	var input Request
	var Tampung_list_warga []Data_list_warga_kontainer
	if c.GetHeader("content-type") == "application/x-www-form-urlencoded" || c.GetHeader("content-type") == "application/x-www-form-urlencoded; charset=utf-8" {

		if err := c.Bind(&input); err != nil {
			fmt.Println("MASUK SINI")
			return
		}

	} else {

		if err := c.BindJSON(&input); err != nil {
			fmt.Println("MASUK SINI")
			return
		}

	}

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}

	if input.NIK != "" {

		var ambil3 Data_list_warga_kontainer
		// Filter pecarian NIK
		query_pencarian := `
		Select a.id_pengguna,
		a.niK,
		a.kk,
		b.jenis_kelamin,
		a.alamat_pengguna,
		a.nama_lengkap,
		CAST(DATE_PART('year', AGE(NOW(), a.tanggal_lahir)) AS VARCHAR) AS umur,
		a.no_telp 
	FROM
		dev.pengguna a, dev.jenis_kelamin b
	where 
		a.jk_id = b.id_jk and a.role_id  = 2 and nik = $1
		`

		err = tx.QueryRow(ctx, query_pencarian, input.NIK).Scan(
			&ambil3.IDPengguna,
			&ambil3.NIK,
			&ambil3.KK,
			&ambil3.JenisKelamin,
			&ambil3.AlamatPengguna,
			&ambil3.NamaLengkap,
			&ambil3.Umur,
			&ambil3.NoTelp,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			err = tx.Commit(ctx)
			if err != nil {
				panic(err.Error())
			}
			return
		}

		Tampung_list_warga = append(Tampung_list_warga, ambil3)

		if len(Tampung_list_warga) > 0 {

			c.JSON(http.StatusOK, gin.H{
				"status":  true,
				"message": "Data Warga Desa tersedia",
				"data":    Tampung_list_warga,
			})
			err = tx.Commit(ctx)
			if err != nil {
				fmt.Println(err.Error())
			}
			return
		} else {

			c.JSON(http.StatusOK, gin.H{
				"status":  true,
				"message": "Data Warga Desa tidak tersedia !",
				"data":    []Data_list_warga_kontainer{},
			})
			err = tx.Commit(ctx)
			if err != nil {
				fmt.Println(err.Error())
			}
			return

		}

	}

	if input.KK != "" {

		fmt.Println("INI KK NYA :", input.KK)

		var ambil3 Data_list_warga_kontainer
		// Filter pecarian NIK
		query_pencarian := `
		Select a.id_pengguna,
		a.niK,
		a.kk,
		b.jenis_kelamin,
		a.alamat_pengguna,
		a.nama_lengkap,
		CAST(DATE_PART('year', AGE(NOW(), a.tanggal_lahir)) AS VARCHAR) AS umur,
		a.no_telp 
	FROM
		dev.pengguna a, dev.jenis_kelamin b
		where 
			a.jk_id = b.id_jk and a.role_id  = 2 and kk = $1
		`

		err = tx.QueryRow(ctx, query_pencarian, input.KK).Scan(
			&ambil3.IDPengguna,
			&ambil3.NIK,
			&ambil3.KK,
			&ambil3.JenisKelamin,
			&ambil3.AlamatPengguna,
			&ambil3.NamaLengkap,
			&ambil3.Umur,
			&ambil3.NoTelp,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			err = tx.Commit(ctx)
			if err != nil {
				panic(err.Error())
			}
			return
		}

		Tampung_list_warga = append(Tampung_list_warga, ambil3)
	} else if input.Umur != "" {

		fmt.Println("INI KK NYA :", input.KK)

		var ambil3 Data_list_warga_kontainer
		// Filter pecarian NIK
		query_pencarian := `
		Select a.id_pengguna,
		a.niK,
		a.kk,
		b.jenis_kelamin,
		a.alamat_pengguna,
		a.nama_lengkap,
		CAST(DATE_PART('year', AGE(NOW(), a.tanggal_lahir)) AS VARCHAR) AS umur,
		a.no_telp 
	FROM
		dev.pengguna a, dev.jenis_kelamin b
	where 
		a.jk_id = b.id_jk and a.role_id  = 2 and
				DATE_PART('year', AGE(NOW(), tanggal_lahir)) = $1;
		`

		err = tx.QueryRow(ctx, query_pencarian, input.Umur).Scan(
			&ambil3.IDPengguna,
			&ambil3.NIK,
			&ambil3.KK,
			&ambil3.JenisKelamin,
			&ambil3.AlamatPengguna,
			&ambil3.NamaLengkap,
			&ambil3.Umur,
			&ambil3.NoTelp,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			err = tx.Commit(ctx)
			if err != nil {
				panic(err.Error())
			}
			return
		}

		Tampung_list_warga = append(Tampung_list_warga, ambil3)
	} else if input.Fullname != "" {

		// fmt.Println("INI KK NYA :", input.KK)

		var ambil3 Data_list_warga_kontainer
		// Filter pecarian NIK
		query_pencarian := `
		Select a.id_pengguna,
		a.niK,
		a.kk,
		b.jenis_kelamin,
		a.alamat_pengguna,
		a.nama_lengkap,
		CAST(DATE_PART('year', AGE(NOW(), a.tanggal_lahir)) AS VARCHAR) AS umur,
		a.no_telp 
	FROM
		dev.pengguna a, dev.jenis_kelamin b
		where 
			a.jk_id = b.id_jk 
			and a.role_id  = 2 and
				UPPER(nama_lengkap) = UPPER($1);
		`

		err = tx.QueryRow(ctx, query_pencarian, input.Fullname).Scan(
			&ambil3.IDPengguna,
			&ambil3.NIK,
			&ambil3.KK,
			&ambil3.JenisKelamin,
			&ambil3.AlamatPengguna,
			&ambil3.NamaLengkap,
			&ambil3.Umur,
			&ambil3.NoTelp,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			err = tx.Commit(ctx)
			if err != nil {
				panic(err.Error())
			}
			return
		}

		Tampung_list_warga = append(Tampung_list_warga, ambil3)
	} else {
		var ambil4 Data_list_warga_kontainer

		query_pencarian := `
		Select a.id_pengguna,
			a.niK,
			a.kk,
			b.jenis_kelamin,
			a.alamat_pengguna,
			a.nama_lengkap,
			CAST(DATE_PART('year', AGE(NOW(), a.tanggal_lahir)) AS VARCHAR) AS umur,
			a.no_telp 
		FROM
			dev.pengguna a, dev.jenis_kelamin b
			where 
				a.jk_id = b.id_jk 
				and a.role_id  = 2
				and a.desa_id = $1;
		`

		rowwarga, err := tx.Query(ctx, query_pencarian, input.IDdesa)

		if err != nil {
			log.Println("Error executing query:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			err = tx.Commit(ctx)
			if err != nil {
				panic(err.Error())
			}
			return
		}

		defer rowwarga.Close()

		for rowwarga.Next() {
			err := rowwarga.Scan(
				&ambil4.IDPengguna,
				&ambil4.NIK,
				&ambil4.KK,
				&ambil4.JenisKelamin,
				&ambil4.AlamatPengguna,
				&ambil4.NamaLengkap,
				&ambil4.Umur,
				&ambil4.NoTelp,
			)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
				err = tx.Commit(ctx)
				if err != nil {
					panic(err.Error())
				}
				return
			}

			Tampung_list_warga = append(Tampung_list_warga, ambil4)
		}
	}

	if len(Tampung_list_warga) > 0 {

		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Data Warga Desa tersedia",
			"data":    Tampung_list_warga,
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	} else {

		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Data Warga Desa tidak tersedia !",
			"data":    []Data_list_warga_kontainer{},
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return

	}

}
