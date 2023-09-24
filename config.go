package goproject

import "github.com/urfave/cli/v2"

type ConfigFlags struct {
	config string
}

var Config = ConfigFlags{}

var CliFlags = []cli.Flag{
	&cli.StringFlag{
		Name:        "config",
		Aliases:     []string{"c"},
		EnvVars:     []string{"CONFIG"},
		Destination: &Config.config,
		Usage:       "Load configuration from `FILE`",
		Required:    false,
	},
}
