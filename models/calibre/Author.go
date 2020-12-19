package calibre

type Author struct {
	ID   uint
	Name string
	Sort string
	Link string
}

type BookAuthorLink struct {
	ID     uint
	Book   uint
	Author uint
}
