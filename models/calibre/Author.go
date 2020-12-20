package calibre

import (
	"encoding/json"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Author struct {
	ID   uint
	Name string
	Sort string
	Link string

	Books [](*Book)
}

type BookAuthorLink struct {
	ID     uint
	Book   uint
	Author uint
}

func (this Author) MarshalJSON() ([]byte, error) {
	return json.Marshal(this.Name)
}

func (this CalibreDB) ReadAuthors(database *sqlx.DB) error {
	this.Authors = make(map[uint](*Author))
	rows, err := database.Queryx("select * from authors")
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			author := new(Author)
			err = rows.StructScan(author)
			if err != nil {
				return err
			}
			fmt.Println(author)
			this.Authors[author.ID] = author
		}
		err = this.GetBooksAuthorsLink(database)

		if err == nil {
			return rows.Err()
		}
	}

	return err
}

func (this CalibreDB) GetBooksAuthorsLink(database *sqlx.DB) error {
	if (this.Authors == nil) || (this.Books == nil) {
		return nil
	}
	rows, err := database.Queryx("select * from books_authors_link")

	if err == nil {
		defer rows.Close()
		var link BookAuthorLink
		for rows.Next() {
			err = rows.StructScan(&link)
			fmt.Println(link)
			if err != nil {
				return err
			}
			this.Authors[link.Author].Books = append(this.Authors[link.Author].Books, this.Books[link.Book])
			this.Books[link.Book].Authors = append(this.Books[link.Book].Authors, this.Authors[link.Author])
		}
	}
	return err
}
