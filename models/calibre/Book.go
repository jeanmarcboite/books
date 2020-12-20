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

	Authors     [](*Author)
	Annotations []Annotation
	Data        Data
	Identifiers []Identifier
	Languages   []Language
	Tags        []string
	Publishers  [](*NameSort)
	Series      [](*NameSort)
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

		type Get func(*CalibreDB, *sqlx.DB) error

		getFunctions := []Get{
			GetComments,
			GetIdentifiers,
			GetLanguages,
			GetTags, GetAuthors, GetPublishers,
			GetSeries, GetAnnotations, GetData}

		for _, f := range getFunctions {
			err = f(this, database)
			if err != nil {
				return err
			}
		}
		err = rows.Err()
	}

	return err
}
