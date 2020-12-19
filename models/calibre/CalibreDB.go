package calibre

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type CalibreDB struct {
	Books map[uint](*Book)
}

func ReadDB(filename string, debug bool) (CalibreDB, error) {
	db := CalibreDB{}
	var err error = nil

	sqlDB, err := sql.Open("sqlite3", filename)

	if err == nil {
		database := sqlx.NewDb(sqlDB, "sqlite3")
		defer database.Close()
		err = database.Ping()

		if err == nil {
			return ReadBooks(database)
		}
	}

	return db, err
}

/*
for row, book := range books {
				booksk[book.ID] = row
			}
			var comments []calibre.Comment
			if result := db.Find(&comments); result.Error != nil {
				fmt.Printf("%+v\n", result)
			} else {
				for _, comment := range comments {
					books[booksk[comment.Book]].Comment = comment.Text
				}
			}

			var authors []calibre.Author
			if result := db.Find(&authors); result.Error != nil {
				fmt.Printf("%+v\n", result)
			} else {
				var links []calibre.BookAuthorLink
				if result := db.Find(&links); result.Error != nil {
					fmt.Printf("%+v\n", result)
				} else {
					for _, link := range links {
						books[booksk[link.Book]].Authors =
							append(books[booksk[link.Book]].Authors,
								link.Author)
					}
				}
			}

			fmt.Println("{}", books)

			for _, v := range booksk {
				fmt.Printf("%+v\n", books[v])
			}
*/
