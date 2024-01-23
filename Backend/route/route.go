package route

import (
	serviceberita "main/service/serviceBerita"
	servicekreativitas "main/service/serviceKreativitas"
	servicemisidesa "main/service/serviceMisiDesa"
	serviceperaturandesa "main/service/servicePeraturanDesa"
	servicepotensidesa "main/service/servicePotensiDesa"
	servicesambutandesa "main/service/serviceSambutanDesa"
	servicesejarahdesa "main/service/serviceSejarahDesa"
	serviceumkm "main/service/serviceUMKM"
	serviceumum "main/service/serviceUmum"
	servicevisi "main/service/serviceVisi"
	servicewilayahdesa "main/service/serviceWilayahDesa"
	servicewisata "main/service/serviceWisata"
	"main/service/servicelogin"
	sevicewargadesabyamin "main/service/seviceWargaDesabyAmin"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	router.POST("/login", servicelogin.Login)
	router.POST("/validasi", servicelogin.ValidasiToken)
	router.POST("/link_url", serviceumum.Cek_URL)

	beritaRoutes := router.Group("/berita")
	{
		beritaRoutes.GET("/:id", serviceberita.DetailBerita)
		beritaRoutes.POST("/create", serviceberita.Tulis_Berita)
		beritaRoutes.DELETE("/delete/:id", serviceberita.DeleteBerita)
		beritaRoutes.GET("/categori", serviceberita.Kategori_berita)
		beritaRoutes.POST("/list", serviceberita.SemuaBerita)
	}

	umkmRoutes := router.Group("/umkm")
	{
		umkmRoutes.POST("/list", serviceumkm.Semuaumkm)
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

	warga := router.Group("/warga")
	{
		// warga.GET("/perkebunan", sevicewargadesabyamin.Wilayah_Perkebunan)
		warga.POST("/list", sevicewargadesabyamin.Warga_desa_by_admin)
		warga.DELETE("/delete/:id", sevicewargadesabyamin.Delete_warga_desa)
		warga.GET("/detail/:id", sevicewargadesabyamin.Get_detail_warga_by_admin)
		warga.POST("/tambah", sevicewargadesabyamin.Create_warga)
		// warga.GET("/desa", sevicewargadesabyamin.Luas_desa)
		warga.PUT("/update", sevicewargadesabyamin.Create_warga)
		warga.GET("/financial", sevicewargadesabyamin.Kategori_financial)
		warga.GET("/pendidikan", sevicewargadesabyamin.Kategori_pendidikan)
		warga.GET("/agama", sevicewargadesabyamin.Kategori_agama)
		warga.GET("/jk", sevicewargadesabyamin.JK_pengguna)

	}

	//sejarahdesa/kepaladesa

	sejarahdesa := router.Group("/sejarahdesa")
	{
		// warga.GET("/perkebunan", sevicewargadesabyamin.Wilayah_Perkebunan)
		sejarahdesa.GET("/kepaladesa/:id", servicesejarahdesa.Kepaladesa)
		sejarahdesa.GET("/sejarahdesa/:id", servicesejarahdesa.Sejarahdesa)
		// warga.GET("/desa", sevicewargadesabyamin.Luas_desa)
	}

}

// ROUTE GOLANG
