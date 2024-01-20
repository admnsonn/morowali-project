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
	"golang.org/x/image/draw"
)

func Create_warga(c *gin.Context) {
	type Request struct {
		Fullname        string `json:"nama_lengkap"`
		NIK             string `json:"nik"`
		KK              string `json:"kk"`
		TempatLahir     string `json:"tempat_lahir"`
		TanggalLahir    string `json:"tanggal_lahir"`
		JenisKelamin    string `json:"jenis_kelamin_id"`
		Alamat          string `json:"alamat_pengguna"`
		Agama           string `json:"agama_id"`
		StatusPenikahan string `json:"status_perkawinan_id"`
		Profesi         string `json:"profesi"`
		NoTelp          string `json:"no_telp"`
		PendidikanLast  string `json:"pendidikan_terakhir_id"`
		FotoProfile     string `json:"foto_profile"`
		RW              string `json:"rw"`
		RT              string `json:"rt"`
		DesaID          string `json:"id_desa"`
		Financial       string `json:"kategori_financial_id"`
		KodePos         string `json:"kode_pos"`
		Kewarganegaraan string `json:"kewarganegaraan"`
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

	var hitung2 int

	cek_nik := `
		SELECT COUNT(*)
		FROM dev.pengguna
		WHERE nik = $1;
		`
	err = tx.QueryRow(ctx, cek_nik, input.NIK).Scan(&hitung2)
	if err != nil {
		log.Fatal(err)
	}

	if hitung2 != 0 {
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

	// Query INSERT

	query := `
		INSERT INTO dev.pengguna (
			nama_lengkap,
			nik,
			kk,
			tempat_lahir,
			tanggal_lahir,
			jk_id ,
			alamat_pengguna,
			agama_id ,
			status_perkawinan_id ,
			profesi,
			no_telp,
			pendidikan_id ,
			foto_pengguna ,
			rw,
			rt,
			desa_id,
			kategori_financial_id ,
			role_id,
			kode_pos ,
			kewarganegaraan,
			created_at ,
			created_by ,
			update_at ,
			updated_by 
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24
		)
		RETURNING id_pengguna;
		`

	var newID int

	err = tx.QueryRow(ctx, query,
		input.Fullname,
		input.NIK,
		input.KK,
		input.TempatLahir,
		input.TanggalLahir,
		input.JenisKelamin,
		input.Alamat,
		input.Agama,
		input.StatusPenikahan,
		input.Profesi,
		input.NoTelp,
		input.PendidikanLast,
		input.FotoProfile,
		input.RW,
		input.RT,
		input.DesaID,
		input.Financial,
		"2",
		input.KodePos,
		input.Kewarganegaraan,
		time.Now(),
		"Admin Sistem",
		time.Now(),
		"Admin Sistem",
	).Scan(&newID)

	if err != nil {
		// log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Query Insert Bermasalah",
		})
	}

	// Commit the transaction
	err = tx.Commit(ctx)
	if err != nil {
		// log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Data Gagal Ditambahkan",
		})
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Data Berhasil Di Tambahkan",
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
