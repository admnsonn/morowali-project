package route

import (
	serviceberita "main/service/serviceBerita"
	serviceidm "main/service/serviceIDM"
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
	router.POST("/pendidikan", serviceumum.Pendidikan)

	beritaRoutes := router.Group("/berita")
	{
		beritaRoutes.GET("/:id", serviceberita.DetailBerita)
		beritaRoutes.POST("/create", serviceberita.Tulis_Berita)
		beritaRoutes.DELETE("/delete/:id", serviceberita.DeleteBerita)
		beritaRoutes.GET("/categori", serviceberita.Kategori_berita)
		beritaRoutes.POST("/list", serviceberita.SemuaBerita)
		beritaRoutes.PUT("/update", serviceberita.UpdateBerita)
	}

	umkmRoutes := router.Group("/umkm")
	{
		umkmRoutes.POST("/list", serviceumkm.Semuaumkm)
		umkmRoutes.GET("/:id", serviceumkm.Detailumkm)
		umkmRoutes.GET("/umkm_kategori", serviceumkm.List_kat_umkm)
		umkmRoutes.POST("/tambah_umkm", serviceumkm.Tambah_UMKM)
		umkmRoutes.DELETE("/delete_umkm/:id", serviceumkm.Delete_umkm)
		umkmRoutes.PUT("/update_umkm", serviceumkm.Update_umkm)
	}

	wisataRoutes := router.Group("/wisata")
	{
		wisataRoutes.POST("/list", servicewisata.Semuawisata)
		wisataRoutes.GET("/:id", servicewisata.Detailwisata)
		wisataRoutes.POST("/tambah_data", servicewisata.Tambah_wisata)
		wisataRoutes.DELETE("/delete_wisata/:id", servicewisata.Delete_wisata)
		wisataRoutes.GET("/kategori_wisata", servicewisata.List_kat_wisata)
		wisataRoutes.PUT("/update_wisata", servicewisata.Update_Wisata)
	}

	potensiDesa := router.Group("/potensi_desa")
	{
		potensiDesa.GET("/list/:id", servicepotensidesa.Semuapotensi_desa)
		potensiDesa.GET("/:id", servicepotensidesa.Detailpotensi_desa)
		potensiDesa.POST("/tambah_potensi", servicepotensidesa.Tambah_potensi)
		potensiDesa.PUT("/update_potensi", servicepotensidesa.Update_Potensi)
		potensiDesa.DELETE("/delete_potensi/:id", servicepotensidesa.Delete_potensi)
	}

	Kreatifitas := router.Group("/kreatifitas")
	{
		Kreatifitas.GET("/list/:id", servicekreativitas.Kreatifitas_desa)
		Kreatifitas.POST("/tambah_kreatif", servicekreativitas.Tambah)
		Kreatifitas.PUT("/update", servicekreativitas.Update)
		Kreatifitas.GET("/detail/:id", servicekreativitas.Detail)
		Kreatifitas.DELETE("/delete/:id", servicekreativitas.Delete)
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

	Wilayah_desa := router.Group("/wilayah_desa")
	{
		Wilayah_desa.GET("/produksi", servicewilayahdesa.Wilayah_produksi)
		// Wilayah_desa.GET("/warga", servicewilayahdesa.Total_warga)
		Wilayah_desa.GET("/desa/:id", servicewilayahdesa.Wilayah_desa)
		Wilayah_desa.POST("/setting", servicewilayahdesa.Setting)

	}

	warga := router.Group("/warga")
	{
		// warga.GET("/perkebunan", sevicewargadesabyamin.Wilayah_Perkebunan)
		warga.POST("/list", sevicewargadesabyamin.Warga_desa_by_admin)
		warga.DELETE("/delete/:id", sevicewargadesabyamin.Delete_warga_desa)
		warga.GET("/detail/:id", sevicewargadesabyamin.Get_detail_warga_by_admin)
		warga.POST("/tambah", sevicewargadesabyamin.Create_warga)
		// warga.GET("/desa", sevicewargadesabyamin.Luas_desa)
		warga.PUT("/update", sevicewargadesabyamin.Update_data_warga)
		warga.GET("/financial", sevicewargadesabyamin.Kategori_financial)
		warga.GET("/pendidikan", sevicewargadesabyamin.Kategori_pendidikan)
		warga.GET("/agama", sevicewargadesabyamin.Kategori_agama)
		warga.GET("/jk", sevicewargadesabyamin.JK_pengguna)
		warga.GET("/pernikahan", sevicewargadesabyamin.Pernikahan)

	}

	//sejarahdesa/kepaladesa

	sejarahdesa := router.Group("/sejarahdesa")
	{
		// warga.GET("/perkebunan", sevicewargadesabyamin.Wilayah_Perkebunan)
		sejarahdesa.GET("/kepaladesa/:id", servicesejarahdesa.Kepaladesa)
		sejarahdesa.GET("/sejarahdesa/:id", servicesejarahdesa.Sejarahdesa)
		// warga.GET("/desa", sevicewargadesabyamin.Luas_desa)
	}

	idmiksike := router.Group("/idmiksike")
	{
		idmiksike.POST("/input", serviceidm.Input)
		idmiksike.GET("/sejarahdesa/:id", servicesejarahdesa.Sejarahdesa)
	}

}

// ROUTE GOLANG
