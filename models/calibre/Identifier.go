package calibre

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type Identifier struct {
	Type string
	Val  string
}

func GetIdentifiers(db *CalibreDB, database *sqlx.DB) error {
	rows, err := database.Queryx("select book, type, val from identifiers")

	if err != nil {
		return err
	}

	defer rows.Close()

	var bookID uint
	for rows.Next() {
		var identifier Identifier

		err = rows.Scan(&bookID, &identifier.Type, &identifier.Val)
		if err != nil {
			return err
		}
		if book, ok := db.Books[bookID]; ok {
			book.Identifiers = append(book.Identifiers, identifier)
		} else {
			log.Error().Uint("id", bookID).Msg("Invalid book id")
		}
	}

	return rows.Err()
}
