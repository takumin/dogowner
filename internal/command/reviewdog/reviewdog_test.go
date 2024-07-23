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
	t.Parallel()

	cases := map[string]struct {
		stdout string
		stderr string
		stdin  string
		args   string
		error  bool
	}{
		"empty":               {"", "", "", "a reviewdog", false},
		"config-path-exists":  {"", "", "", "a reviewdog -c /dev/null", false},
		"config-path-unknown": {"", "", "", "a reviewdog -c /noneexists", true},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			stdout := bytes.NewBuffer(make([]byte, bytes.MinRead))
			stderr := bytes.NewBuffer(make([]byte, bytes.MinRead))
			stdin := strings.NewReader(tt.stdin)
			args := strings.Split(tt.args, " ")

			app := cli.App{
				Commands:  []*cli.Command{reviewdog.NewCommands(config.NewConfig(), []cli.Flag{})},
				Reader:    stdin,
				Writer:    stdout,
				ErrWriter: stderr,
			}

			err := app.Run(args)

			switch {
			case tt.error == true && err == nil:
				t.Errorf("expected running to not be nil")
			case tt.error == false && err != nil:
				t.Error("unexpected error:", err)
			}
		})
	}
}
