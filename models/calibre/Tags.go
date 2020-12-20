package calibre

import (
	"github.com/jmoiron/sqlx"
)

func (this *CalibreDB) GetTags(database *sqlx.DB) error {
	tags, err := this.GetStrings(database, "tags")
	if err == nil {
		return this.GetBooksTagsLink(database, tags)
	}

	return err
}

func (this CalibreDB) GetBooksTagsLink(database *sqlx.DB, tags map[uint]string) error {
	if this.Books == nil {
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
			this.Books[book].Tags = append(this.Books[book].Tags, tags[tag])
		}
	}
	return err
}
