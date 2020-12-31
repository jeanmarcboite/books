package calibre

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type Ratings struct {
	ID     uint
	Rating uint
}

func (this Ratings) Add(db *CalibreDB) {
	if book, ok := db.Books[this.ID]; ok {
		book.Rating = this.Rating
	} else {
		log.Error().Uint("id", this.ID).Msg("Invalid book id")
	}
}

func (this Ratings) StructScan(rows *sqlx.Rows) (TableRowData, error) {
	err := rows.StructScan(&this)

	return this, err
}

func GetRatings(db *CalibreDB, database *sqlx.DB) error {
	var ratings Ratings

	return getTable(db, database, "ratings", ratings, func(rows *sqlx.Rows) error {
		return rows.StructScan(&ratings)
	})
}
