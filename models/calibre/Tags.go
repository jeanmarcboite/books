package calibre

import (
	"github.com/jmoiron/sqlx"
)

func GetTags(db *CalibreDB, database *sqlx.DB) error {
	tags, err := db.GetStrings(database, "tags")
	if err == nil {
		return GetBooksTagsLink(db, database, tags)
	}

	return err
}

func GetBooksTagsLink(db *CalibreDB, database *sqlx.DB, tags map[uint]string) error {
	if db.Books == nil {
		return nil
	}
	rows, err := database.Queryx("select book, tag from books_tags_link")

	if err == nil {
		defer rows.Close()
		var book uint
		var tag uint
		for rows.Next() {
			err = rows.Scan(&book, &tag)
			if err != nil {
				return err
			}
			db.Books[book].Tags = append(db.Books[book].Tags, tags[tag])
		}
	}
	return err
}
