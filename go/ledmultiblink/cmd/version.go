package cmd

import (
	"fmt"
	"io"
)

// PrintMblinkVersion prints the tool versions string
func PrintMblinkVersion(w io.Writer) {

	major := "1"
	minor := "0"
	patch := "15"
	bdate := "2019-11-25-20-27-47"

	fmt.Fprintf(w, "mblink - raspberry pi multi led blinker - v%s.%sp%s (%s)\n", major, minor, patch, bdate)
	fmt.Fprintf(w, "  author : rbr <raibru@web.de>\n")
	fmt.Fprintf(w, "  license: MIT\n\n")
}
