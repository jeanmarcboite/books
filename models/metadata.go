package models

import (
	"strings"
)

// Metadata -- book metadata
type Metadata struct {
	ID             string
	Title          string
	SubTitle       string
	Author         string
	FirstAuthor    Author
	Authors        []Author
	Categories     []string
	Series         string
	Tags           string
	Ratings        string
	RatingsPercent string
	ReviewsCount   int
	RatingsSum     int
	RatingsCount   int
	RatingDist     string
	URL            map[string]string
	Cover          string
	Covers         []string
	Identifiers    Identifiers
	PublishDate    string
	Publishers     []string
	PublishCountry string
	Description    string
	Subjects       string
	PopularShelves []Shelf
	NumberOfPages  int
	Preview        string
	PhysicalFormat string
	IsEbook        string
	LanguageCode   string
	Legal          string

	ISBN string

	RAW interface{}
	XML string
}

// Identifiers -- book identifiers
type Identifiers struct {
	ISBN13       []string
	ISBN10       []string
	Amazon       []string
	ASIN         string
	KindleASIN   string
	Google       []string
	Gutenberg    []string
	Goodreads    []string
	Librarything []string
}

type Shelf struct {
	// Text  string `xml:",chardata"`
	Name  string `xml:"name,attr"`
	Count string `xml:"count,attr"`
}

// GetAuthors -- return the author(s)
func (m Metadata) GetAuthors() string {
	authors := make([]string, len(m.Authors))
	for k, author := range m.Authors {
		authors[k] = author.Name
	}

	return strings.Join(authors, ", ")
}
