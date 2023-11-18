package main

import (
	"backendpgx7071/database"
	"backendpgx7071/route"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Atur header untuk mengizinkan akses dari origin yang sesuai (http://127.0.0.1:8000)
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		// Izinkan cookies saat permintaan di-cross-origin
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// Atur metode HTTP yang diizinkan
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		// Atur header yang diizinkan termasuk 'x-requested-with'
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, x-requested-with")

		// Tangani permintaan OPTIONS
		if c.Request.Method == "OPTIONS" {
			// Jika permintaan adalah OPTIONS, balas dengan status 204 (No Content)
			c.AbortWithStatus(204)
			return
		}

		// Lanjutkan ke middleware selanjutnya
		c.Next()
	}
}

func main() {

	database.NewConnect()

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.Use(CORSMiddleware())

	route.Routes(router)

	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	gin.SetMode(gin.ReleaseMode)

	server := &http.Server{
		Addr:         ":7071",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second, // Ini adalah timeout untuk koneksi terjaga.
	}

	fmt.Println("Dijalankn pada port :7071")

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}
