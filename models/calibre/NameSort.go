package calibre

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type NameSort struct {
	Name string
	Sort sql.NullString
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
		var val = new(NameSort)
		err = rows.Scan(&ID, &val.Name, &val.Sort)
		if err != nil {
			return names, err
		}

		names[ID] = val
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
			fmt.Println(book, name)
			appendName(book, names[name])
		}
	}
	return err
}

func GetPublishers(db *CalibreDB, database *sqlx.DB) error {
	return GetNames(db, database, "publisher", "publishers", func(book uint, data *NameSort) {
		db.Books[book].Publishers = append(db.Books[book].Publishers, data)
	})
}

func GetSeries(db *CalibreDB, database *sqlx.DB) error {
	return GetNames(db, database, "series", "series", func(book uint, data *NameSort) {
		db.Books[book].Series = append(db.Books[book].Series, data)
	})
}
