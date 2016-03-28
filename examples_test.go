package xtermcolor_test

import (
	"fmt"
	"image/color"

	"github.com/tomnomnom/xtermcolor"
)

func ExampleFromColor() {
	fmt.Println(xtermcolor.FromColor(color.RGBA{120, 210, 120, 255}))

	// Output:
	// 114
}

func ExampleFromInt() {
	fmt.Println(xtermcolor.FromInt(0xCC66FFFF))

	// Output:
	// 171
}

func ExampleFromHexStr() {
	code, err := xtermcolor.FromHexStr("#CC66FF")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(code)

	// Output:
	// 171
}
