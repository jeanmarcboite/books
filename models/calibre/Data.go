package calibre

import (
	"github.com/jmoiron/sqlx"
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

		db.Books[dataRow.Book].Data =
			Data{
				Format:           dataRow.Format,
				UncompressedSize: dataRow.Uncompressed_size,
				Name:             dataRow.Name}
	}

	return rows.Err()
}
