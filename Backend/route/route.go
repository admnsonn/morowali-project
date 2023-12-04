package route

import (
	serviceberita "backendpgx7071/service/serviceBerita"
	servicekreativitas "backendpgx7071/service/serviceKreativitas"
	servicemisidesa "backendpgx7071/service/serviceMisiDesa"
	serviceperaturandesa "backendpgx7071/service/servicePeraturanDesa"
	servicepotensidesa "backendpgx7071/service/servicePotensiDesa"
	serviceumkm "backendpgx7071/service/serviceUMKM"
	servicevisi "backendpgx7071/service/serviceVisi"
	servicewilayahdesa "backendpgx7071/service/serviceWilayahDesa"
	servicewisata "backendpgx7071/service/serviceWisata"

	"backendpgx7071/service/servicelogin"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	router.POST("/login", servicelogin.Login)
	router.POST("/validasi", servicelogin.ValidasiToken)

	beritaRoutes := router.Group("/berita")
	{
		beritaRoutes.GET("/", serviceberita.SemuaBerita)
		beritaRoutes.GET("/:id", serviceberita.DetailBerita)
	}

	umkmRoutes := router.Group("/umkm")
	{
		umkmRoutes.GET("/", serviceumkm.Semuaumkm)
		umkmRoutes.GET("/:id", serviceumkm.Detailumkm)
	}

	wisataRoutes := router.Group("/wisata")
	{
		wisataRoutes.GET("/", servicewisata.Semuawisata)
		wisataRoutes.GET("/:id", servicewisata.Detailwisata)
	}

	potensiDesa := router.Group("/potensi_desa")
	{
		potensiDesa.GET("/", servicepotensidesa.Semuapotensi_desa)
		potensiDesa.GET("/:id", servicepotensidesa.Detailpotensi_desa)
	}

	Kreatifitas := router.Group("/kreatifitas")
	{
		Kreatifitas.GET("/", servicekreativitas.Kreatifitas_desa)
	}

	Visi := router.Group("/visi")
	{
		Visi.GET("/", servicevisi.Visi_desa)
	}

	Misi := router.Group("/misi")
	{
		Misi.GET("/", servicemisidesa.Misi_desa)
	}

	Peraturan := router.Group("/peraturan")
	{
		Peraturan.GET("/", serviceperaturandesa.Peraturan_desa)
	}

	Wilayah_desa := router.Group("/Wilayah_desa")
	{
		Wilayah_desa.GET("/perkebunan", servicewilayahdesa.Wilayah_Perkebunan)
		Wilayah_desa.GET("/warga", servicewilayahdesa.Total_warga)
		Wilayah_desa.GET("/desa", servicewilayahdesa.Luas_desa)
	}

}
