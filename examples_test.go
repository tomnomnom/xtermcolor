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
	// 177
}
