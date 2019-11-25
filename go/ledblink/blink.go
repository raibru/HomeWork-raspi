package main

import (
	"os"

	"rbr.de/ledblink/cmd"
)

func main() {
	defer os.Exit(0)
	cmd.Execute()
}
