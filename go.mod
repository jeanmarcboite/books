module github.com/jeanmarcboite/books

go 1.15

replace github.com/jeanmarcboite/epub => ../epub

require (
	github.com/adrg/xdg v0.2.3
	github.com/basgys/goxml2json v1.1.0
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/gobuffalo/packr/v2 v2.8.1
	github.com/jeanmarcboite/epub v0.0.0-00010101000000-000000000000
	github.com/jeanmarcboite/librarytruc v0.0.0-20201129183928-0b5142a38fe3
	github.com/knadh/koanf v0.14.0
	github.com/rs/zerolog v1.20.0
)
