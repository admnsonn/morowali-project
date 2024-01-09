package sevicewargadesabyamin

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Warga_desa_by_admin(c *gin.Context) {
	type Data_list_warga_kontainer struct {
		IDPengguna     int    `json:"id_pengguna"`
		NIK            int    `json:"nik"`
		NamaLengkap    string `json:"nama_lengkap"`
		AlamatPengguna string `json:"alamat_pengguna"`
		NoTelp         string `json:"no_telp"`
		KK             int    `json:"kk"`
		JenisKelamin   string `json:"jenis_kelamin"`
	}

	type Request struct {
		IDdesa           int    `json:"id_desa"`
		RTPengguna       string `json:"rt"`
		RWPengguna       string `json:"rw"`
		TingkatPedidikan string `json:"tingkat_pendidikan"`
		NIK              string `json:"nik"`
	}

	var input Request
	var Tampung_list_warga []Data_list_warga_kontainer
	var rt, rw, kode_pos string

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

	if input.NIK != "" {

		var ambil3 Data_list_warga_kontainer
		// Filter pecarian NIK
		query_pencarian := `
			Select 

			id_pengguna          ,
			nik                  ,
			kk					 ,
			jenis_kelamin		 ,
			nama_lengkap         ,
			alamat_pengguna      ,
			rt                   ,
			rw                   ,
			kode_pos             ,
			no_telp 
			from dev.pengguna where nik = $1
		`

		err = tx.QueryRow(ctx, query_pencarian, input.NIK).Scan(
			&ambil3.IDPengguna,
			&ambil3.NIK,
			&ambil3.KK,
			&ambil3.JenisKelamin,
			&ambil3.NamaLengkap,
			&ambil3.AlamatPengguna,
			&rt,
			&rw,
			&kode_pos,
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

			Tampung_list_warga[0].AlamatPengguna = Tampung_list_warga[0].AlamatPengguna + ", RT: " + rt + ", RW: " + rw + ", Kode Post: " + kode_pos

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

	query := ``

	if input.TingkatPedidikan != "" {
		if input.TingkatPedidikan == "SD" {
			query = query + `
			SELECT 
			c.id_pengguna          ,
			c.nik                  ,
			c.kk					 ,
			c.jenis_kelamin		 ,
			c.nama_lengkap         ,
			c.alamat_pengguna      ,
			c.rt                   ,
			c.rw                   ,
			c.kode_pos             ,
			c.no_telp
			FROM dev.pendidikan a, dev.pengguna_pendidikan b, dev.pengguna c
			WHERE 
			c.id_pengguna = b.pengguna_id_pengguna 
			and a.id_pendidikan  = b.pendidikan_pengguna_id 
			and (tingkat_pendidikan = 'SD')
			AND pengguna_id NOT IN (
				SELECT pengguna_id
				FROM dev.pendidikan
				WHERE tingkat_pendidikan IN ('SMP', 'SMA', 'D1', 'D2', 'D3', 'D4', 'S1', 'S2', 'S3')
			)
			and c.desa_id = $1
			and c.role_id = 2
			`

		} else if input.TingkatPedidikan == "SMP" {
			query = query + `
			SELECT 
			c.id_pengguna          ,
			c.nik                  ,
			c.kk					 ,
			c.jenis_kelamin		 ,
			c.nama_lengkap         ,
			c.alamat_pengguna      ,
			c.rt                   ,
			c.rw                   ,
			c.kode_pos             ,
			c.no_telp
			FROM dev.pendidikan a, dev.pengguna_pendidikan b, dev.pengguna c
			WHERE 
			c.id_pengguna = b.pengguna_id_pengguna 
			and a.id_pendidikan  = b.pendidikan_pengguna_id 
			and (tingkat_pendidikan = 'SMP')
			AND pengguna_id NOT IN (
				SELECT pengguna_id
				FROM dev.pendidikan
				WHERE tingkat_pendidikan IN ('SMK', 'SMA', 'D1', 'D2', 'D3', 'D4', 'S1', 'S2', 'S3')
			)
			and c.desa_id = $1
			and c.role_id = 2
			`

		} else if input.TingkatPedidikan == "SMK" || input.TingkatPedidikan == "SMA" {
			query = query + `
			SELECT 
			c.id_pengguna          ,
			c.nik                  ,
			c.kk					 ,
			c.jenis_kelamin		 ,
			c.nama_lengkap         ,
			c.alamat_pengguna      ,
			c.rt                   ,
			c.rw                   ,
			c.kode_pos             ,
			c.no_telp
			FROM dev.pendidikan a, dev.pengguna_pendidikan b, dev.pengguna c
			WHERE 
			c.id_pengguna = b.pengguna_id_pengguna 
			and a.id_pendidikan  = b.pendidikan_pengguna_id 
			and (tingkat_pendidikan = 'SMK' or tingkat_pendidikan = 'SMA')
			AND pengguna_id NOT IN (
				SELECT pengguna_id
				FROM dev.pendidikan
				WHERE tingkat_pendidikan IN ( 'D1', 'D2', 'D3', 'D4', 'S1', 'S2', 'S3')
			)
			and c.desa_id = $1
			and c.role_id = 2
			`

		} else if input.TingkatPedidikan == "D1" {
			query = query + `
			SELECT 
			c.id_pengguna          ,
			c.nik                  ,
			c.kk					 ,
			c.jenis_kelamin		 ,
			c.nama_lengkap         ,
			c.alamat_pengguna      ,
			c.rt                   ,
			c.rw                   ,
			c.kode_pos             ,
			c.no_telp
			FROM dev.pendidikan a, dev.pengguna_pendidikan b, dev.pengguna c
			WHERE 
			c.id_pengguna = b.pengguna_id_pengguna 
			and a.id_pendidikan  = b.pendidikan_pengguna_id 
			and (tingkat_pendidikan = 'D1')
			AND pengguna_id NOT IN (
				SELECT pengguna_id
				FROM dev.pendidikan
				WHERE tingkat_pendidikan IN ('D2', 'D3', 'D4', 'S1', 'S2', 'S3')
			)
			and c.desa_id = $1
			and c.role_id = 2
			`

		} else if input.TingkatPedidikan == "D2" {
			query = query + `
			SELECT 
			c.id_pengguna          ,
			c.nik                  ,
			c.kk					 ,
			c.jenis_kelamin		 ,
			c.nama_lengkap         ,
			c.alamat_pengguna      ,
			c.rt                   ,
			c.rw                   ,
			c.kode_pos             ,
			c.no_telp
			FROM dev.pendidikan a, dev.pengguna_pendidikan b, dev.pengguna c
			WHERE 
			c.id_pengguna = b.pengguna_id_pengguna 
			and a.id_pendidikan  = b.pendidikan_pengguna_id 
			and (tingkat_pendidikan = 'D2')
			AND pengguna_id NOT IN (
				SELECT pengguna_id
				FROM dev.pendidikan
				WHERE tingkat_pendidikan IN ('D3', 'D4', 'S1', 'S2', 'S3')
			)
			and c.desa_id = $1
			and c.role_id = 2
			`

		} else if input.TingkatPedidikan == "D3" {
			query = query + `
			SELECT 
			c.id_pengguna          ,
			c.nik                  ,
			c.kk					 ,
			c.jenis_kelamin		 ,
			c.nama_lengkap         ,
			c.alamat_pengguna      ,
			c.rt                   ,
			c.rw                   ,
			c.kode_pos             ,
			c.no_telp
			FROM dev.pendidikan a, dev.pengguna_pendidikan b, dev.pengguna c
			WHERE 
			c.id_pengguna = b.pengguna_id_pengguna 
			and a.id_pendidikan  = b.pendidikan_pengguna_id 
			and (tingkat_pendidikan = 'D3')
			AND pengguna_id NOT IN (
				SELECT pengguna_id
				FROM dev.pendidikan
				WHERE tingkat_pendidikan IN ('D4', 'S1', 'S2', 'S3')
			)
			and c.desa_id = $1
			and c.role_id = 2
			`

		} else if input.TingkatPedidikan == "S1" || input.TingkatPedidikan == "D4" {
			fmt.Println("Masuk ke bagian S1 ATAU D3")

			query = query + `
			SELECT
			c.id_pengguna          ,
			c.nik                  ,
			c.kk					 ,
			c.jenis_kelamin		 ,
			c.nama_lengkap         ,
			c.alamat_pengguna      ,
			c.rt                   ,
			c.rw                   ,
			c.kode_pos             ,
			c.no_telp
			FROM dev.pendidikan a, dev.pengguna_pendidikan b, dev.pengguna c
			WHERE 
			c.id_pengguna = b.pengguna_id_pengguna 
			and a.id_pendidikan  = b.pendidikan_pengguna_id 
			and (tingkat_pendidikan = 'D4' or tingkat_pendidikan = 'S1')
			AND pengguna_id NOT IN (
				SELECT pengguna_id
				FROM dev.pendidikan
				WHERE tingkat_pendidikan IN ('S2', 'S3')
			)
			and c.desa_id = $1
			and c.role_id = 2
			`

		} else if input.TingkatPedidikan == "S2" {
			query = query + `
			SELECT 
			c.id_pengguna          ,
			c.nik                  ,
			c.kk					 ,
			c.jenis_kelamin		 ,
			c.nama_lengkap         ,
			c.alamat_pengguna      ,
			c.rt                   ,
			c.rw                   ,
			c.kode_pos             ,
			c.no_telp
			FROM dev.pendidikan a, dev.pengguna_pendidikan b, dev.pengguna c
			WHERE 
			c.id_pengguna = b.pengguna_id_pengguna 
			and a.id_pendidikan  = b.pendidikan_pengguna_id 
			and (tingkat_pendidikan = 'S2')
			AND pengguna_id NOT IN (
				SELECT pengguna_id
				FROM dev.pendidikan
				WHERE tingkat_pendidikan IN ('S3')
			)
			and c.desa_id = $1
			and c.role_id = 2
			`

		} else {
			query = query + `
			SELECT 
			c.id_pengguna          ,
			c.nik                  ,
			c.kk					 ,
			c.jenis_kelamin		 ,
			c.nama_lengkap         ,
			c.alamat_pengguna      ,
			c.rt                   ,
			c.rw                   ,
			c.kode_pos             ,
			c.no_telp
			FROM dev.pendidikan a, dev.pengguna_pendidikan b, dev.pengguna c
			WHERE 
			c.id_pengguna = b.pengguna_id_pengguna 
			and a.id_pendidikan  = b.pendidikan_pengguna_id 
			and (tingkat_pendidikan = 'S3')
			and c.desa_id = $1
			and c.role_id = 2
			`

		}

		if input.RWPengguna != "" {
			if input.RTPengguna != "" {

			} else {

				query = query + `
					and c.rw = $2 
				`

				fmt.Println(query)

				row, err := tx.Query(ctx, query, input.IDdesa, input.RWPengguna)

				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
					err = tx.Commit(ctx)
					if err != nil {
						panic(err.Error())
					}
					return
				}

				defer row.Close()

				for row.Next() {
					var ambil Data_list_warga_kontainer

					err := row.Scan(
						&ambil.IDPengguna,
						&ambil.NIK,
						&ambil.KK,
						&ambil.JenisKelamin,
						&ambil.NamaLengkap,
						&ambil.AlamatPengguna,
						&rt,
						&rw,
						&kode_pos,
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

					Tampung_list_warga = append(Tampung_list_warga, ambil)
				}

			}

			// bedasarkan pendidikan, rw

		}

	} else {
		if input.RWPengguna != "" {
			if input.RTPengguna != "" {
				query = `
				select
				id_pengguna          ,
				nik                  ,
				kk					 ,
				jenis_kelamin		 ,
				nama_lengkap         ,
				alamat_pengguna      ,
				rt                   ,
				rw                   ,
				kode_pos             ,
				no_telp
				from dev.pengguna
				where desa_id = $1
				and role_id = 2
				and rw = $2 
				and rt = $3 
				`

				row, err := tx.Query(ctx, query, input.IDdesa, input.RWPengguna, input.RTPengguna)

				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
					err = tx.Commit(ctx)
					if err != nil {
						panic(err.Error())
					}
					return
				}

				defer row.Close()

				for row.Next() {
					var ambil Data_list_warga_kontainer

					err := row.Scan(
						&ambil.IDPengguna,
						&ambil.NIK,
						&ambil.KK,
						&ambil.JenisKelamin,
						&ambil.NamaLengkap,
						&ambil.AlamatPengguna,
						&rt,
						&rw,
						&kode_pos,
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

					Tampung_list_warga = append(Tampung_list_warga, ambil)
				}

				// Bedasarkan RT DAN RW
			} else {
				query = `
				select
				id_pengguna          ,
				nik                  ,
				kk					 ,
				jenis_kelamin		 ,
				nama_lengkap         ,
				alamat_pengguna      ,
				rt                   ,
				rw                   ,
				kode_pos             ,
				no_telp
				from dev.pengguna
				where desa_id = $1
				and role_id = 2
				and rw = $2 

				`

				row, err := tx.Query(ctx, query, input.IDdesa, input.RWPengguna)

				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
					err = tx.Commit(ctx)
					if err != nil {
						panic(err.Error())
					}
					return
				}

				defer row.Close()

				for row.Next() {
					var ambil Data_list_warga_kontainer

					err := row.Scan(
						&ambil.IDPengguna,
						&ambil.NIK,
						&ambil.KK,
						&ambil.JenisKelamin,
						&ambil.NamaLengkap,
						&ambil.AlamatPengguna,
						&rt,
						&rw,
						&kode_pos,
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

					Tampung_list_warga = append(Tampung_list_warga, ambil)
				}
			}

			// Bedasarkan RW
		} else {
			query = `
				select
					id_pengguna          ,
					nik                  ,
					kk					 ,
					jenis_kelamin		 ,
					nama_lengkap         ,
					alamat_pengguna      ,
					rt                   ,
					rw                   ,
					kode_pos             ,
					no_telp
				from dev.pengguna
				where desa_id = $1
				and role_id = 2
			`

			row, err := tx.Query(ctx, query, input.IDdesa)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
				err = tx.Commit(ctx)
				if err != nil {
					panic(err.Error())
				}
				return
			}

			defer row.Close()

			for row.Next() {
				var ambil Data_list_warga_kontainer

				err := row.Scan(
					&ambil.IDPengguna,
					&ambil.NIK,
					&ambil.KK,
					&ambil.JenisKelamin,
					&ambil.NamaLengkap,
					&ambil.AlamatPengguna,
					&rt,
					&rw,
					&kode_pos,
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

				Tampung_list_warga = append(Tampung_list_warga, ambil)
			}
		}

	}

	if len(Tampung_list_warga) > 0 {

		Tampung_list_warga[0].AlamatPengguna = Tampung_list_warga[0].AlamatPengguna + ", RT: " + rt + ", RW: " + rw + ", Kode Post: " + kode_pos

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
