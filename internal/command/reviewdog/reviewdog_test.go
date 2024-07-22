package reviewdog_test

import (
	"testing"

	"github.com/urfave/cli/v2"

	"github.com/takumin/dogowner/internal/command/reviewdog"
	"github.com/takumin/dogowner/internal/config"
)

func TestNewCommands(t *testing.T) {
	cfg := config.NewConfig()
	flags := []cli.Flag{}
	cmd := reviewdog.NewCommands(cfg, flags)

	if cmd.Name != "reviewdog" {
		t.Errorf("expected command name to be 'reviewdog', but got '%s'", cmd.Name)
	}

	if cmd.Usage != "running reviewdog" {
		t.Errorf("expected command usage to be 'command reviewdog', but got '%s'", cmd.Usage)
	}

	for _, subcmd := range cmd.Subcommands {
		if subcmd == nil {
			t.Errorf("expected subcommand to not be nil")
		}
	}

	if err := cmd.Action(&cli.Context{}); err != nil {
		t.Errorf("expected action func to not be nil")
	}
}
