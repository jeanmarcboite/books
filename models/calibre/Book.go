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
	Authors      []uint
}

type Comment struct {
	ID   uint
	Book uint
	Text string
}

func ReadBooks(database *sqlx.DB) (CalibreDB, error) {
	db := CalibreDB{}
	db.Books = make(map[uint](*Book))
	rows, err := database.Queryx("select * from books")

	if err == nil {
		defer rows.Close()
		for rows.Next() {
			book := new(Book)
			err = rows.StructScan(book)
			if err != nil {
				return db, err
			}
			db.Books[book.ID] = book
		}

		GetComments(database, &db)

		return db, rows.Err()
	}

	return db, err
}

func GetComments(database *sqlx.DB, db *CalibreDB) error {
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
