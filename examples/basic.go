package main

import (
	"fmt"
	"image/color"

	"github.com/tomnomnom/xtermcolor"
)

func main() {
	fmt.Println(xtermcolor.Colors.Convert(color.RGBA{128, 64, 32, 255}))
	fmt.Println(xtermcolor.FromColor(color.RGBA{120, 210, 120, 255}))
	fmt.Println(xtermcolor.FromInt(0xCC66FFFF))

	code, err := xtermcolor.FromHexStr("#FEFEFE")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(code)
}
