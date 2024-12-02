package db

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type DbPlugin struct {
	db *sqlx.DB
}

func Config(databaseType string, connectionStr string) DbPlugin {
	db, err := sqlx.Connect(databaseType, connectionStr)
	if err != nil {
		log.Fatalln(err)
	}

	return DbPlugin{
		db,
	}
}
