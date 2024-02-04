package servicewilayahdesa

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Setting(c *gin.Context) {
	type Wilayah_update struct {
		Longitude      string `json:"longitude"`
		Latitude       string `json:"latitude"`
		Status_wilayah string `json:"status_wilayah"`
		LuasWilayah    string `json:"luas_wilayah"`
		ID             string `json:"id_desa"`
	}

	var input Wilayah_update

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
	defer tx.Rollback(ctx)

	if input.Status_wilayah == "update" {
		query := `
			UPDATE dev.desa
			SET luas_wilayah = $1, longitude  = $2, latitude = $3 
			WHERE id_desa = $4
		`

		_, err = tx.Exec(ctx, query, input.LuasWilayah, input.Longitude, input.Latitude, input.ID)
		if err != nil {
			fmt.Println("MASUK SINI")
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			return
		}

		err = tx.Commit(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"status": true, "message": "WILAYAH DESA updated successfully"})
	} else {
		query := `
			UPDATE dev.desa
			SET luas_wilayah = $1, longitude  = $2, latitude = $3 , status_wilayah = $4
			WHERE id_desa = $5
		`

		_, err = tx.Exec(ctx, query, input.LuasWilayah, input.Longitude, input.Latitude, "update", input.ID)
		if err != nil {
			fmt.Println("MASUK SINI")
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			return
		}

		err = tx.Commit(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"status": true, "message": "WILAYAH DESA updated successfully"})
	}

}
