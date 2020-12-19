package calibre

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	books []Book
}

func ReadDB(filename string, debug bool) (DB, error) {
	db := DB{}
	s, err := sql.Open("sqlite3", filename)
	if err != nil {
		log.Fatal(err)
	}
	database := sqlx.NewDb(s, "sqlite3")
	defer database.Close()
	err = database.Ping()
	if err != nil {
		log.Fatal(err)
	}
	rows, err := database.Queryx("select * from books")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		book := Book{}
		err = rows.StructScan(&book)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(book.ID, book.Title)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
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