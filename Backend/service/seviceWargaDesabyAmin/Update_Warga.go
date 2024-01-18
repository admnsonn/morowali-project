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
		IDPengguna          string `json:"id_pengguna"`
		NIK                 string `json:"nik"`
		NoTelp              string `json:"no_telp"`
		TanggalLahir        string `json:"tanggal_lahir"`
		TempatLahir         string `json:"tempat_lahir"`
		JenisKelamin        string `json:"jenis_kelamin"`
		StatusPenikahan     string `json:"status_perkawinan"`
		Profesi             string `json:"profesi"`
		KK                  string `json:"kk"`
		Umur                string `json:"umur"`
		Agama               string `json:"agama"`
		Desa                string `json:"desa"`
		DesaID              int    `json:"id_desa"`
		NamaLengkap         string `json:"nama_lengkap"`
		Alamat              string `json:"alamat_pengguna"`
		RT                  string `json:"rt"`
		RW                  string `json:"rw"`
		KodePos             string `json:"kode_pos"`
		FotoProfile         string `json:"foto_pengguna"`
		KategoriFinancialID string `json:"kategori_financial_id"`
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

	// Lakukan Update

	updateQuery := `
		UPDATE dev.pengguna
		SET
			nik = $2,
			no_telp = $3,
			tanggal_lahir = $4,
			tempat_lahir = $5,
			jenis_kelamin = $6,
			status_perkawinan= $7,
			profesi= $8,
			kk= $9,
			umur= $10,
			agama = $11,
			desa = $12,
			id_desa = $13,
			nama_lengkap = $14,
			alamat_pengguna = $15,
			rt = $16,
			rw = $17,
			kode_pos = $18,
			foto_pengguna = $19,
			kategori_financial_id = $20
		WHERE
			id_pengguna = $1
	`

	// Execute update query with input parameters
	_, err = tx.Exec(ctx, updateQuery)
	if err != nil {
		fmt.Println("Gagal melakukan UPDATE:", err)
		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil Update",
	})

}
