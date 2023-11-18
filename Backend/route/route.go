package route

import (
	serviceberita "backendpgx7071/service/serviceBerita"
	serviceumkm "backendpgx7071/service/serviceUMKM"
	servicewisata "backendpgx7071/service/serviceWisata"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

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

}
