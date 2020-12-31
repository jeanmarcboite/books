package calibre

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type dataRow struct {
	ID                uint
	Book              uint
	Format            string
	Uncompressed_size uint
	Name              string
}

type Data struct {
	Format           string
	UncompressedSize uint
	Name             string
}

func GetData(db *CalibreDB, database *sqlx.DB) error {
	rows, err := database.Queryx("select * from data")

	if err != nil {
		return err
	}

	defer rows.Close()
	for rows.Next() {
		var dataRow dataRow
		err = rows.StructScan(&dataRow)
		if err != nil {
			return err
		}
		if book, ok := db.Books[dataRow.Book]; ok {
			book.Data =
				Data{
					Format:           dataRow.Format,
					UncompressedSize: dataRow.Uncompressed_size,
					Name:             dataRow.Name}
		} else {
			log.Error().Uint("id", dataRow.Book).Msg("Invalid book id")
		}
	}

	return rows.Err()
}
