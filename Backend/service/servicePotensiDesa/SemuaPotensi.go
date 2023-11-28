package servicepotensidesa

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Semuapotensi_desa(c *gin.Context) {
	type Data_potensi_desa struct {
		ID        int    `json:"id_potensi"`
		Judul     string `json:"judul_potensi"`
		Deskripsi string `json:"deskripsi"`
		Foto      string `json:"foto_potensi_desa"`
		SubJudul  string `json:"sub_judul"`
	}

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}

	var id_desa int

	desa := `
		select id_desa from desa d where nama_desa = 'Morowali'
	`

	err = tx.QueryRow(ctx, desa).Scan(&id_desa)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		err = tx.Commit(ctx)
		if err != nil {
			panic(err.Error())
		}
		return
	}

	fmt.Print(id_desa)

	potensi_desa := `
	select 

	a.id_potensi       ,
	a.judul_potensi    ,
	a.deskripsi        ,
	a.foto_potensi_desa,
	a.sub_judul
		
	from potensi_desa a , desa b
	where a.desa_id = $1
	and a.desa_id = B.id_desa
	`

	row, err := tx.Query(ctx, potensi_desa, id_desa)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		err = tx.Commit(ctx)
		if err != nil {
			panic(err.Error())
		}
		return
	}

	defer row.Close()

	var Tampung_potensi_desa []Data_potensi_desa

	for row.Next() {
		var ambil Data_potensi_desa
		row.Scan(
			&ambil.ID,
			&ambil.Judul,
			&ambil.Deskripsi,
			&ambil.Foto,
			&ambil.SubJudul,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			err = tx.Commit(ctx)
			if err != nil {
				panic(err.Error())
			}
			return
		}

		Tampung_potensi_desa = append(Tampung_potensi_desa, ambil)

	}

	fmt.Println(Tampung_potensi_desa)

	if len(Tampung_potensi_desa) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "potensi_desa tersedia",
			"data":    Tampung_potensi_desa,
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return

	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "potensi_desa tidak ada",
			"data":    []Data_potensi_desa{},
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}

}
