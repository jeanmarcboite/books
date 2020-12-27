package calibre

import (
	"github.com/jmoiron/sqlx"
)

type Ratings struct {
	ID     uint
	Rating uint
}

func (this Ratings) Add(db *CalibreDB) {
	db.Books[this.ID].Rating = this.Rating
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
