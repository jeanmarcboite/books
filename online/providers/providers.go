package providers

import (
	"github.com/jeanmarcboite/books/models"
	"github.com/jeanmarcboite/books/online/providers/goodreads"
	"github.com/jeanmarcboite/books/online/providers/google"
	"github.com/jeanmarcboite/books/online/providers/openlibrary"
)

type ProviderLookUpISBN func(isbn string) (models.Metadata, error)

type Provider struct {
	Name         string
	LookUpISBN   func(isbn string) (models.Metadata, error)
	SearchAuthor func(author string) (models.Author, error)
}

var Providers = []Provider{
	{Name: goodreads.Name(), LookUpISBN: goodreads.LookUpISBN, SearchAuthor: goodreads.SearchAuthor},
	{Name: google.Name(), LookUpISBN: google.LookUpISBN},
	{Name: openlibrary.Name(), LookUpISBN: openlibrary.LookUpISBN},
}
