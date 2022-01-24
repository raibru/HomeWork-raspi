package main

import (
	"os"

	"rbr.de/ledrgb/cmd"
)

func main() {
	defer os.Exit(0)
	cmd.Execute()
}
