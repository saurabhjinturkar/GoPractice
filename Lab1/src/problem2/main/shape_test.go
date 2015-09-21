package main

import (
	"testing"
)

type testcase struct {
	shape     Shape
	area      float64
	perimeter float64
}

var tests = []testcase{
	{Circle{7}, 153.93804002589985, 43.982297150257104},
	{Rectangle{10, 20}, 200, 60},
	{Circle{0}, 0, 0},
	{Rectangle{0, 0}, 0, 0},
}

func TestArea(t *testing.T) {

	for _, test := range tests {
		var area float64
		area = test.shape.Area()
		if area != test.area {
			t.Error("Problem in calculating area. Expected area is ", test.area, " Calculated area is ", area)
		}
	}

}

func TestPerimeter(t *testing.T) {

	for _, test := range tests {
		var perimeter float64
		perimeter = test.shape.Perimeter()
		if perimeter != test.perimeter {
			t.Error("Problem in calculating perimeter. Expected perimeter is ", test.perimeter, " Calculated perimeter is ", perimeter)
		}
	}

}
