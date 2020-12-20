package calibre

import (
	"github.com/jmoiron/sqlx"
)

type Identifier struct {
	Type string
	Val  string
}

func (this *CalibreDB) GetIdentifiers(database *sqlx.DB) error {
	rows, err := database.Queryx("select book, type, val from identifiers")

	if err != nil {
		return err
	}

	defer rows.Close()

	var book uint
	for rows.Next() {
		var identifier Identifier

		err = rows.Scan(&book, &identifier.Type, &identifier.Val)
		if err != nil {
			return err
		}

		this.Books[book].Identifiers = append(this.Books[book].Identifiers, identifier)
	}

	return rows.Err()
}
