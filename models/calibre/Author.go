package calibre

import (
	"github.com/jmoiron/sqlx"
)

type Author struct {
	ID   uint
	Name string
	Sort string
	Link string

	Books []uint
}

type BookAuthorLink struct {
	ID     uint
	Book   uint
	Author uint
}

/*
func (this Author) MarshalJSON() ([]byte, error) {
	return json.Marshal(this.Name)
}
*/
func GetAuthors(db *CalibreDB, database *sqlx.DB) error {
	db.Authors = make(map[uint](*Author))
	rows, err := database.Queryx("select * from authors")
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			author := new(Author)
			err = rows.StructScan(author)
			if err != nil {
				return err
			}
			db.Authors[author.ID] = author
		}
		err = GetBooksAuthorsLink(db, database)
		if err == nil {
			return rows.Err()
		}
	}

	return err
}

func GetBooksAuthorsLink(db *CalibreDB, database *sqlx.DB) error {
	if db.Books == nil {
		return nil
	}
	rows, err := database.Queryx("select * from books_authors_link")

	if err == nil {
		defer rows.Close()
		var link BookAuthorLink
		for rows.Next() {
			err = rows.StructScan(&link)
			if err != nil {
				return err
			}
			db.Authors[link.Author].Books = append(db.Authors[link.Author].Books, link.Book)
			db.Books[link.Book].Authors = append(db.Books[link.Book].Authors, link.Author)
		}
	}
	return err
}
