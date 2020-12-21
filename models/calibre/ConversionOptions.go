package calibre

import (
	"github.com/jmoiron/sqlx"
)

type ConversionOptions struct {
	Format string
	Data   []byte
}

func GetConversionOptions(db *CalibreDB, database *sqlx.DB) error {
	rows, err := database.Queryx("select * from conversion_options")

	if err != nil {
		return err
	}

	defer rows.Close()
	for rows.Next() {
		var id, book uint
		var format string
		var data []byte
		err = rows.Scan(&id, &format, &book, &data)
		if err != nil {
			return err
		}

		db.Books[book].ConversionOptions = ConversionOptions{Format: format, Data: data}
	}

	return rows.Err()
}
