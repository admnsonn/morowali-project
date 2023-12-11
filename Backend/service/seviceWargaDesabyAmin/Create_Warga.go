package sevicewargadesabyamin

import (
	servicecheck "backendpgx7071/service/serviceCheck"
	"bytes"
	"context"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
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
		Username        string `json:"username"`
		Password        string `json:"password"`
		NIK             string `json:"nik"`
		NoTelp          string `json:"no_telp"`
		FotoProfile     string `json:"foto_pengguna"`
		TanggalLahir    string `json:"tanggal_lahir"`
		TempatLahir     string `json:"tempat_lahir"`
		JenisKelamin    string `json:"jenis_kelamin"`
		StatusPenikahan string `json:"status_perkawinan"`
		Profesi         string `json:"profesi"`
		KK              string `json:"kk"`
		Umur            string `json:"umur"`
		FinancialID     string `json:"kategori_financial_id"`
		Agama           string `json:"agama"`
		Desa            string `json:"desa"`
		DesaID          int    `json:"id_desa"`
		NamaLengkap     string `json:"nama_lengkap"`
		Alamat          string `json:"alamat_pengguna"`
		RT              string `json:"rt"`
		RW              string `json:"rw"`
		KodePos         string `json:"kode_pos"`
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

	var hitung1, hitung2 int

	cek_usename := `
			SELECT COUNT(*)
			FROM dev.pengguna
			WHERE UPPER(username) = UPPER($1);
		`
	err = tx.QueryRow(ctx, cek_usename, input.Username).Scan(&hitung1)
	if err != nil {
		log.Fatal(err)
	}

	if hitung1 != 0 {
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

	hasher := md5.New()
	hasher.Write([]byte(input.Password))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))

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

	var jk_id, nikah_id, agama_id int

	jk_id = servicecheck.CheckGender(input.JenisKelamin)

	nikah_id = servicecheck.CheckGender(input.StatusPenikahan)

	agama_id = servicecheck.CheckGender(input.Agama)

	// Query INSERT
	query := `
		INSERT INTO dev.pengguna (
			username, 
			password,
			role_pengguna, 
			role_id,
			created_at, 
			created_by, 
			update_at, 
			updated_by,
			tanggal_lahir, 
			tempat_lahir, 
			jk_id, 
			jenis_kelamin,
			status_perkawinan_id, 
			status_perkawinan, 
			profesi, 
			nik, 
			kk,
			umur, 
			kategori_financial_id, 
			agama,
			agama_id, 
			desa, 
			desa_id,
			nama_lengkap, 
			foto_pengguna, 
			alamat_pengguna, 
			rt, 
			rw,
			kode_pos,  
			no_telp
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30
		)
	`

	// Execute query dengan parameter input
	_, err = tx.Exec(ctx, query,
		input.Username, hashedPassword, "Warga", 2, time.Now(), "Admin Sistem", time.Now(), "Admin Sistem",
		input.TanggalLahir, input.TempatLahir, jk_id, input.JenisKelamin, nikah_id, input.StatusPenikahan,
		input.Profesi, input.NIK, input.KK, input.Umur, input.FinancialID, input.Agama, agama_id, input.Desa, input.DesaID,
		input.NamaLengkap, input.FotoProfile, input.Alamat, input.RT, input.RW, input.KodePos, input.NoTelp,
	)
	if err != nil {
		fmt.Println("Gagal melakukan INSERT:", err)
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
