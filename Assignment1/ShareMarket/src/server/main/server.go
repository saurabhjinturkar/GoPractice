package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"regexp"
	"strconv"
	"strings"
	"core/commons"
)

var transcationId int = 0
var validator = regexp.MustCompile(`^([A-Z]{4}:[0-9]{1,3}%)(,[A-Z]{4}:[0-9]{1,3}%)*$`)

func getNewTransactionId() int {
	transcationId += 1
	return transcationId
}

// Function to validate stock request string
func ValidateQuery(query string) bool {
	return validator.MatchString(query)
}

var transactions map[int]commons.Transaction

func getStockSymbols(stocks []commons.Stock) []string {
	symbols := []string{}
	for _, stock := range stocks {
		symbols = append(symbols, stock.Symbol)
	}
	return symbols
}

func fetchStockValues(symbols []string) (stocks []commons.Stock) {

	// Append stock symbols to the url
	url := "http://finance.yahoo.com/d/quotes.csv?s="
	urls := []string{url}
	for _, symbol := range symbols {
		urls = append(urls, symbol)
	}
	urls = append(urls, "&f=nsb")
	url = strings.Join(urls, "+")

	log.Println("Calling Yahoo! API with URL: ", url)

	// Make HTTP GET Request to YAHOO API URL
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Read CSV data
	reader := csv.NewReader(resp.Body)
	records, parseErr := reader.ReadAll()
	if parseErr != nil {
		log.Fatal(err)
	}

	// Create an array of stocks with stock name, symbol and price.
	// Number of stocks is initialized to 0
	for _, record := range records {
		f, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			panic(record)
		}
		stocks = append(stocks, commons.Stock{record[0], record[1], 0, f})
	}

	log.Println("Fetched stock values are: ", stocks)
	return stocks
}

func createPurchaseRequests(request string) (stocks []commons.Stock, percentages []float64, err error) {
	// Split the request string into stock symbols and percentages
	requests := strings.Split(request, ",")
	symbols := []string{}
	percentages = []float64{}
	for _, request := range requests {
		values := strings.Split(request, ":")
		symbols = append(symbols, values[0])
		percentString := strings.Replace(values[1], "%", "", -1)
		percentage, err := strconv.ParseFloat(percentString, 64)
		if err != nil {
			return nil, nil, errors.New("Percentage can not be parsed. Parsable example: GOOG:50%,YHOO:50%")
		}
		percentages = append(percentages, percentage)
	}

	// Check if percentage sum is equal to 100, else return error
	isValid := CheckPercentages(percentages)
	if !isValid {
		return nil, nil, errors.New("Sum of the percentage values should be 100.")
	}

	// Fetch stock values from Yahoo API for symbols passed
	stocks = fetchStockValues(symbols)
	return stocks, percentages, nil
}

func buyStocks(stocks []commons.Stock, percentages []float64, balance float64) int {

	// Create an empty transaction
	transaction := commons.Transaction{stocks, getNewTransactionId(), 0, 0}

	// Iterate over stocks and find number of shares that can be purchased in given budget
	for i := 0; i < len(stocks); i++ {
		allottedBalance := percentages[i] * balance / 100
		stocks[i].Number = int(allottedBalance / stocks[i].Price)
		transaction.InvestedAmount += float64(stocks[i].Number) * stocks[i].Price
		transaction.UninvestedAmount += (allottedBalance - float64(stocks[i].Number)*stocks[i].Price)
	}

	// Remove stocks which are not purchased
	transaction.CleanTransaction()

	// Make an entry in transactions map
	transactions[transaction.TransactionId] = transaction
	log.Println("Completed transaction: ", transaction)
	log.Println("Current state of transactions : ", transactions)
	return transaction.TransactionId
}

// Function to check if sum of the percentages is 100
// returns true if it is 100; otherwise false
func CheckPercentages(percentages []float64) (isValid bool) {
	sum := 0.0
	for _, percentage := range percentages {
		sum += percentage
	}
	if sum != 100 {
		isValid = false
	} else {
		isValid = true
	}
	return isValid
}

func getTransactionDetails(transactionId int) commons.TransactionReply {
	transaction := transactions[transactionId]
	stocks := transaction.Stocks

	// Fetch current stock values
	newStockValues := fetchStockValues(getStockSymbols(stocks))

	// Format response string
	responseString := commons.FormatResponse(stocks, newStockValues)

	// Find current market value of the stocks
	currentMarketValue := 0.00
	for i, stock := range newStockValues {
		currentMarketValue += (stock.Price * float64(stocks[i].Number))
	}
	reply := commons.TransactionReply{transaction.TransactionId, responseString, currentMarketValue, transaction.UninvestedAmount}
	return reply
}

// Initialize transactions map
func init() {
	transactions = make(map[int]commons.Transaction)
}

func main() {
	stockService := new(StockService)
	rpc.Register(stockService)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8080")
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	fmt.Println("Listening on port 8080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
