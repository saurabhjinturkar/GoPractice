package main

import (
	"core/commons"
	"fmt"
	"net/rpc/jsonrpc"
	"log"
)

// This is sample client. Other two functions are used by interactive client and cmd client
func main1() {
	client, err := jsonrpc.Dial("tcp", "localhost:8080")

	if err != nil {
		panic(err)
	}
	defer client.Close()

	args := & commons.StockArgs{"GOOG:100%", 2000.0}
	var reply commons.StockReply

	for i := 0; i < 1; i++ {

		err = client.Call("StockService.PurchaseShares", args, &reply)
		if err != nil {
			log.Fatal("error:", err)
		}
		fmt.Println(reply)
	}
	
	for i := 2; i <= 6; i++ {
		args:= &commons.TransactionArgs {i}
		var reply commons.TransactionReply
		err = client.Call("StockService.GetTransactionDetail", args, &reply)
		if err != nil {
			log.Fatal("error:", err)
		}
		fmt.Println(reply)
	}
}

func purchase(stockString string, budget float64) commons.StockReply {
	client, err := jsonrpc.Dial("tcp", "localhost:8080")

	if err != nil {
		panic(err)
	}
	defer client.Close()

	args := &commons.StockArgs{stockString, budget}
	var reply commons.StockReply

	err = client.Call("StockService.PurchaseShares", args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}
	return reply
}

func getTransacionDetails(transactionId int) commons.TransactionReply {
	client, err := jsonrpc.Dial("tcp", "localhost:8080")

	if err != nil {
		panic(err)
	}
	defer client.Close()

	args := &commons.TransactionArgs{transactionId}
	var reply commons.TransactionReply
	err = client.Call("StockService.GetTransactionDetail", args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}
	return reply
}