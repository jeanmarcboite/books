package goodreads

import (
	"encoding/xml"
	"fmt"
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

	var goodreads Response

	/* response could be: <error>Page not found</error> */
	xml.Unmarshal(response, &goodreads)

	if goodreads.XMLName.Local == "GoodreadsResponse" {
		return parseBook(goodreads.Books[0])
	}

	return models.Metadata{}, fmt.Errorf("Nothing found on goodreads for '%v'", what)
}

func parseBook(goodreads Book) (models.Metadata, error) {
	meta := models.Metadata{
		ID:      goodreads.ID,
		Title:   goodreads.Title,
		Authors: []models.Author{},
		Identifiers: models.Identifiers{
			ISBN10:     []string{goodreads.ISBN},
			ISBN13:     []string{goodreads.ISBN13},
			ASIN:       goodreads.ASIN,
			KindleASIN: goodreads.KindleASIN,
		},
		PublishCountry: goodreads.CountryCode,
		Publishers:     []string{goodreads.Publisher},
		Description:    goodreads.Description,
		Cover:          goodreads.ImageURL,
		IsEbook:        goodreads.IsEbook,
		ReviewsCount:   goodreads.Work.ReviewsCount,
		RatingsSum:     goodreads.Work.RatingsSum,
		RatingsCount:   goodreads.Work.RatingsCount,
		Ratings:        goodreads.Work.RatingDist,

		RAW: goodreads,
	}

	return meta, nil
}

func SearchAuthor(author string) (models.Author, error) {
	fmt.Println(net.Koanf.String("goodreads.url.nocover"))
	fmt.Println(net.Koanf.String("goodreads.url.showAuthor"))
	_, err := getResponse(author, net.Koanf.String("goodreads.url.showAuthor"))

	if err != nil {
		return models.Author{}, err
	}

	return models.Author{}, err
}
