package config_test

import (
	"reflect"
	"testing"

	"github.com/takumin/dogowner/internal/config"
)

func TestNewConfig(t *testing.T) {
	if !reflect.DeepEqual(config.NewConfig(), &config.Config{Reviewdog: &config.Reviewdog{}}) {
		t.Error("expected config struct to be equal, but got not equal")
	}

	if !reflect.DeepEqual(config.NewConfig(config.LogLevel("TEST")), &config.Config{LogLevel: "TEST", Reviewdog: &config.Reviewdog{}}) {
		t.Error("expected config struct to be equal, but got not equal")
	}
}
