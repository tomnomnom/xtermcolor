package xtermcolors

import "testing"

type intTest struct {
	in  int
	exp int
}

type rgbTest struct {
	r, g, b int
	exp     int
}

var intTests = []intTest{
	{0x000000, 0},
	{0xFF0000, 9},
	{0xCC66FF, 171},
}

var rgbTests = []rgbTest{
	{255, 0, 255, 13},
	{123, 123, 123, 8},
	{63, 254, 128, 84},
}

func TestInts(t *testing.T) {
	for _, test := range intTests {
		act := FromInt(test.in)
		if act != test.exp {
			t.Errorf("FromInt(0x%06.X) should be %d; got %d", test.in, test.exp, act)
		}
	}
}

func TestRGB(t *testing.T) {
	for _, test := range rgbTests {
		act := FromRGB(test.r, test.g, test.b)
		if act != test.exp {
			t.Errorf("FromRGB(%d, %d, %d) should be %d; got %d", test.r, test.g, test.b, test.exp, act)
		}
	}
}
