package serviceberita

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type Data_berita_admin struct {
	ID         int    `json:"id_berita"`
	Judul      string `json:"judul"`
	Deskripsi  string `json:"deskripsi"`
	FotoBerita string `json:"foto_berita"`
	Kategori   string `json:"kategori"`
}

// CreateBerita creates a new berita
func CreateBerita(c *gin.Context) {
	var inputBerita Data_berita_admin
	if err := c.ShouldBindJSON(&inputBerita); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid input"})
		return
	}

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}
	defer tx.Rollback(ctx)

	var id_desa int
	desa := `select id_desa from dev.desa d where nama_desa = 'Morowali'`
	err = tx.QueryRow(ctx, desa).Scan(&id_desa)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	query := `
		INSERT INTO dev.berita (judul, deskripsi, foto_berita, kategori_id, desa_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id_berita
	`

	var idBerita int
	err = tx.QueryRow(ctx, query, inputBerita.Judul, inputBerita.Deskripsi, inputBerita.FotoBerita, inputBerita.Kategori, id_desa).Scan(&idBerita)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Berita created successfully", "id_berita": idBerita})
	err = tx.Commit(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
	}
}

// GetBerita retrieves a specific berita by ID
func GetBerita(c *gin.Context) {
	idBerita := c.Param("id")

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}
	defer tx.Rollback(ctx)

	// Modify the SQL query based on your database schema
	query := `
		SELECT 
			a.id_berita,
			a.judul,
			a.deskripsi,
			a.foto_berita,
			b.kategori
		FROM
		dev.berita a, dev.kategori_berita b
		WHERE
			a.kategori_id = b.id_kategori_berita
			AND a.id_berita = $1
	`

	var result Data_berita_admin
	err = tx.QueryRow(ctx, query, idBerita).Scan(
		&result.ID,
		&result.Judul,
		&result.Deskripsi,
		&result.FotoBerita,
		&result.Kategori,
	)

	if err == pgx.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"status": false, "message": "Berita not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Berita retrieved successfully", "data": result})
	err = tx.Commit(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
	}
}

// UpdateBerita updates an existing berita
func UpdateBerita(c *gin.Context) {
	idBerita := c.Param("id")

	var inputBerita Data_berita_admin
	if err := c.ShouldBindJSON(&inputBerita); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid input"})
		return
	}

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}
	defer tx.Rollback(ctx)

	// Modify the SQL query based on your database schema
	query := `
		UPDATE dev.berita
		SET judul = $1, deskripsi = $2, foto_berita = $3, kategori_id = $4
		WHERE id_berita = $5
	`

	_, err = tx.Exec(ctx, query, inputBerita.Judul, inputBerita.Deskripsi, inputBerita.FotoBerita, inputBerita.Kategori, idBerita)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Berita updated successfully"})
	err = tx.Commit(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
	}
}

// DeleteBerita deletes a berita by ID
func DeleteBerita(c *gin.Context) {
	idBerita := c.Param("id")

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}
	defer tx.Rollback(ctx)

	// Modify the SQL query based on your database schema
	query := `DELETE FROM dev.berita WHERE id_berita = $1`

	result, err := tx.Exec(ctx, query, idBerita)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": false, "message": "Berita not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Berita deleted successfully"})
	err = tx.Commit(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
	}
}
