package serviceumkm

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

func Tambah_UMKM(c *gin.Context) {
	type UMKM struct {
		Nama         string `json:"nama_umkm"`
		Konten       string `json:"konten_umkm"`
		KategoriID   int    `json:"kategori_umkm_id"`
		Foto         string `json:"foto_umkm"`
		DesaID       int    `json:"desa_id"`
		NomorTelepon string `json:"no_telp_umkm"`
		Alamat       string `json:"alamat_umkm"`
	}

	var input UMKM

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

	if input.Foto != "" {

		percent := 0.5

		match := strings.HasPrefix(input.Foto, "data:image/")

		if match {

			dataParts := strings.Split(input.Foto, ",")
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

			newWidth := int(float64(im2.Bounds().Dx()) * percent)
			newHeight := int(float64(im2.Bounds().Dy()) * percent)
			newImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
			draw.ApproxBiLinear.Scale(newImg, newImg.Bounds(), im2, im2.Bounds(), draw.Over, nil)

			var buf bytes.Buffer
			err = jpeg.Encode(&buf, newImg, nil)
			if err != nil {
				log.Println("Error encoding image to JPEG:", err)
				return
			}

			newImgBase64 := base64.StdEncoding.EncodeToString(buf.Bytes())

			input.Foto = newImgBase64

			log.Println("Processed image in base64:", newImgBase64)
		} else {
			fmt.Println("Image format is not valid.")
		}
	} else {
		fmt.Println("Foto is empty.")
	}

	var id_next, id_now int

	cek_berita := `
	SELECT MAX(id_umkm) AS last_id
	FROM dev.umkm;
	`
	err = tx.QueryRow(ctx, cek_berita).Scan(&id_now)

	if err != nil {
		id_next = id_now + 1
	} else {
		id_next = id_now + 1
	}

	query := `
		INSERT INTO dev.umkm (
			id_umkm ,nama_umkm , konten_umkm , kategori_umkm_id , foto_umkm , desa_id , no_telp_umkm , alamat,
			created_at,
			created_by,
			update_at,
			updated_by
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`

	_, err = tx.Exec(ctx, query, id_next, input.Nama, input.Konten, input.KategoriID, input.Foto, input.DesaID, input.NomorTelepon, input.Alamat,
		time.Now(), "admin", time.Now(), "admin")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Error committing transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Berita berhasil ditambahkan"})
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
