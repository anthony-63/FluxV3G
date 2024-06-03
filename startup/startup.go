package startup

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func initLogger() {
	log_file, err := os.Create("data/log.txt")

	writer := io.MultiWriter(os.Stdout)

	if err == nil {
		writer = io.MultiWriter(os.Stdout, log_file)
	}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: writer})
}

func initFilesystem() {
	if _, err := os.ReadDir("data/"); os.IsNotExist(err) {
		os.Mkdir("data/", os.ModeDir)
	}
}

func Startup() {
	initFilesystem()
	initLogger()

	log.Info().Msg("Starting up")
}
