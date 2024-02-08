package route

import (
	servicepemerintahgo "main/service/ServicePemerintah.go"
	serviceproduksi "main/service/ServiceProduksi"
	serviceberita "main/service/serviceBerita"
	serviceidm "main/service/serviceIDM"
	serviceinformasipembangunango "main/service/serviceInformasiPembangunan.go"
	servicekreativitas "main/service/serviceKreativitas"
	servicepegawai "main/service/servicePegawai"
	serviceperaturandesa "main/service/servicePeraturanDesa"
	servicepotensidesa "main/service/servicePotensiDesa"
	serviceprodukhukum "main/service/serviceProdukHukum"
	serviceumkm "main/service/serviceUMKM"
	serviceumum "main/service/serviceUmum"
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
	router.GET("/galeri/:id", serviceumum.GaleriFoto)
	router.GET("/agama/:id", serviceumum.Agama)

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

	Pemerintah := router.Group("/pemerintah")
	{
		Pemerintah.GET("/visi-misi/:id", servicepemerintahgo.Visi_Misi)
		Pemerintah.PUT("/visi-misi/setting", servicepemerintahgo.Setting_visimisi)
		Pemerintah.GET("/sambutan/:id", servicepemerintahgo.Sambutan)
		Pemerintah.PUT("/sambutan/setting", servicepemerintahgo.Setting_sambutan)
		Pemerintah.GET("/sejarah/:id", servicepemerintahgo.Sejarah)
		Pemerintah.PUT("/sejarah/setting", servicepemerintahgo.Setting_sejarah)
		Pemerintah.GET("/proker/:id", servicepemerintahgo.Proker)
		Pemerintah.PUT("/proker/setting", servicepemerintahgo.Setting_Proker)
		Pemerintah.GET("/peraturan/:id", servicepemerintahgo.Peraturan)
		Pemerintah.PUT("/peraturan/setting", servicepemerintahgo.Setting_Peraturan)

		Pemerintah.GET("/kepdes/:id", servicepemerintahgo.Kepaladesa)
		Pemerintah.POST("/kepdes/tambah", servicepemerintahgo.Tambah_kepalaDesa)

		Pemerintah.GET("/organigram/:id", servicepemerintahgo.Organigram)
		Pemerintah.PUT("/organigram/setting", servicepemerintahgo.Setting_organigram)
	}

	Peraturan := router.Group("/peraturan")
	{
		Peraturan.GET("/", serviceperaturandesa.Peraturan_desa)
	}

	Pegawain := router.Group("/pegawai")
	{
		Pegawain.POST("/", servicepegawai.List)
		Pegawain.DELETE("/delete/:id", servicepegawai.Delete)
		Pegawain.POST("/tambah", servicepegawai.Tambah)
	}

	Hukum := router.Group("/hukum")
	{
		Hukum.POST("", serviceprodukhukum.List)
		Hukum.DELETE("/delete/:id", serviceprodukhukum.Delete)
		Hukum.POST("/tambah", serviceprodukhukum.Tambah)
		Hukum.GET("/kategori-law", serviceprodukhukum.List_Kategori)
	}

	Pembangunan := router.Group("/pembangunan")
	{
		Pembangunan.POST("", serviceinformasipembangunango.List)
		Pembangunan.DELETE("/delete/:id", serviceinformasipembangunango.Delete)
		Pembangunan.POST("/tambah", serviceinformasipembangunango.Tambah)
		Pembangunan.GET("/kategori-law", serviceprodukhukum.List_Kategori)
	}

	Wilayah_desa := router.Group("/wilayah_desa")
	{
		Wilayah_desa.GET("/desa/:id", servicewilayahdesa.Wilayah_desa)
		Wilayah_desa.POST("/setting", servicewilayahdesa.Setting)
	}

	Produksi_desa := router.Group("/produksi")
	{
		Produksi_desa.GET("/list/:id", serviceproduksi.Wilayah_produksi)
	}

	warga := router.Group("/warga")
	{
		warga.POST("/list", sevicewargadesabyamin.Warga_desa_by_admin)
		warga.DELETE("/delete/:id", sevicewargadesabyamin.Delete_warga_desa)
		warga.GET("/detail/:id", sevicewargadesabyamin.Get_detail_warga_by_admin)
		warga.POST("/tambah", sevicewargadesabyamin.Create_warga)
		warga.PUT("/update", sevicewargadesabyamin.Update_data_warga)
		warga.GET("/financial", sevicewargadesabyamin.Kategori_financial)
		warga.GET("/pendidikan", sevicewargadesabyamin.Kategori_pendidikan)
		warga.GET("/agama", sevicewargadesabyamin.Kategori_agama)
		warga.GET("/jk", sevicewargadesabyamin.JK_pengguna)
		warga.GET("/pernikahan", sevicewargadesabyamin.Pernikahan)

	}

	idmiksike := router.Group("/idmiksike")
	{
		idmiksike.POST("/input", serviceidm.Input)
	}

}
