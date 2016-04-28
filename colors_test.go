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
	{color.RGBA{127, 127, 127, 255}, 8},
	{color.RGBA{63, 254, 128, 255}, 84},
	{color.RGBA{215, 0, 95, 255}, 161},
	{color.RGBA{204, 102, 255, 255}, 171},
	{color.RGBA{17, 156, 17, 255}, 34},
}

func TestRGBA(t *testing.T) {
	for _, test := range rgbaTests {
		act := FromColor(test.in)
		if act != test.exp {
			t.Errorf("FromRGB(%v) should be %d (%v); got %d (%v)", test.in, test.exp, Colors[test.exp], act, Colors[act])
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
	{0xCC66FFFF, 171},
}

func TestInts(t *testing.T) {
	for _, test := range intTests {
		act := FromInt(test.in)
		if act != test.exp {
			t.Errorf("FromInt(0x%08.X) (%v) should be %d (%v); got %d (%v)", test.in, intToRGBA(test.in), test.exp, Colors[test.exp], act, Colors[act])
		}
	}
}

// HexStr Tests
type hexStrTest struct {
	in  string
	exp uint8
	err error
}

var hexStrTests = []hexStrTest{
	{"#000000", 0, nil},
	{"FF0000", 9, nil},
	{"#CC66FF", 171, nil},
	{"", 0, ErrorEmptyHexStr},
	{"#nothex", 0, ErrorHexParse},
}

func TestHexStr(t *testing.T) {
	for _, test := range hexStrTests {
		act, err := FromHexStr(test.in)

		if err != test.err {
			t.Errorf("FromHexStr error value should have been %s but was %s", test.err, err)
		}
		if act != test.exp {
			t.Errorf("FromHexStr(%s) should be %d (%v); got %d (%v)", test.in, test.exp, Colors[test.exp], act, Colors[act])
		}
	}
}
