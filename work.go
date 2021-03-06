package books

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/jeanmarcboite/books/models"
	"github.com/jeanmarcboite/books/online"
	"github.com/jeanmarcboite/books/online/net"
	"github.com/jeanmarcboite/epub"
)

// WorkFromISBN -- look online for a book
func WorkFromISBN(isbn string) (models.Work, error) {
	metadata, _ := online.LookUpISBN(isbn)

	return work(metadata, nil)
}

// WorksFromTitle -- look online for a book
func WorksFromTitle(title string) ([]models.Work, error) {
	log.Debug().Str("title", title).Msg("WorksFromTitle")
	metadata, _ := online.LookUpTitle(title)
	books := make([]models.Work, len(metadata))
	for k, book := range metadata {
		books[k], _ = work(book, nil)
	}
	return books, nil
}

// WorkFromFilename -- read file, look up online for metadata
func WorkFromFilename(filename string) (models.Work, error) {
	log.Debug().Str("filename", filename).Msg("WorksFromFilename")
	ereader, error := epub.OpenReader(filename)

	if error != nil {
		return models.Work{Error: error}, error
	}

	work, error := workFromEpub(ereader)
	ereader.Close()

	return work, error
}

// WorkFromData -- read epub, look up online for metadata
func WorkFromData(zipData []byte, size int64) (models.Work, error) {
	ereader, error := epub.OpenBuffer(zipData, size)

	if error != nil {
		return models.Work{Error: error}, error
	}

	work, error := workFromEpub(ereader)

	return work, error
}

func workFromEpub(epub *epub.EpubReaderCloser) (models.Work, error) {
	isbn, err := epub.GetISBN()
	if err != nil {
		return models.Work{Epub: epub}, nil
	}
	metadata, _ := online.LookUpISBN(isbn)

	return work(metadata, epub)
}

func work(metadata map[string]models.Metadata, epub *epub.EpubReaderCloser) (models.Work, error) {
	this := models.Work{Online: metadata, Epub: epub}
	this.URL = make(map[string]string)

	if epub != nil {
		epubMetadata := epub.Rootfiles[0].Metadata
		this.ID = "epub"
		this.Title = epubMetadata.Title
		if epubMetadata.Creator.Role == "aut" {
			this.Author = epubMetadata.Creator.Text
		}

		if len(epubMetadata.Publisher) > 0 {
			this.Publishers = []string{epubMetadata.Publisher}
		}
		this.Description = epubMetadata.Description
		// TODO: identifier
		this.Cover, _ = epub.GetCover()
	}

	for online := range metadata {
		//printFieldNames(metadata[online])
		mo := metadata[online]
		s := reflect.ValueOf(&mo).Elem()
		t := s.Type()

		for i := 0; i < s.NumField(); i++ {
			assign(&this, online, t.Field(i).Name)
		}
		if metadata[online].Link != "" {
			this.URL[online] = metadata[online].Link
		} else if metadata[online].ID != "" {
			if net.Koanf.Get(online+".url.show") != nil {
				this.URL[online] = fmt.Sprintf(net.Koanf.String(online+".url.show"), metadata[online].ID)
			}
		}
	}

	if this.Cover == "" || strings.Contains(this.Cover, "nophoto") {
		if online, ok := this.Online["google"]; ok {
			this.Cover = online.Cover
		}
	}

	if this.Cover == "" || strings.Contains(this.Cover, "nophoto") {
		this.Cover = fmt.Sprintf(net.Koanf.String("librarything.url.cover"),
			net.Koanf.String("librarything.key"), this.ISBN)
	}

	this.Author = this.GetAuthors()
	if this.RatingsPercent == "" && this.RatingsSum > 0 {
		this.RatingsPercent = fmt.Sprintf("%6.2f",
			float64(this.RatingsSum)/float64(this.RatingsCount))
	}

	return this, nil
}

func display(this *models.Work, key string) {
	type T struct {
		A int
		B string
	}
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}

func printFieldNames(this models.Metadata) {
	s := reflect.ValueOf(&this).Elem()
	t := s.Type()
	for i := 0; i < s.NumField(); i++ {
		fmt.Printf("%d: %s\n", i, t.Field(i).Name)
	}
}

func assign(this *models.Work, key string, fieldName string) {
	if fieldName != "RAW" && fieldName != "XML" {
		value := reflect.ValueOf(this.Online[key]).FieldByName(fieldName)
		field := reflect.ValueOf(this).Elem().FieldByName(fieldName)

		if field.IsZero() {
			// A Value can be changed only if it is
			// addressable and was not obtained by
			// the use of unexported struct fields.
			if field.IsValid() && field.CanSet() {
				field.Set(value)
				/**
				if field.Kind() == reflect.String {
					field.SetString(value.String())
				}
				**/
			}
		}
		if false {
			if fieldName == "Title" || fieldName == "Authors" {
				fmt.Println(fieldName, field, field.String(), field.Kind(), value.String())
			}
		}
	}
}
