package route

import (
	serviceberita "backendpgx7071/service/serviceBerita"
	servicekreativitas "backendpgx7071/service/serviceKreativitas"
	servicemisidesa "backendpgx7071/service/serviceMisiDesa"
	serviceperaturandesa "backendpgx7071/service/servicePeraturanDesa"
	servicepotensidesa "backendpgx7071/service/servicePotensiDesa"
	servicesambutandesa "backendpgx7071/service/serviceSambutanDesa"
	servicesejarahdesa "backendpgx7071/service/serviceSejarahDesa"
	serviceumkm "backendpgx7071/service/serviceUMKM"
	serviceumum "backendpgx7071/service/serviceUmum"
	servicevisi "backendpgx7071/service/serviceVisi"
	servicewilayahdesa "backendpgx7071/service/serviceWilayahDesa"
	servicewisata "backendpgx7071/service/serviceWisata"
	"backendpgx7071/service/servicelogin"
	sevicewargadesabyamin "backendpgx7071/service/seviceWargaDesabyAmin"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	router.POST("/login", servicelogin.Login)
	router.POST("/validasi", servicelogin.ValidasiToken)
	router.POST("/link_url", serviceumum.Cek_URL)

	beritaRoutes := router.Group("/berita")
	{
		beritaRoutes.GET("/list/:id", serviceberita.SemuaBerita)
		beritaRoutes.GET("/:id", serviceberita.DetailBerita)
	}

	umkmRoutes := router.Group("/umkm")
	{
		umkmRoutes.GET("/list/:id", serviceumkm.Semuaumkm)
		umkmRoutes.GET("/:id", serviceumkm.Detailumkm)
	}

	wisataRoutes := router.Group("/wisata")
	{
		wisataRoutes.GET("/list/:id", servicewisata.Semuawisata)
		wisataRoutes.GET("/:id", servicewisata.Detailwisata)
	}

	potensiDesa := router.Group("/potensi_desa")
	{
		potensiDesa.GET("/list/:id", servicepotensidesa.Semuapotensi_desa)
		potensiDesa.GET("/:id", servicepotensidesa.Detailpotensi_desa)
	}

	Kreatifitas := router.Group("/kreatifitas")
	{
		Kreatifitas.GET("/list/:id", servicekreativitas.Kreatifitas_desa)
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

	Sambutan := router.Group("/sambutan")
	{
		Sambutan.GET("/", servicesambutandesa.Sambutan_desa)
	}

	Wilayah_desa := router.Group("/Wilayah_desa")
	{
		Wilayah_desa.GET("/perkebunan", servicewilayahdesa.Wilayah_Perkebunan)
		Wilayah_desa.GET("/warga", servicewilayahdesa.Total_warga)
		Wilayah_desa.GET("/desa", servicewilayahdesa.Luas_desa)
	}

	warga := router.Group("/warga") // Menambahkan grup "/warga" di dalam grup "/admin"
	{
		// warga.GET("/perkebunan", sevicewargadesabyamin.Wilayah_Perkebunan)
		warga.POST("/list", sevicewargadesabyamin.Warga_desa_by_admin)
		warga.DELETE("/delete/:id", sevicewargadesabyamin.Delete_warga_desa)
		warga.GET("/detail/:id", sevicewargadesabyamin.Get_detail_warga_by_admin)
		warga.POST("/tambah", sevicewargadesabyamin.Create_warga)
		// warga.GET("/desa", sevicewargadesabyamin.Luas_desa)
	}

	//sejarahdesa/kepaladesa

	sejarahdesa := router.Group("/sejarahdesa") // Menambahkan grup "/warga" di dalam grup "/admin"
	{
		// warga.GET("/perkebunan", sevicewargadesabyamin.Wilayah_Perkebunan)
		sejarahdesa.GET("/kepaladesa/:id", servicesejarahdesa.Kepaladesa)
		sejarahdesa.GET("/sejarahdesa/:id", servicesejarahdesa.Sejarahdesa)
		// warga.GET("/desa", sevicewargadesabyamin.Luas_desa)
	}

}
