package calibre

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type Comment struct {
	ID   uint
	Book uint
	Text string
}

func (this Comment) Add(db *CalibreDB) {
	if book, ok := db.Books[this.Book]; ok {
		book.Comment = this.Text
	} else {
		log.Error().Uint("id", this.Book).Msg("Invalid book id")
	}
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
