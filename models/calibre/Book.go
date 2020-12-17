package calibre

import (
	"time"
)

type Book struct {
	ID           uint
	Title        string
	Sort         string
	SeriesIndex  float64
	AuthorSort   string
	Isbn         string
	Lccn         string
	Path         string
	Flags        int
	Uuid         string
	HasCover     bool
	Timestamp    time.Time
	Pubdate      time.Time
	LastModified time.Time
	Comment      string
}

type Comment struct {
	ID   uint
	Book uint
	Text string
}
