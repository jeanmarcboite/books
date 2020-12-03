package online

import (
	"github.com/jeanmarcboite/books/models"
	"github.com/jeanmarcboite/books/online/providers"
	"github.com/jeanmarcboite/books/online/providers/goodreads"
	"github.com/jeanmarcboite/books/online/providers/openlibrary"
)

// LookUpISBN -- lookup a work on goodreads and openlibrary, with isbn
func LookUpISBN(isbn string) (map[string]models.Metadata, error) {
	metadata := make(map[string]models.Metadata)

	// LibraryThing returns "APIs Temporarily disabled."

	for name, provider := range providers.Providers {
		m, err := provider.LookUpISBN(isbn)
		if err == nil {
			metadata[name] = m
		}
	}

	return metadata, nil
}

// SearchTitle --
func LookUpTitle(title string) ([]map[string]models.Metadata, error) {
	docs, err := openlibrary.LookUpTitle(title)
	if err != nil {
		return nil, err
	}

	books := make([]map[string]models.Metadata, len(docs))
	for k, doc := range docs {
		metadata := make(map[string]models.Metadata)
		if err == nil {
			metadata["openlibrary"] = doc
			if doc.ISBN != "" {
				g, err := goodreads.LookUpISBN(doc.ISBN)
				if err == nil {
					metadata["goodreads"] = g
				}

			}
			books[k] = metadata
		}
	}

	return books, err
}
