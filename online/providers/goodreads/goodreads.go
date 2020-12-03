package goodreads

import (
	"encoding/xml"
)

func Name() string {
	return "goodreads"
}

// GoodreadsResponse was generated 2020-12-03 15:11:23 by box on redkeep.
type GoodreadsResponseShowAuthor struct {
	XMLName xml.Name `xml:"GoodreadsResponse"`
	Text    string   `xml:",chardata"`
	Request struct {
		Text           string `xml:",chardata"`
		Authentication string `xml:"authentication"`
		Key            string `xml:"key"`
		Method         string `xml:"method"`
	} `xml:"Request"`
	Author struct {
		Text string `xml:",chardata"`
		ID   string `xml:"id,attr"`
		Name string `xml:"name"`
		Link string `xml:"link"`
	} `xml:"author"`
}
