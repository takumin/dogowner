package reviewdog

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/takumin/dogowner/internal/config"
)

func NewCommands(cfg *config.Config, flags []cli.Flag) *cli.Command {
	flags = append(flags, []cli.Flag{
		&cli.StringFlag{
			Name:        "config-path",
			Aliases:     []string{"c"},
			Usage:       "reviewdog config path",
			EnvVars:     []string{"REVIEWDOG_CONFIG_PATH"},
			Value:       cfg.Reviewdog.ConfigPath,
			Destination: &cfg.Reviewdog.ConfigPath,
			Action: func(ctx *cli.Context, s string) error {
				_, err := os.Stat(s)
				if os.IsNotExist(err) {
					return fmt.Errorf("not found reviewdog config file: %s", s)
				}
				cfg.Reviewdog.ConfigPath = s
				return nil
			},
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
		slog.DebugContext(ctx.Context, "reviewdog",
			slog.String("config-path", cfg.Reviewdog.ConfigPath),
			slog.Any("args", ctx.Args().Slice()),
		)
		return nil
	}
}
