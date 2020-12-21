package calibre

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type CalibreDB struct {
	Books         map[uint](*Book)
	Authors       map[uint](*Author)
	CustomColumns map[uint]CustomColumn
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
		if err == nil {
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
