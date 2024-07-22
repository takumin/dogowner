package reviewdog_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/urfave/cli/v2"

	"github.com/takumin/dogowner/internal/command/reviewdog"
	"github.com/takumin/dogowner/internal/config"
)

func TestNewCommands(t *testing.T) {
	args := []string{"a", "reviewdog"}
	stdin := strings.NewReader("")
	stdout := bytes.NewBuffer(make([]byte, bytes.MinRead))
	stderr := bytes.NewBuffer(make([]byte, bytes.MinRead))
	app := cli.App{
		Commands:  []*cli.Command{reviewdog.NewCommands(config.NewConfig(), []cli.Flag{})},
		Reader:    stdin,
		Writer:    stdout,
		ErrWriter: stderr,
	}
	if err := app.Run(args); err != nil {
		t.Errorf("expected running to not be nil")
	}
}
