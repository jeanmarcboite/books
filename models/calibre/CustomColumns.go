package calibre

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type TableRowData interface {
	Add(db *CalibreDB)
	StructScan(rows *sqlx.Rows) (TableRowData, error)
}

type CustomColumn struct {
	ID            uint
	Label         string
	Name          string
	Datatype      string
	MarkForDelete bool `db:"mark_for_delete"`
	Editable      bool
	Display       string
	IsMultiple    bool `db:"is_multiple"`
	Normalized    bool
}

func (custom_column CustomColumn) Add(db *CalibreDB) {
	db.CustomColumns[custom_column.ID] = custom_column
}
func (custom_column CustomColumn) StructScan(rows *sqlx.Rows) (TableRowData, error) {
	err := rows.StructScan(&custom_column)

	return custom_column, err
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

func GetCustomColumns(db *CalibreDB, database *sqlx.DB) error {
	var custom_column CustomColumn
	db.CustomColumns = make(map[uint]CustomColumn)
	return getTable(db, database, "custom_columns", custom_column, func(rows *sqlx.Rows) error {
		return rows.StructScan(&custom_column)
	})
}
