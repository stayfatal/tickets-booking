package testhelpers

import (
	"testing"
	"ticketsbooking/libs/config"

	"github.com/jmoiron/sqlx"
)

func PreparePostgres(t *testing.T) (*sqlx.Tx, error) {
	db, err := config.NewPostgresDB()
	if err != nil {
		return nil, err
	}

	tx, err := db.Beginx()
	if err != nil {
		return nil, err
	}

	t.Cleanup(func() {
		tx.Rollback()
		db.Close()
	})

	return tx, nil
}
