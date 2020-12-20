package calibre

import (
	"github.com/jmoiron/sqlx"
)

func GetComments(db *CalibreDB, database *sqlx.DB) error {
	rows, err := database.Queryx("select * from comments")

	if err != nil {
		return err
	}

	defer rows.Close()
	for rows.Next() {
		var comment Comment
		err = rows.StructScan(&comment)
		if err != nil {
			return err
		}

		db.Books[comment.Book].Comment = comment.Text
	}

	return rows.Err()
}
