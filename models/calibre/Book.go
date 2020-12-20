package calibre

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Book struct {
	ID           uint
	Title        string
	Sort         string
	SeriesIndex  float64 `db:"series_index"`
	AuthorSort   string  `db:"author_sort"`
	Isbn         string
	Lccn         string
	Path         string
	Flags        int
	Uuid         string
	HasCover     bool `db:"has_cover"`
	Timestamp    time.Time
	Pubdate      time.Time
	LastModified time.Time `db:"last_modified"`
	Comment      string

	Authors [](*Author)
}

type Comment struct {
	ID   uint
	Book uint
	Text string
}

func (this *CalibreDB) ReadBooks(database *sqlx.DB) error {
	this.Books = make(map[uint](*Book))
	rows, err := database.Queryx("select * from books")

	if err == nil {
		defer rows.Close()
		for rows.Next() {
			book := new(Book)
			err = rows.StructScan(book)
			if err != nil {
				return err
			}
			this.Books[book.ID] = book
		}

		err = this.GetComments(database)
		if err != nil {
			return err
		}
		err = this.ReadAuthors(database)

		if err == nil {
			err = rows.Err()
		}
	}

	return err
}

func (this *CalibreDB) GetComments(database *sqlx.DB) error {
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

		this.Books[comment.Book].Comment = comment.Text
	}

	return rows.Err()
}
