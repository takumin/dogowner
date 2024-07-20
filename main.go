package main

import (
	"os"

	"github.com/takumin/dogowner/internal/command"
)

func main() {
	os.Exit(command.Main(
		os.Stdout,
		os.Stderr,
		os.Stdin,
		os.Args,
	))
}
