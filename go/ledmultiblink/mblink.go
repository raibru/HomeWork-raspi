package main

import (
	"os"

	"rbr.de/ledmultiblink/cmd"
)

func main() {
	defer os.Exit(0)
	cmd.Execute()
}
