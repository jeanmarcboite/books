package calibre

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type NullString sql.NullString
type NameSort struct {
	Name string
	Sort string
}

func GetNamesSort(database *sqlx.DB, from string) (map[uint](*NameSort), error) {
	rows, err := database.Queryx("select * from " + from)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	names := map[uint](*NameSort){}
	for rows.Next() {
		var ID uint
		var name string
		var sort sql.NullString
		err = rows.Scan(&ID, &name, &sort)
		if err != nil {
			return names, err
		}
		val := NameSort{Name: name}
		if sort.Valid {
			val.Sort = sort.String
		}

		names[ID] = &val
	}

	return names, err
}

func GetNames(db *CalibreDB, database *sqlx.DB, what string, from string, appendName func(uint, *NameSort)) error {
	names, err := GetNamesSort(database, from)
	if err == nil {
		return GetBooksNamesLink(db, database, what, from, names, appendName)
	}

	return err
}

func GetBooksNamesLink(
	db *CalibreDB,
	database *sqlx.DB,
	what string,
	from string,
	names map[uint](*NameSort),
	appendName func(uint, *NameSort)) error {
	if db.Books == nil {
		return nil
	}
	rows, err := database.Queryx("select book, " + what + " from books_" + from + "_link")

	if err == nil {
		defer rows.Close()
		var book uint
		var name uint
		for rows.Next() {
			err = rows.Scan(&book, &name)
			if err != nil {
				return err
			}
			appendName(book, names[name])
		}
	}
	return err
}

func GetPublishers(db *CalibreDB, database *sqlx.DB) error {
	return GetNames(db, database, "publisher", "publishers", func(bookID uint, data *NameSort) {
		if book, ok := db.Books[bookID]; ok {
			book.Publishers = append(book.Publishers, data)
		} else {
			log.Error().Uint("id", bookID).Msg("Invalid book id")
		}
	})
}

func GetSeries(db *CalibreDB, database *sqlx.DB) error {
	return GetNames(db, database, "series", "series", func(bookID uint, data *NameSort) {
		if book, ok := db.Books[bookID]; ok {
			book.Series = append(book.Series, data)
		} else {
			log.Error().Uint("id", bookID).Msg("Invalid book id")
		}
	})
}
