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

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"golang.org/x/image/draw"
)

func Update_data_warga(c *gin.Context) {
	type Request struct {
		IDPengguna      string `json:"id_pengguna"`
		Fullname        string `json:"nama_lengkap"`
		NIK             string `json:"nik"`
		KK              string `json:"kk"`
		TempatLahir     string `json:"tempat_lahir"`
		TanggalLahir    string `json:"tanggal_lahir"`
		JenisKelamin    string `json:"jenis_kelamin"`
		Alamat          string `json:"alamat_pengguna"`
		Agama           string `json:"agama"`
		StatusPenikahan string `json:"status_perkawinan"`
		Profesi         string `json:"profesi"`
		NoTelp          string `json:"no_telp"`
		PendidikanLast  string `json:"pendidikan_terakhir"`
		FotoProfile     string `json:"foto_profile"`
	}

	var input Request
	// var id_pendidikan int

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
	defer tx.Rollback(context.Background())

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

	query_update_pendidikan := `
		UPDATE dev.pendidikan
		SET
			tingkat_pendidikan = $1
		WHERE
			id_pendidikan = $2;
	`

	_, err = tx.Exec(ctx, query_update_pendidikan, input.PendidikanLast, input.IDPengguna)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		err = tx.Commit(ctx)
		if err != nil {
			panic(err.Error())
		}
		return
	}

	// query_update_pengguna := `
	// 	UPDATE dev.pengguna
	// 	SET
	// 		nik = $1,
	// 		kk = $2,
	// 		tempat_lahir = $3,
	// 		tanggal_lahir = $4,
	// 		jenis_kelamin = $5',
	// 		alamat_pengguna = $6,
	// 		agama = 'Islam',
	// 		status_perkawinan = $7,
	// 		profesi = $8,
	// 		no_telp = $9,
	// 		pendidikan_terakhir = $10,
	// 		foto_profile = $11
	// 	WHERE
	// 		id_pengguna = $12;

	// `

}
