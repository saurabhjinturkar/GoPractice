package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main2() {
	fmt.Println("Welcome!")

	for {
		fmt.Println("==========================================")
		fmt.Println("Please enter your choice: ")
		fmt.Println("1. Purchase Shares")
		fmt.Println("2. View Transactions")
		fmt.Println("3. Exit")
		reader := bufio.NewReader(os.Stdin)
		t, _, err := reader.ReadRune()

		choice := int(t) - 48
		if err != nil {
			panic(err)
		}

		switch choice {
		case 1:
			reader.Reset(os.Stdin)
			fmt.Println("Enter budget:")
			t, _, err := reader.ReadLine()
			if err != nil {
				panic(err)
			}
			budget, err := strconv.ParseFloat(string(t), 64)
			if err != nil {
				panic(err)
			}
			fmt.Println(budget)

			fmt.Println("Enter Stock symbols and percentages:")
			fmt.Println("Example: GOOG:100%,YHOO:50%")
			t, _, err = reader.ReadLine()
			if err != nil {
				panic(err)
			}

			fmt.Println("Purchasing...")
			reply := purchase(string(t), budget)
			fmt.Println("TradeId==>", reply.TradeId)
			fmt.Println("Transaction Summary==>", reply.Stocks)
			fmt.Println("Unvested Amount==>", reply.UnvestedAmount)

		case 2:
			fmt.Println("Enter transaction Id: ")
			reader.Reset(os.Stdin)
			t, _, err := reader.ReadLine()
			if err != nil {
				panic(err)
			}
			id, err := strconv.Atoi(string(t))
			if err != nil {
				panic(err)
			}
			reply := getTransacionDetails(id)
			fmt.Println("TradeId==>", reply.TradeId)
			fmt.Println("Transaction Summary==>", reply.Stocks)
			fmt.Println("Current Marketvalue==>", reply.CurrentMarketValue)
			fmt.Println("Unvested Amount==>", reply.UnvestedAmount)
		case 3:
			fmt.Println("Goodbye!")
			os.Exit(0)
		}
	}
}


