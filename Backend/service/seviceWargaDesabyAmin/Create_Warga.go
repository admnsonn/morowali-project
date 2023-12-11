package sevicewargadesabyamin

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/rwcarlsen/goexif/exif"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/image/draw"
)

func Create_warga(c *gin.Context) {
	type Request struct {
		Username     string `json:"username"`
		Password     string `json:"password"`
		NIK          string `json:"nik"`
		NoTelp       string `json:"no_telp"`
		FotoProfile  string `json:"foto_pengguna"`
		TanggalLahir string `json:"tanggal_lahir"`
		TempatLahir  string `json:"tempat_lahir"`
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

	var hitung int

	if input.Username != "" {
		cek_usename := `
			SELECT COUNT(*)
			FROM dev.pengguna
			WHERE UPPER(username) = UPPER($1);
		`
		err = tx.QueryRow(ctx, cek_usename, input.Username).Scan(&hitung)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			err = tx.Commit(ctx)
			if err != nil {
				panic(err.Error())
			}
			return
		}

		if hitung != 0 {
			c.JSON(http.StatusOK, gin.H{
				"status":  true,
				"message": "Username Sudah Terdaftar !",
			})
			err = tx.Commit(ctx)
			if err != nil {
				fmt.Println(err.Error())
			}
			return
		}
	}

	if input.NIK != "" {
		cek_usename := `
		SELECT COUNT(*)
		FROM dev.pengguna
		WHERE nik = $1;
		`
		err = tx.QueryRow(ctx, cek_usename, input.NIK).Scan(&hitung)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			err = tx.Commit(ctx)
			if err != nil {
				panic(err.Error())
			}
			return
		}

		if hitung != 0 {
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
	}

	costFactor := 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), costFactor)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if input.FotoProfile != "" {

		percent := 0.5

		match := strings.HasPrefix(input.FotoProfile, "data:image/")

		if match {
			dataParts := strings.Split(input.FotoProfile, ",")
			if len(dataParts) != 2 {
				fmt.Println("Invalid base64 image format.")
				return
			}

			data := dataParts[1]
			dataBytes, err := base64.StdEncoding.DecodeString(data)
			if err != nil {
				fmt.Println("Error decoding base64 data:", err)
				return
			}

			imgReader := bytes.NewReader(dataBytes)
			img, _, err := image.Decode(imgReader)
			if err != nil {
				fmt.Println("Error decoding image:", err)
				return
			}

			fmt.Println("Image successfully decoded.")
			exifOrientation, err := getExifOrientation(dataBytes)
			if err != nil {
				fmt.Println("Error getting EXIF orisentation:", err)
				return
			}
			var im2 image.Image

			if exifOrientation != 0 {
				switch exifOrientation {
				case 8:
					im2 = rotateImage(img, -90)
				case 3:
					im2 = rotateImage(img, 180)
				case 6:
					im2 = rotateImage(img, 90)
				default:
					im2 = img
				}
			}

			// Calculate new width and height
			newWidth := int(float64(im2.Bounds().Dx()) * percent)
			newHeight := int(float64(im2.Bounds().Dy()) * percent)
			newImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
			draw.ApproxBiLinear.Scale(newImg, newImg.Bounds(), im2, im2.Bounds(), draw.Over, nil)

			// Convert newImg to base64
			var buf bytes.Buffer
			err = jpeg.Encode(&buf, newImg, nil)
			if err != nil {
				log.Println("Error encoding image to JPEG:", err)
				return
			}

			newImgBase64 := base64.StdEncoding.EncodeToString(buf.Bytes())

			input.FotoProfile = newImgBase64

			log.Println("Processed image in base64:", newImgBase64)
		} else {
			fmt.Println("Image format is not valid.")
		}
	} else {
		fmt.Println("FotoProfile is empty.")
	}

	query := `
	INSERT INTO dev.pengguna (
		username, 1
		password, 2
		role_pengguna, 3
		role_id, 4
		created_at, 5
		created_by, 6
		update_at, 7
		updated_by, 8
		tanggal_lahir, 9
		tempat_lahir, 10
		jk_id, 11
		jenis_kelamin, 12
		status_perkawinan_id, 13
		status_perkawinan, 14 
		profesi, 15 
		nik, 16
		kk, 17
		umur, 18
		kategori_financial_id, 19
		agama, 20
		agama_id, 21
		desa, 22
		desa_id, 23
		nama_lengkap, 24
		foto_pengguna, 25
		alamat_id, 26
		alamat_pengguna, 27
		rt, 28
		rw, 29
		kode_pos,  30
		no_telp 31
	) VALUES (
		$1, $2, $3, $4, $5, 
		$6, $7, $8, $9, $10,
		$11, $12, $13, $14, $15, 
		$16, $17, $18, $19, $20,
		$21, $22, $23, $24, $25, 
		$26, $27, $28, $29, $30,
		$31
	)
	`

	_, err = tx.Exec(ctx, query,
		input.Username, hashedPassword, "Warga", 2, time.Now(),
		"Admin", time.Now(), "Admin", input.NIK, input.FotoProfile,
	)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		fmt.Println(err.Error())
		fmt.Println("ERROR DI QUERY = Pada kondisi foto kosong")
		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil Input",
	})

}

func rotateImage(src image.Image, angle float64) image.Image {
	dst := image.NewRGBA(image.Rect(0, 0, src.Bounds().Dy(), src.Bounds().Dx()))

	switch angle {
	case -90:
		fmt.Println("MASUK -90")

		draw.Copy(dst, image.Pt(0, 0), src, src.Bounds(), draw.Over, nil)
		for x := 0; x < dst.Bounds().Dx(); x++ {
			for y := 0; y < dst.Bounds().Dy(); y++ {
				dst.Set(x, y, src.At(src.Bounds().Dx()-y-1, x))
			}
		}
	case 90:
		fmt.Println("MASUK 90")

		draw.Copy(dst, image.Pt(0, 0), src, src.Bounds(), draw.Over, nil)
		for x := 0; x < dst.Bounds().Dx(); x++ {
			for y := 0; y < dst.Bounds().Dy(); y++ {
				dst.Set(x, y, src.At(y, src.Bounds().Dy()-x-1))
			}
		}
	case 180:
		fmt.Println("MASUK 180")

		draw.Copy(dst, image.Pt(0, 0), src, src.Bounds(), draw.Over, nil)
		for x := 0; x < dst.Bounds().Dx(); x++ {
			for y := 0; y < dst.Bounds().Dy(); y++ {
				dst.Set(x, y, src.At(src.Bounds().Dx()-x-1, src.Bounds().Dy()-y-1))
			}
		}
	default:

		fmt.Println("MASUK DEFAULT")
		draw.Copy(dst, image.Pt(0, 0), src, src.Bounds(), draw.Over, nil)
	}

	return dst
}

func getExifOrientation(imageData []byte) (int, error) {
	exifData, err := exif.Decode(bytes.NewReader(imageData))
	if err != nil {
		return 0, err
	}

	orientationTag, err := exifData.Get(exif.Orientation)
	if err != nil {
		return 0, err
	}

	orientationValue, err := orientationTag.Int(0)
	if err != nil {
		return 0, err
	}

	return orientationValue, nil
}
