package models

import (
	"github.com/jeanmarcboite/epub"
)

// Info -- Book info and metadata
type Work struct {
	Metadata
	Online map[string]Metadata
	Epub   *epub.EpubReaderCloser
	Error  error
}
