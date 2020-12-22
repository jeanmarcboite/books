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

	Authors           [](*Author)
	Annotations       []Annotation
	Data              Data
	ConversionOptions ConversionOptions
	Identifiers       []Identifier
	Languages         []Language
	Tags              []string
	Publishers        [](*NameSort)
	Series            [](*NameSort)
	LastReadPosition  LastReadPosition
}

func (this Book) Add(db *CalibreDB) {
	book := this
	db.Books[this.ID] = &book
}

func (this Book) StructScan(rows *sqlx.Rows) (TableRowData, error) {
	err := rows.StructScan(&this)

	return this, err
}

func (this *CalibreDB) ReadBooks(database *sqlx.DB) error {
	var book Book
	this.Books = make(map[uint]*Book)

	err := getTable(this, database, "books", book, func(rows *sqlx.Rows) error {
		return rows.StructScan(&book)
	})

	if err == nil {
		type Get func(*CalibreDB, *sqlx.DB) error

		getFunctions := []Get{
			GetComments,
			GetIdentifiers,
			GetLanguages,
			GetTags, GetAuthors, GetPublishers,
			GetSeries, GetAnnotations, GetData,
			GetConversionOptions, GetFeeds,
			GetLastReadPositions}

		for _, f := range getFunctions {
			err = f(this, database)
			if err != nil {
				return err
			}
		}
	}

	return err
}
