package calibre

import (
	"time"
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
