package cmd

import (
	"fmt"
	"io"
)

// PrintVersion prints the tool versions string
func PrintVersion(w io.Writer) {
	major := "1"
	minor := "0"
	patch := "8"
	bdate := "2019-11-25-20-57-15"

	fmt.Fprintf(w, "blink - raspberry pi led blinker - v%s.%sp%s (%s)\n", major, minor, patch, bdate)
	fmt.Fprintf(w, "  author : rbr <raibru@web.de>\n")
	fmt.Fprintf(w, "  license: MIT\n\n")
}
