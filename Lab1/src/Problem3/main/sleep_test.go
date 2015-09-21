package main

import (
	"testing"
	"time"
)

type testpair struct {
	sleepTime    int
	expectedTime int
}

var testcases = []testpair{
	{500, 501},
	{0, 1},
}

func TestSleep(t *testing.T) {

	for _, testcase := range testcases {
		starttime := time.Now()
		Sleep(testcase.sleepTime)
		elapsedtime := time.Since(starttime)
		t.Log("Elapsed time", elapsedtime)
		if elapsedtime > (time.Duration(testcase.expectedTime) * time.Millisecond )	|| elapsedtime < (time.Duration(testcase.sleepTime) * time.Millisecond) {
			t.Error("Sleep took ", elapsedtime, "to execute. Expected time:" , testcase.expectedTime)
		}
	}
}
