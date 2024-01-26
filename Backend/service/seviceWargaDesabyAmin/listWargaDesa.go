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
		ID                    int    `json:"id_pengguna"`
		NIK                   string `json:"niK"`
		KK                    string `json:"kk"`
		JenisKelaminID        int    `json:"jk_id"`
		JenisKelamin          string `json:"jenis_kelamin"`
		AlamatPengguna        string `json:"alamat_pengguna"`
		NamaLengkap           string `json:"nama_lengkap"`
		Umur                  string `json:"umur"`
		NoTelp                string `json:"no_telp"`
		KategoriFinancial     int    `json:"kategori_financial_id"`
		KategoriFinancialName string `json:"kategori_financial"`
		AgamaID               int    `json:"agama_id"`
		NamaAgama             string `json:"nama_agama"`
		PendidikanID          int    `json:"pendidikan_id"`
		NamaPendidikan        string `json:"nama_pendidikan"`
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

	var ambil4 Data_list_warga_kontainer

	query_pencarian := `
		Select a.id_pengguna,
			a.niK,
			a.kk,
			a.jk_id ,
			b.jenis_kelamin,
			a.alamat_pengguna,
			a.nama_lengkap,
			CAST(DATE_PART('year', AGE(NOW(), a.tanggal_lahir)) AS VARCHAR) AS umur,
			a.no_telp ,
			a.kategori_financial_id ,
			c.kategori_financial  ,
			a.agama_id ,
			d.nama_agama ,
			a.pendidikan_id ,
			e.category_name 
		FROM
			dev.pengguna a, dev.jenis_kelamin b, dev.kategori_financial c , dev.agama d , dev.kategori_pendidikan e 
			where 
				a.jk_id = b.id_jk and 
				a.pendidikan_id  = e.id and 
				a.agama_id  = d.id_agama and 
				a.kategori_financial_id  = c.id_kategori_financial 
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
			&ambil4.ID,
			&ambil4.NIK,
			&ambil4.KK,
			&ambil4.JenisKelaminID,
			&ambil4.JenisKelamin,
			&ambil4.AlamatPengguna,
			&ambil4.NamaLengkap,
			&ambil4.Umur,
			&ambil4.NoTelp,
			&ambil4.KategoriFinancial,
			&ambil4.KategoriFinancialName,
			&ambil4.AgamaID,
			&ambil4.NamaAgama,
			&ambil4.PendidikanID,
			&ambil4.NamaPendidikan,
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
