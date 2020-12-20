package calibre

type Annotation struct {
	ID             uint
	Book           *Book
	Format         string
	UserType       string `db:"user_type"`
	User           string
	Timestamp      float64
	AnnotID        string `db:"annot_id"`
	AnnotType      string `db:"annot_type"`
	AnnotData      string `db:"annot_data"`
	SearchableText string `db:"searchable_text"`
}
