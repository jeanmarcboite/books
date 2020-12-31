package calibre

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type Rating struct {
	ID    uint
	Value uint
}

type BookRatingLink struct {
	ID     uint
	Book   uint
	Rating uint
}

func GetRatings(db *CalibreDB, database *sqlx.DB) error {
	db.Ratings = make(map[uint]uint)
	rows, err := database.Queryx("select id, rating from ratings")
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var rating Rating
			err = rows.Scan(&rating.ID, &rating.Value)
			if err != nil {
				return err
			}
			db.Ratings[rating.ID] = rating.Value
		}
		err = GetBooksRatingsLink(db, database)
		if err == nil {
			return rows.Err()
		}
	}

	return err
}

func GetBooksRatingsLink(db *CalibreDB, database *sqlx.DB) error {
	if db.Books == nil {
		return nil
	}
	rows, err := database.Queryx("select * from books_ratings_link")

	if err == nil {
		defer rows.Close()
		var link BookRatingLink
		for rows.Next() {
			err = rows.StructScan(&link)
			if err != nil {
				return err
			}

			rating, ratingOK := db.Ratings[link.Rating]
			book, bookOK := db.Books[link.Book]

			if bookOK && ratingOK {
				book.Rating = rating
			} else {
				if !ratingOK {
					log.Error().Uint("id", link.Rating).Msg("Invalid rating id")

				}
				if !bookOK {
					log.Error().Uint("id", link.Book).Msg("Invalid book id")
				}
			}
		}
	}
	return err
}
