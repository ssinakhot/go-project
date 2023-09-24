package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/ssinakhot/goproject"
	"github.com/urfave/cli/v2"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	err := godotenv.Load()
	if err != nil {
		log.Warn().Msg("failed to load .env file")
	}
	app := &cli.App{
		Name:  "Executable",
		Flags: goproject.CliFlags,
		Action: func(*cli.Context) error {
			log.Info().Msg("Hello World")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal().Stack().Err(err).Msg("")
	}
}
