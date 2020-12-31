package calibre

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type ConversionOption struct {
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
		var id, bookID uint
		var format string
		var data []byte
		err = rows.Scan(&id, &format, &bookID, &data)
		if err != nil {
			return err
		}
		if book, ok := db.Books[bookID]; ok {
			book.ConversionOption = ConversionOption{Format: format, Data: data}
		} else {
			log.Error().Uint("id", bookID).Msg("Invalid book id")
		}
	}

	return rows.Err()
}
