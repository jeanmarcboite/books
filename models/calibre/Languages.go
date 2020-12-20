package calibre

import (
	"github.com/jmoiron/sqlx"
)

type Language struct {
	LangCode  string
	ItemOrder uint
}

func (this *CalibreDB) GetLanguages(database *sqlx.DB) error {
	languages, err := this.GetStrings(database, "languages")
	if err == nil {
		return this.GetBooksLanguagesLink(database, languages)
	}

	return err
}

func (this CalibreDB) GetBooksLanguagesLink(database *sqlx.DB, languages map[uint]string) error {
	if this.Books == nil {
		return nil
	}
	rows, err := database.Queryx("select book, lang_code, item_order from books_languages_link")

	if err == nil {
		defer rows.Close()
		var language Language
		var lang_code uint
		for rows.Next() {
			var book uint
			err = rows.Scan(&book, &lang_code, &language.ItemOrder)
			if err != nil {
				return err
			}
			language.LangCode = languages[lang_code]
			this.Books[book].Languages = append(this.Books[book].Languages, language)
		}
	}
	return err
}
