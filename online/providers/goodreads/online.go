package goodreads

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"

	"github.com/jeanmarcboite/books/models"
	"github.com/jeanmarcboite/books/online/net"
	"github.com/rs/zerolog/log"
)

// LookUpISBN -- lookup a work on goodreads, with isbn
func LookUpISBN(isbn string) (models.Metadata, error) {
	return getMetadata(isbn, net.Koanf.String("goodreads.url.isbn"))
}

// SearchTitle -- search for a work with a title
func SearchTitle(title string) (models.Metadata, error) {
	return getMetadata(strings.Join(strings.Fields(title), "+"),
		net.Koanf.String("goodreads.url.title"))
}

func getResponse(what string, where string) ([]byte, error) {
	url := fmt.Sprintf(where, what)
	log.Debug().Str("where", where).Str("what", what).Str("url", url).Msg("getResponse")
	response, err := net.HTTPGetWithKey(url,
		net.Koanf.String("goodreads.keyname"),
		net.Koanf.String("goodreads.key"))
	if err != nil {
		log.Error().Str("url", url).Msg(err.Error())
		return nil, err
	}

	return response, nil
}

func getMetadata(what string, where string) (models.Metadata, error) {
	response, err := getResponse(what, where)
	if err != nil {
		return models.Metadata{}, err
	}

	var goodreads GoodreadsResponse

	/* response could be: <error>Page not found</error> */
	xml.Unmarshal(response, &goodreads)

	if goodreads.XMLName.Local == "GoodreadsResponse" {
		return parseBook(goodreads)
	}

	return models.Metadata{}, fmt.Errorf("Nothing found on goodreads for '%v'", what)
}

func parseBook(goodreads GoodreadsResponse) (models.Metadata, error) {
	reviewsCount, _ := strconv.Atoi(goodreads.Book.Work.ReviewsCount.Text)
	ratingsSum, _ := strconv.Atoi(goodreads.Book.Work.RatingsSum.Text)
	ratingsCount, _ := strconv.Atoi(goodreads.Book.Work.RatingsCount.Text)

	meta := models.Metadata{
		ID:      goodreads.Book.ID,
		Title:   goodreads.Book.Title,
		Authors: []models.Author{goodreads.Book.Authors.Author},
		Identifiers: models.Identifiers{
			ISBN10:     []string{goodreads.Book.ISBN},
			ISBN13:     []string{goodreads.Book.Isbn13},
			ASIN:       goodreads.Book.Asin,
			KindleASIN: goodreads.Book.KindleAsin,
		},
		PublishCountry: goodreads.Book.CountryCode,
		Publishers:     []string{goodreads.Book.Publisher},
		Description:    goodreads.Book.Description,
		Cover:          goodreads.Book.ImageURL,
		IsEbook:        goodreads.Book.IsEbook,
		ReviewsCount:   reviewsCount,
		RatingsSum:     ratingsSum,
		RatingsCount:   ratingsCount,
		RatingDist:     goodreads.Book.Work.RatingDist,

		RAW: goodreads,
	}

	return meta, nil
}

func SearchAuthor(author string) (models.Author, error) {
	response, err := getResponse(author, net.Koanf.String("goodreads.url.searchAuthor"))

	if err != nil {
		log.Error().Str("author", author).Msg("Cannot show author")
		return models.Author{}, err
	}

	var goodreads GoodreadsResponseShowAuthor
	xml.Unmarshal(response, &goodreads)
	if goodreads.XMLName.Local != "GoodreadsResponse" {
		log.Error().Str("author", author).Str("XMLName.Local", goodreads.XMLName.Local).Msg("Invalid goodreads response")
		return models.Author{}, err
	}

	return models.Author{
		ID:   goodreads.Author.ID,
		Name: goodreads.Author.Name,
		Link: goodreads.Author.Link,
	}, nil
}
