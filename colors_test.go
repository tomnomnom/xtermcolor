package xtermcolor

import (
	"image/color"
	"testing"
)

// RGBA Tests
type rgbaTest struct {
	in  color.RGBA
	exp uint8
}

var rgbaTests = []rgbaTest{
	{color.RGBA{255, 0, 255, 255}, 13},
	{color.RGBA{120, 120, 120, 255}, 8},
	{color.RGBA{63, 254, 128, 255}, 84},
	{color.RGBA{215, 0, 95, 255}, 161},
	{color.RGBA{204, 102, 255, 255}, 177},
	{color.RGBA{17, 156, 17, 255}, 71},
}

func TestRGBA(t *testing.T) {
	for _, test := range rgbaTests {
		act := FromColor(test.in)
		if act != test.exp {
			t.Errorf("FromRGB(%v) should be %d (%v); got %d (%v)", test.in, test.exp, colors[test.exp], act, colors[act])
		}
	}
}

// Int Tests
type intTest struct {
	in  uint32
	exp uint8
}

var intTests = []intTest{
	{0x000000FF, 0},
	{0xFF0000FF, 9},
	{0xCC66FFFF, 177},
}

func TestInts(t *testing.T) {
	for _, test := range intTests {
		act := FromInt(test.in)
		if act != test.exp {

			//t.Errorf("FromInt(0x%06.X) should be %d; got %d", test.in, test.exp, act)
			t.Errorf("FromInt(0x%08.X) (%v) should be %d (%v); got %d (%v)", test.in, intToRGBA(test.in), test.exp, colors[test.exp], act, colors[act])
		}
	}
}
