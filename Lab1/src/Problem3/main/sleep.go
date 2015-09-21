package main

import (
	"time"
	"fmt"
)

func Sleep(duration int) {

	select {
	case <-time.After(time.Millisecond * time.Duration(duration)):
		return
	}
}

func main() {
	var duration int
	fmt.Println("Enter sleep duration:")
	fmt.Scanf("%d", &duration)
	fmt.Println("sleeping...")
	Sleep(duration)
	fmt.Println("out of sleep.. after " , duration , " ms")
}
