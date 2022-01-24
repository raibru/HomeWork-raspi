package cmd

import (
	"fmt"
	"io"
)

// PrintVersion prints the tool versions string
func PrintVersion(w io.Writer) {
	major := "0"
	minor := "1"
	patch := "0"
	bdate := "2022-01-24-20-24-06"

	fmt.Fprintf(w, "rgbled - raspberry pi led in RGB - v%s.%sp%s (%s)\n", major, minor, patch, bdate)
	fmt.Fprintf(w, "  author : rbr <raibru@web.de>\n")
	fmt.Fprintf(w, "  license: MIT\n\n")
}
