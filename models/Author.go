package models

// Author -- Author name and id
type Author0 struct {
	Name string
	Key  string
	ID   string
	Link string
}
type Author struct {
	Key      string
	Text     string `xml:",chardata"`
	ID       string `xml:"id"`
	Name     string `xml:"name"`
	Role     string `xml:"role"`
	ImageURL struct {
		Text    string `xml:",chardata"`
		Nophoto string `xml:"nophoto,attr"`
	} `xml:"image_url"`
	SmallImageURL struct {
		Text    string `xml:",chardata"`
		Nophoto string `xml:"nophoto,attr"`
	} `xml:"small_image_url"`
	Link             string `xml:"link"`
	AverageRating    string `xml:"average_rating"`
	RatingsCount     string `xml:"ratings_count"`
	TextReviewsCount string `xml:"text_reviews_count"`
}
