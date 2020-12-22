package calibre

import (
	"github.com/jmoiron/sqlx"
)

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

func GetCustomColumns(db *CalibreDB, database *sqlx.DB) error {
	var custom_column CustomColumn
	db.CustomColumns = make(map[uint]CustomColumn)
	return getTable(db, database, "custom_columns", custom_column, func(rows *sqlx.Rows) error {
		return rows.StructScan(&custom_column)
	})
}
