package testhelpers

import (
	"booking/services/auth/config"
	"testing"

	"github.com/jmoiron/sqlx"
)

func PreparePostgres(t *testing.T) (*sqlx.Tx, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	db, err := config.NewPostgresDb(cfg)
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
