package main

import (
	"os"
	"os/signal"
	"syscall"

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
			log.Info().Msg("Executable is now running. Press CTRL-C to exit.")
			signalChannel := make(chan os.Signal, 1)
			signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
			<-signalChannel
			log.Info().Msg("Good bye!")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal().Stack().Err(err).Msg("")
	}
}
