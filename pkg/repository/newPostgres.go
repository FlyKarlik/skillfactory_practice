package repository

import "github.com/jmoiron/sqlx"

func NewPostgresDb(databasePostgresURL string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", databasePostgresURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
