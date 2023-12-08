package serviceumum

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Cek_URL(c *gin.Context) {
	type Umum struct {
		IDdesa int `json:"id_desa"`
	}

	type Request struct {
		URL string `json:"cari_link_url"`
	}

	var input Request
	var respon Umum

	if c.GetHeader("content-type") == "application/x-www-form-urlencoded" || c.GetHeader("content-type") == "application/x-www-form-urlencoded; charset=utf-8" {

		if err := c.Bind(&input); err != nil {
			fmt.Println("MAsuk sini || Error masuk if ketika form")
			return
		}

	} else {

		if err := c.BindJSON(&input); err != nil {
			fmt.Println("MAsuk sini || Error masuk else ketika form")
			return
		}

	}

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(input.URL)

	query := `
	SELECT id_desa FROM dev.link_url WHERE url = $1;
	`

	ROW, err := tx.Query(ctx, query, input.URL)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		err = tx.Commit(ctx)
		if err != nil {
			panic(err.Error())
		}
		return
	}

	defer ROW.Close()

	var Tampung_umum []Umum

	for ROW.Next() {
		err := ROW.Scan(
			&respon.IDdesa,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			err = tx.Commit(ctx)
			if err != nil {
				panic(err.Error())
			}
			return
		}

		Tampung_umum = append(Tampung_umum, respon)
	}

	if len(Tampung_umum) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Link desa benar, dan ID DESA tersedia!",
			"id_desa": respon.IDdesa,
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Link desa tidak benar, atau ID DESA tidak tersedia!",
			"id_desa": []Umum{},
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}

}
