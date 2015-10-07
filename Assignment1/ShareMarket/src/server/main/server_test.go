package main

import (
	"testing"
)

type testcase struct {
	percentages    []float64
	expectedResult bool
}

var tests = []testcase{
	{[]float64{50, 50}, true},
	{[]float64{0, 100.0}, true},
	{[]float64{100.0, 1.0}, false},
	{[]float64{50.1, 50.1}, false},
}

type validatorTestCase struct {
	query          string
	expectedResult bool
}

var validatorTests = []validatorTestCase{
	{"GOOG:50%", true},
	{"YHOO:50%,MSFT:50%", true},
	{"YHOO:50%,MSFT:50%,AAPL:25%", true},
	{"YHOO50%,MSFT:50%", false},
	{"YHOO:50,MSFT:50%", false},
	{"YHOO:50%MSFT:50%", false},
	{"YHOO:50%,MFT:50%", false},
	{"", false},
}

func TestCheckPercentages(t *testing.T) {
	for _, test := range tests {
		var result bool
		result = CheckPercentages(test.percentages)
		if result != test.expectedResult {
			t.Error("Problem in calculating check value. Expected check value is ", test.expectedResult, " Calculated check value is ", result)
		}
	}
}

func TestValidateQuery(t *testing.T) {
	for i, test := range validatorTests {
		var result bool
		result = ValidateQuery(test.query)
		if result != test.expectedResult {
			t.Error("For test #", i+1, ":", test, "Problem in validator. Expected value is ", test.expectedResult, " Calculated value is ", result)
		}
	}
}
