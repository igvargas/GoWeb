package db

import (
	"database/sql"

	"github.com/DATA-DOG/go-txdb"
)

func init() {
	dataSource := "root:@tcp(localhost:3306)/dbGoWeb"
	txdb.Register("txdb", "mysql", dataSource)
}

func InitDb() (*sql.DB, error) {
	db, err := sql.Open("txdb", "indentificar")
	if err != nil {
		return db, err
	}
	return db, nil
}
