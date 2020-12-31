package calibre

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
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
		var bookID uint
		var tagID uint
		for rows.Next() {
			err = rows.Scan(&bookID, &tagID)
			if err != nil {
				return err
			}
			tag, tagOK := tags[tagID]
			book, bookOK := db.Books[bookID]
			if tagOK && bookOK {
				book.Tags = append(book.Tags, tag)
			} else {
				if !tagOK {
					log.Error().Uint("id", tagID).Msg("Invalid tag id")

				}
				if !bookOK {
					log.Error().Uint("id", bookID).Msg("Invalid book id")
				}
			}
		}
	}
	return err
}
