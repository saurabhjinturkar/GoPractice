package commons

import (
	"fmt"
	"strconv"
)

type Stock struct {
	Name   string
	Symbol string
	Number int
	Price  float64
}
type Transaction struct {
	Stocks                           []Stock
	TransactionId                    int
	InvestedAmount, UninvestedAmount float64
}

func (t *Transaction) CleanTransaction() {
	fmt.Println(*t)
	stocks := t.Stocks
//	var temp []stock
	for i := len(stocks) - 1; i >= 0; i-- {
		if stocks[i].Number == 0 {
			fmt.Println("Removing... ", stocks[i])
			// remove that stock
			stocks = append(stocks[:i], stocks[i+1:]...)
		}
	}
	t.Stocks = stocks
}

type StockArgs struct {
	StockSymbolAndPercentage string
	Budget                   float64
}

type StockReply struct {
	TradeId        int
	Stocks         string
	UnvestedAmount float64
}

type TransactionArgs struct {
	TradeId int
}

type TransactionReply struct {
	TradeId                            int
	Stocks                             string
	CurrentMarketValue, UnvestedAmount float64
}

func (s *StockReply) FormatResponse(stocks []Stock) string {
	responseString := ""
	for _, stock := range stocks {
		responseString += stock.Symbol + ":" + strconv.Itoa(stock.Number) + ":$" + strconv.FormatFloat(stock.Price*float64(stock.Number), 'f', 2, 64) + ","
	}
	return responseString
}

func FormatResponse(stocks[] Stock, newStockValues[] Stock) string {
	stocksOutput := ""
	for i, stock := range stocks {
		var isUp = ""
		if stock.Price < newStockValues[i].Price {
			isUp = "+"
		} else if stock.Price > newStockValues[i].Price {
			isUp = "-"
		}
		stocksOutput += stock.Symbol + ":" + strconv.Itoa(stock.Number) + ":" + isUp + "$" + strconv.FormatFloat(newStockValues[i].Price, 'f', 2, 64) + ","

	}
	return stocksOutput
}
