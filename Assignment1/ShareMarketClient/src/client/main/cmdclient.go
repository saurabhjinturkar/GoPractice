package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)


//Sample input: <PROGRAM NAME> GOOG:50%,MSFT:50% 2000
// Sample input:<PROGRAM NAME> 3
func main() {
	
	if len(os.Args) == 3 {
		budget, err := strconv.ParseFloat(os.Args[2], 64)
		if err != nil {
			panic(err)
		}
		reply := purchase(os.Args[1], budget)
		b, err := json.Marshal(reply)
		if err != nil {
			panic(err)
			return
		}
		fmt.Println(string(b))
	} else if len(os.Args) == 2 {
		id, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
		reply := getTransacionDetails(id)
		b, err := json.Marshal(reply)
		if err != nil {
			panic(err)
			return
		}
		fmt.Println(string(b))
	} else {
		fmt.Println("Incorrect number of arguments")
	}
}
