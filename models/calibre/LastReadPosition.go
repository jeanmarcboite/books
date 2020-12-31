package calibre

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type LastReadPosition struct {
	ID      uint
	Book    uint
	Format  string
	User    string
	Device  string
	Cfi     string
	Epoch   float64
	PosFrac float64 `db:"pos_frac"`
}

func (this LastReadPosition) Add(db *CalibreDB) {
	if book, ok := db.Books[this.ID]; ok {
		book.LastReadPosition = this
	} else {
		log.Error().Uint("id", this.ID).Msg("Invalid book id")
	}
}

func (this LastReadPosition) StructScan(rows *sqlx.Rows) (TableRowData, error) {
	err := rows.StructScan(&this)

	return this, err
}

func GetLastReadPositions(db *CalibreDB, database *sqlx.DB) error {
	var comment Comment

	return getTable(db, database, "comments", comment, func(rows *sqlx.Rows) error {
		return rows.StructScan(&comment)
	})
}
