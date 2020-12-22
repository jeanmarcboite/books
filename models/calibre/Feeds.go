package calibre

import (
	"github.com/jmoiron/sqlx"
)

type Feed struct {
	ID     uint
	Title  string
	Script string
}

func (this Feed) Add(db *CalibreDB) {
	db.Feeds[this.ID] = this
}

func (this Feed) StructScan(rows *sqlx.Rows) (TableRowData, error) {
	err := rows.StructScan(&this)

	return this, err
}

func GetFeeds(db *CalibreDB, database *sqlx.DB) error {
	var feed Feed
	db.Feeds = make(map[uint]Feed)
	return getTable(db, database, "feeds", feed, func(rows *sqlx.Rows) error {
		return rows.StructScan(&feed)
	})
}
