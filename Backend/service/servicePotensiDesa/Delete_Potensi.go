package servicepotensidesa

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Delete_potensi(c *gin.Context) {
	id := c.Param("id")

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}
	defer tx.Rollback(ctx)

	query := `DELETE FROM dev.potensi_desa WHERE id_potensi = $1`

	result, err := tx.Exec(ctx, query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": false, "message": "POTENSI not found"})
		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "POTENSI deleted successfully"})

}
