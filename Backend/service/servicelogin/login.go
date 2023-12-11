package servicelogin

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Login(c *gin.Context) {
	type Pengguna struct {
		ID           int    `json:"id_pengguna"`
		RolePengguna string `json:"role_pengguna"`
		RoleID       int    `json:"role_id"`
		DesaID       int    `json:"desa_id"`
		Token        string `json:"token"`
	}

	type Request struct {
		Username string `json:"username"`
		Password string `json:"password"`
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

	if input.Username == "" {
		c.JSON(400, gin.H{
			"status":  false,
			"message": "Username Kosong",
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}

	if input.Password == "" {
		c.JSON(400, gin.H{
			"status":  false,
			"message": "Password Kosong",
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}

	var validasi_pass string

	cek_password := `
		select password from dev.pengguna where UPPER(username) = UPPER($1)
	`
	err = tx.QueryRow(ctx, cek_password, input.Username).Scan(&validasi_pass)
	if err != nil {
		log.Fatal(err)
	}

	hasher := md5.New()
	hasher.Write([]byte(input.Password))
	hashedInputPassword := hex.EncodeToString(hasher.Sum(nil))

	if hashedInputPassword != validasi_pass {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Password Salah, Login Gagal",
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}

	query_pengguna := `
		select 
			id_pengguna,
			role_pengguna,
			role_id,
			desa_id,
			COALESCE(token, '') AS token
		from dev.pengguna
		where UPPER(username) = UPPER($1)
	`
	row1, err := tx.Query(ctx, query_pengguna, input.Username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		err = tx.Commit(ctx)
		if err != nil {
			panic(err.Error())
		}
		return
	}

	defer row1.Close()

	var Tampung_pengguna []Pengguna

	for row1.Next() {
		var ambil Pengguna
		err := row1.Scan(
			&ambil.ID,
			&ambil.RolePengguna,
			&ambil.RoleID,
			&ambil.DesaID,
			&ambil.Token,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			err = tx.Commit(ctx)
			if err != nil {
				panic(err.Error())
			}
			return
		}

		Tampung_pengguna = append(Tampung_pengguna, ambil)
	}

	if len(Tampung_pengguna) > 0 {
		Tampung_pengguna[0].Token, err = generateToken(Tampung_pengguna[0].ID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			err = tx.Commit(ctx)
			if err != nil {
				panic(err.Error())
			}
			return
		}

		err = updateTokenInDatabase(Tampung_pengguna[0].ID, Tampung_pengguna[0].Token)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			err = tx.Commit(ctx)
			if err != nil {
				panic(err.Error())
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Berhasil Login!",
			"data":    Tampung_pengguna,
		})
		err = tx.Commit(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return

	}

}
