package main

import (
	"fmt"

	"github.com/tomnomnom/xtermcolor"
)

func main() {
	fmt.Println(xtermcolor.FromInt(0xCC66FF))
	fmt.Println(xtermcolor.FromRGB(120, 210, 120))
}
