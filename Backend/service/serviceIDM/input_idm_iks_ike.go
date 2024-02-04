package serviceidm

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Input(c *gin.Context) {
	type IDM struct {
		Tahun      string `json:"tahun"`
		Indikator  string `json:"indikator"`
		Skor       string `json:"skor"`
		Keterangan string `json:"keterangan"`
		Kegiatan   string `json:"kegiatan"`
		Nilai      string `json:"nilai"`
		Pusat      string `json:"pusat"`
		Provinsi   string `json:"provinsi"`
		Kabupaten  string `json:"kabupaten"`
		Desa       string `json:"desa"`
		CSR        string `json:"csr"`
		Lainya     string `json:"lainya"`
	}

	type IKS struct {
		Tahun      string `json:"tahun"`
		Indikator  string `json:"indikator"`
		Skor       string `json:"skor"`
		Keterangan string `json:"keterangan"`
		Kegiatan   string `json:"kegiatan"`
		Nilai      string `json:"nilai"`
		Pusat      string `json:"pusat"`
		Provinsi   string `json:"provinsi"`
		Kabupaten  string `json:"kabupaten"`
		Desa       string `json:"desa"`
		CSR        string `json:"csr"`
		Lainya     string `json:"lainya"`
	}

	type IKE struct {
		Tahun      string `json:"tahun"`
		Indikator  string `json:"indikator"`
		Skor       string `json:"skor"`
		Keterangan string `json:"keterangan"`
		Kegiatan   string `json:"kegiatan"`
		Nilai      string `json:"nilai"`
		Pusat      string `json:"pusat"`
		Provinsi   string `json:"provinsi"`
		Kabupaten  string `json:"kabupaten"`
		Desa       string `json:"desa"`
		CSR        string `json:"csr"`
		Lainya     string `json:"lainya"`
	}

	type Data struct {
		IDMinput []IDM `json:"data_idm"`
		IKSinput []IKS `json:"data_iks"`
		IKEinput []IKE `json:"data_ike"`
	}

	type Input_data struct {
		DesaID string `json:"desa_id"`
		Data   []Data `json:"data"`
	}

	var input Input_data

	if c.GetHeader("content-type") == "application/x-www-form-urlencoded" || c.GetHeader("content-type") == "application/x-www-form-urlencoded; charset=utf-8" {

		if err := c.Bind(&input); err != nil {
			return
		}

	} else {

		if err := c.BindJSON(&input); err != nil {
			return
		}

	}

	if len(input.Data) != 0 {
		ch := make(chan struct{})

		go func() {
			for _, v := range input.Data[0].IDMinput {
				fmt.Println("INI DATA IDM INPUT", v)
			}

			ch <- struct{}{}
		}()

		go func() {
			for _, v := range input.Data[0].IKSinput {
				fmt.Println("INI DATA IKS INPUT", v)
			}

			ch <- struct{}{}
		}()

		go func() {
			for _, v := range input.Data[0].IKEinput {
				fmt.Println("INI DATA IKE INPUT", v)
			}

			ch <- struct{}{}
		}()

		for i := 0; i < 3; i++ {
			<-ch
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": input})
}
