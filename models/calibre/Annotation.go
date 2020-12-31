package calibre

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type AnnotationRow struct {
	ID              uint
	book            uint
	format          string
	user_type       string `db:"user_type"`
	user            string
	timestamp       float64
	annot_id        string `db:"annot_id"`
	annot_type      string `db:"annot_type"`
	annot_data      string `db:"annot_data"`
	searchable_text string `db:"searchable_text"`
}

type Annotation struct {
	Format         string
	UserType       string `db:"user_type"`
	User           string
	Timestamp      float64
	AnnotID        string `db:"annot_id"`
	AnnotType      string `db:"annot_type"`
	AnnotData      string `db:"annot_data"`
	SearchableText string `db:"searchable_text"`
}

func GetAnnotations(db *CalibreDB, database *sqlx.DB) error {
	rows, err := database.Queryx("select * from annotations")

	if err != nil {
		return err
	}

	defer rows.Close()
	for rows.Next() {
		var annotationRow AnnotationRow
		err = rows.StructScan(&annotationRow)
		if err != nil {
			return err
		}
		if book, ok := db.Books[annotationRow.book]; ok {
			book.Annotations = append(book.Annotations,
				Annotation{
					Format:         annotationRow.format,
					UserType:       annotationRow.user_type,
					User:           annotationRow.user,
					Timestamp:      annotationRow.timestamp,
					AnnotID:        annotationRow.annot_id,
					AnnotType:      annotationRow.annot_type,
					AnnotData:      annotationRow.annot_data,
					SearchableText: annotationRow.searchable_text,
				})
		} else {
			log.Error().Uint("id", annotationRow.book).Msg("Invalid book id")
		}
	}

	return rows.Err()
}
