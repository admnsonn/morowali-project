package servicekreativitas

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Update(c *gin.Context) {
	type Creativity struct {
		Title       string `json:"judul_kreatifitas"`
		Description string `json:"deskripsi"`
		Photo       string `json:"foto_kreatifitas"`
		ID          int    `json:"id_kreatifitas"`
		Pengunjung  string `json:"pengunjung"`
	}

	var input Creativity

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
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}
	defer tx.Rollback(ctx)

	query := `
		UPDATE dev.kreatifitas
		SET judul_kreatifitas = $1, deskripsi  = $2, foto_kreatifitas = $3, pengunjung = $4
		WHERE id_kreatifitas = $5
	`

	_, err = tx.Exec(ctx, query, input.Title, input.Description, input.Photo, input.Pengunjung, input.ID)
	if err != nil {
		fmt.Println("MASUK SINI")
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "KREATIFITAS updated successfully"})
}
