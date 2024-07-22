package reviewdog

import (
	"github.com/urfave/cli/v2"

	"github.com/takumin/dogowner/internal/config"
)

func NewCommands(cfg *config.Config, flags []cli.Flag) *cli.Command {
	flags = append(flags, []cli.Flag{
		&cli.StringFlag{
			Name:        "config-path",
			Aliases:     []string{"c"},
			Usage:       "reviewdog config path",
			EnvVars:     []string{"CONFIG_PATH"},
			Value:       cfg.ConfigPath,
			Destination: &cfg.ConfigPath,
		},
	}...)
	return &cli.Command{
		Name:    "reviewdog",
		Aliases: []string{"r", "rd"},
		Usage:   "running reviewdog",
		Flags:   flags,
		Action:  action(cfg),
	}
}

func action(cfg *config.Config) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		return nil
	}
}
