package main

import (
	"core/commons"
	"log"
	"errors"
)

type StockService struct{}

func (s *StockService) PurchaseShares(args *commons.StockArgs, reply *commons.StockReply) error {
	log.Println("Processing buy request for:", *args)

	if ValidateQuery(args.StockSymbolAndPercentage) == false {
		return errors.New("Stock symbol and percentage string is not in valid format. You are passing: " + args.StockSymbolAndPercentage + " Example of valid format: GOOG:50%,YHOO:50%")
	}

	stocks, percentages, err := createPurchaseRequests(args.StockSymbolAndPercentage)
	if err != nil {
		return err
	}
	transId := buyStocks(stocks, percentages, args.Budget)
	reply.TradeId = transId
	reply.Stocks = reply.FormatResponse(transactions[transId].Stocks)
	reply.UnvestedAmount = transactions[transId].UninvestedAmount
	return nil
}

func (s *StockService) GetTransactionDetail(args *commons.TransactionArgs, reply *commons.TransactionReply) error {
	log.Println("Processing transaction details request for:", *args)
	*reply = getTransactionDetails(args.TradeId)
	return nil
}
