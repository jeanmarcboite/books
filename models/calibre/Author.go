package calibre

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
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

			author, authorOK := db.Authors[link.Author]
			book, bookOK := db.Books[link.Book]

			if bookOK && authorOK {
				author.Books = append(author.Books, link.Book)
				book.Authors = append(book.Authors, link.Author)

			} else {
				if !authorOK {
					log.Error().Uint("id", link.Author).Msg("Invalid author id")

				}
				if !bookOK {
					log.Error().Uint("id", link.Book).Msg("Invalid book id")
				}
			}
		}
	}
	return err
}
