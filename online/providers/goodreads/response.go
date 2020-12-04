package goodreads

import (
	"encoding/xml"

	"github.com/jeanmarcboite/books/models"
)

// Response -- GoodreadsResponse
// GoodreadsResponse was generated 2020-12-03 16:27:20 by box on redkeep.
type GoodreadsResponse struct {
	XMLName xml.Name `xml:"GoodreadsResponse"`
	Text    string   `xml:",chardata"`
	Request struct {
		Text           string `xml:",chardata"`
		Authentication string `xml:"authentication"`
		Key            string `xml:"key"`
		Method         string `xml:"method"`
	} `xml:"Request"`
	Book struct {
		Text             string `xml:",chardata"`
		ID               string `xml:"id"`
		Title            string `xml:"title"`
		ISBN             string `xml:"isbn"`
		Isbn13           string `xml:"isbn13"`
		Asin             string `xml:"asin"`
		KindleAsin       string `xml:"kindle_asin"`
		MarketplaceID    string `xml:"marketplace_id"`
		CountryCode      string `xml:"country_code"`
		ImageURL         string `xml:"image_url"`
		SmallImageURL    string `xml:"small_image_url"`
		PublicationYear  string `xml:"publication_year"`
		PublicationMonth string `xml:"publication_month"`
		PublicationDay   string `xml:"publication_day"`
		Publisher        string `xml:"publisher"`
		LanguageCode     string `xml:"language_code"`
		IsEbook          string `xml:"is_ebook"`
		Description      string `xml:"description"`
		Work             struct {
			Text string `xml:",chardata"`
			ID   struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"id"`
			BooksCount struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"books_count"`
			BestBookID struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"best_book_id"`
			ReviewsCount struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"reviews_count"`
			RatingsSum struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"ratings_sum"`
			RatingsCount struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"ratings_count"`
			TextReviewsCount struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"text_reviews_count"`
			OriginalPublicationYear struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"original_publication_year"`
			OriginalPublicationMonth struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"original_publication_month"`
			OriginalPublicationDay struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"original_publication_day"`
			OriginalTitle      string `xml:"original_title"`
			OriginalLanguageID struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
				Nil  string `xml:"nil,attr"`
			} `xml:"original_language_id"`
			MediaType  string `xml:"media_type"`
			RatingDist string `xml:"rating_dist"`
			DescUserID struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"desc_user_id"`
			DefaultChapteringBookID struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
				Nil  string `xml:"nil,attr"`
			} `xml:"default_chaptering_book_id"`
			DefaultDescriptionLanguageCode struct {
				Text string `xml:",chardata"`
				Nil  string `xml:"nil,attr"`
			} `xml:"default_description_language_code"`
			WorkURI string `xml:"work_uri"`
		} `xml:"work"`
		AverageRating      string `xml:"average_rating"`
		NumPages           string `xml:"num_pages"`
		Format             string `xml:"format"`
		EditionInformation string `xml:"edition_information"`
		RatingsCount       string `xml:"ratings_count"`
		TextReviewsCount   string `xml:"text_reviews_count"`
		URL                string `xml:"url"`
		Link               string `xml:"link"`
		Authors            struct {
			Text   string        `xml:",chardata"`
			Author models.Author `xml:"author"`
		} `xml:"authors"`
		ReviewsWidget  string `xml:"reviews_widget"`
		PopularShelves struct {
			// Text  string `xml:",chardata"`
			Shelf []models.Shelf `xml:"shelf"`
		} `xml:"popular_shelves"`
		BookLinks struct {
			Text     string `xml:",chardata"`
			BookLink struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id"`
				Name string `xml:"name"`
				Link string `xml:"link"`
			} `xml:"book_link"`
		} `xml:"book_links"`
		BuyLinks struct {
			Text    string `xml:",chardata"`
			BuyLink []struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id"`
				Name string `xml:"name"`
				Link string `xml:"link"`
			} `xml:"buy_link"`
		} `xml:"buy_links"`
		SeriesWorks  string `xml:"series_works"`
		SimilarBooks struct {
			Text string `xml:",chardata"`
			Book []struct {
				Text               string `xml:",chardata"`
				ID                 string `xml:"id"`
				URI                string `xml:"uri"`
				Title              string `xml:"title"`
				TitleWithoutSeries string `xml:"title_without_series"`
				Link               string `xml:"link"`
				SmallImageURL      string `xml:"small_image_url"`
				ImageURL           string `xml:"image_url"`
				NumPages           string `xml:"num_pages"`
				Work               struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id"`
				} `xml:"work"`
				ISBN             string `xml:"isbn"`
				Isbn13           string `xml:"isbn13"`
				AverageRating    string `xml:"average_rating"`
				RatingsCount     string `xml:"ratings_count"`
				PublicationYear  string `xml:"publication_year"`
				PublicationMonth string `xml:"publication_month"`
				PublicationDay   string `xml:"publication_day"`
				Authors          struct {
					Text   string `xml:",chardata"`
					Author struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id"`
						Name string `xml:"name"`
						Link string `xml:"link"`
					} `xml:"author"`
				} `xml:"authors"`
			} `xml:"book"`
		} `xml:"similar_books"`
	} `xml:"book"`
}
