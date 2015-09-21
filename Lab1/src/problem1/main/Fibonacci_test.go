package main

import (
	"testing"
)

type testpair struct {
	value  int
	output int64
}

var tests = []testpair{
	{0, 0},
	{1, 1},
	{2, 1},
	{10, 55},
}

func TestFibonacci(t *testing.T) {
	for _, pair := range tests {
		o := Fibonacci(pair.value)
		if o != pair.output {
			t.Error("For", pair.value, " expected is ", pair.output, " got ", o)
		}
	}
}
