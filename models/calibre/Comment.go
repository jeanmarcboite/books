package calibre

import (
	"github.com/jmoiron/sqlx"
)

type Comment struct {
	ID   uint
	Book uint
	Text string
}

func (this Comment) Add(db *CalibreDB) {
	db.Books[this.Book].Comment = this.Text
}

func (this Comment) StructScan(rows *sqlx.Rows) (TableRowData, error) {
	err := rows.StructScan(&this)

	return this, err
}

func GetComments(db *CalibreDB, database *sqlx.DB) error {
	var comment Comment

	return getTable(db, database, "comments", comment, func(rows *sqlx.Rows) error {
		return rows.StructScan(&comment)
	})
}
