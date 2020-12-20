package calibre

import (
	"database/sql"
	"encoding/json"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type CalibreDB struct {
	Books   map[uint](*Book)
	Authors map[uint](*Author)
}

func (this CalibreDB) String() string {
	res, _ := json.MarshalIndent(this, "", "   ")
	return string(res)
}

func ReadDB(filename string, debug bool) (CalibreDB, error) {
	db := CalibreDB{}
	var err error = nil

	sqlDB, err := sql.Open("sqlite3", filename)

	if err == nil {
		database := sqlx.NewDb(sqlDB, "sqlite3")
		defer database.Close()
		err = database.Ping()

		if err == nil {
			err = db.ReadBooks(database)
		}
	}

	return db, err
}
