package servicelogin

import (
	"context"

	"github.com/jackc/pgx/v4"
)

func updateTokenInDatabase(userID int, newToken string) error {
	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `
        UPDATE dev.pengguna
        SET token = $1
        WHERE id_pengguna = $2
    `

	_, err = tx.Exec(ctx, updateQuery, newToken, userID)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}
