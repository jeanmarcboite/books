package books

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	SetLogLevel(zerolog.InfoLevel)
}

func SetLogLevel(level zerolog.Level) {
	zerolog.SetGlobalLevel(level)
}
