package main

import (
	"flag"
	"fmt"
	"image/color"
	"os"
	"strconv"

	"github.com/tomnomnom/xtermcolor"
)

const (
	exitOK = iota
	exitNoArgs
	exitHexParse
	exitRBGParse
)

var usage = `
Usage:
  xtermcolor [<hexString>|<red> <green> <blue>]

Examples:
  xtermcolor CC66FF
  xtermcolor 210 128 0
`

func handleError(code int, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		fmt.Fprintf(os.Stderr, usage)
		os.Exit(code)
	}
}

func getRGBArgs() (color.Color, error) {
	parseErr := func(input string) error {
		return fmt.Errorf("Failed to parse argument [%s], must be in the range 0-255", input)
	}

	r, err := strconv.ParseUint(flag.Arg(0), 10, 8)
	if err != nil {
		return nil, parseErr(flag.Arg(0))
	}
	g, err := strconv.ParseUint(flag.Arg(1), 10, 8)
	if err != nil {
		return nil, parseErr(flag.Arg(1))
	}
	b, err := strconv.ParseUint(flag.Arg(2), 10, 8)
	if err != nil {
		return nil, parseErr(flag.Arg(2))
	}

	return color.RGBA{uint8(r), uint8(g), uint8(b), 255}, nil
}

func main() {
	flag.Parse()

	switch flag.NArg() {
	case 1:
		hexStr := flag.Arg(0)
		code, err := xtermcolor.FromHexStr(hexStr)
		handleError(exitHexParse, err)
		fmt.Println(code)
	case 3:
		rgba, err := getRGBArgs()
		handleError(exitRBGParse, err)
		code := xtermcolor.FromColor(rgba)
		fmt.Println(code)
	default:
		handleError(exitNoArgs, fmt.Errorf("Invalid or no arguments"))
	}

	os.Exit(0)
}
