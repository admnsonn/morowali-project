package sevicewargadesabyamin

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Get_detail_warga_by_admin(c *gin.Context) {
	type Data_Detail struct {
		IDPengguna        int       `json:"id_pengguna"`
		RolePengguna      string    `json:"role_pengguna"`
		TanggalLahir      time.Time `json:"tanggal_lahir"`
		TempatLahir       string    `json:"tempat_lahir"`
		JenisKelamin      string    `json:"jenis_kelamin"`
		StatusPerkawinan  string    `json:"status_perkawinan"`
		Profesi           string    `json:"profesi"`
		NIK               int       `json:"nik"`
		KK                int       `json:"kk"`
		Umur              int       `json:"umur"`
		KategoriFinancial string    `json:"kategori_financial"`
		Agama             string    `json:"agama"`
		Desa              string    `json:"desa"`
		NamaLengkap       string    `json:"nama_lengkap"`
		FotoPengguna      string    `json:"foto_pengguna"`
		AlamatPengguna    string    `json:"alamat_pengguna"`
		RT                string    `json:"rt"`
		RW                string    `json:"rw"`
		KodePos           string    `json:"kode_pos"`
		NoTelp            string    `json:"no_telp"`
	}

	id := c.Param("id")

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}

	query := `

		Select 
			a.id_pengguna,
			a.role_pengguna,
			a.tanggal_lahir,
			a.tempat_lahir,
			a.jenis_kelamin,
			a.status_perkawinan,
			a.profesi,
			a.nik,
			a.kk,
			a.umur,
			b.kategori_financial,
			a.agama,
			a.desa,
			a.nama_lengkap,
			a.foto_pengguna,
			a.alamat_pengguna,
			a.rt,
			a.rw,
			a.kode_pos,
			a.no_telp
		from dev.pengguna a, dev.kategori_financial b 
		where b.id_kategori_financial = a.kategori_financial_id 
		and a.id_pengguna = $1

	`

	var Tampung_detail_warga []Data_Detail
	var ambil Data_Detail

	err = tx.QueryRow(ctx, query, id).Scan(
		&ambil.IDPengguna,
		&ambil.RolePengguna,
		&ambil.TanggalLahir,
		&ambil.TempatLahir,
		&ambil.JenisKelamin,
		&ambil.StatusPerkawinan,
		&ambil.Profesi,
		&ambil.NIK,
		&ambil.KK,
		&ambil.Umur,
		&ambil.KategoriFinancial,
		&ambil.Agama,
		&ambil.Desa,
		&ambil.NamaLengkap,
		&ambil.FotoPengguna,
		&ambil.AlamatPengguna,
		&ambil.RT,
		&ambil.RW,
		&ambil.KodePos,
		&ambil.NoTelp,
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
