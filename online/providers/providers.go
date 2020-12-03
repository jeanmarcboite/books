package providers

import (
	"github.com/jeanmarcboite/books/models"
	"github.com/jeanmarcboite/books/online/providers/goodreads"
	"github.com/jeanmarcboite/books/online/providers/google"
	"github.com/jeanmarcboite/books/online/providers/openlibrary"
)

type ProviderLookUpISBN func(isbn string) (models.Metadata, error)

type Provider struct {
	LookUpISBN   func(isbn string) (models.Metadata, error)
	SearchAuthor func(author string) (models.Author, error)
}

var Providers = map[string]Provider{
	goodreads.Name():   {LookUpISBN: goodreads.LookUpISBN, SearchAuthor: goodreads.SearchAuthor},
	google.Name():      {LookUpISBN: google.LookUpISBN},
	openlibrary.Name(): {LookUpISBN: openlibrary.LookUpISBN},
}
