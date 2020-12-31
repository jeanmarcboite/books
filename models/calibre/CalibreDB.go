package calibre

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
)

type CalibreDB struct {
	ID            string
	Filename      string
	Books         map[uint]*Book
	Authors       map[uint](*Author)
	Ratings       map[uint]uint
	CustomColumns map[uint]CustomColumn
	Feeds         map[uint]Feed
}

type TableRowData interface {
	Add(db *CalibreDB)
	StructScan(rows *sqlx.Rows) (TableRowData, error)
}

func (this CalibreDB) String() string {
	res, _ := json.MarshalIndent(this, "", "   ")
	return string(res)
}

func ReadDB(filename string, debug bool) (CalibreDB, error) {
	log.Debug().Str("filename", filename).Msg("ReadDB")
	db := CalibreDB{Filename: filename}
	var err error = nil

	sqlDB, err := sql.Open("sqlite3", filename)

	if err == nil {
		log.Debug().Str("db", filename).Msg("open database")
		database := sqlx.NewDb(sqlDB, "sqlite3")
		defer database.Close()
		err = database.Ping()
		rows, err := database.Queryx("select uuid from library_id")
		if err == nil {
			defer rows.Close()
			for rows.Next() {
				err = rows.Scan(&db.ID)
				if err != nil {
					return db, fmt.Errorf("scanning library_id: %s", err.Error())
				}
				log.Debug().Str("db", filename).Str("uuid", db.ID).Msg("open database")
			}

			err = rows.Err()
		}

		if err == nil {
			log.Debug().Str("db", filename).Str("uuid", db.ID).Msg("read books")
			err = db.ReadBooks(database)
		}
		if err == nil {
			log.Debug().Str("db", filename).Str("uuid", db.ID).Msg("get custom columns")
			err = GetCustomColumns(&db, database)
		}
	}

	return db, err
}

func (this *CalibreDB) GetStrings(database *sqlx.DB, from string) (map[uint]string, error) {
	rows, err := database.Queryx("select * from " + from)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	strings := map[uint]string{}
	for rows.Next() {
		var ID uint
		var val string
		err = rows.Scan(&ID, &val)
		if err != nil {
			return strings, err
		}

		strings[ID] = val
	}
	return strings, err
}

func getTable(db *CalibreDB, database *sqlx.DB, table string, data TableRowData, StructScan func(rows *sqlx.Rows) error) error {
	rows, err := database.Queryx("select * from " + table)

	if err == nil {
		defer rows.Close()
		for rows.Next() {
			data, err = data.StructScan(rows)
			if err != nil {
				return fmt.Errorf("scanning table %s: %s", table, err.Error())
			}
			data.Add(db)
		}

		err = rows.Err()
	}

	return err
}
