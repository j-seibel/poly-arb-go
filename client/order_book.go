package client

import (
	"fmt"
	"math"
	"strconv"
)

var OrderBooks = make(map[string]*NegRiskOrderBook)

func InitOrderBooks() {
	for key, value := range NegRiskMarketMap {
		OrderBooks[key] = &NegRiskOrderBook{value[0].condition_id, value[0].yes_token_id, value[0].no_token_id, make(map[string]*OrderBook), len(value), float64(len(value))}
		defaultAskMap := make(map[float64]float64)
		defaultAskMap[1] = 10e10
		for _, market := range value {
			OrderBooks[key].order_books[market.no_token_id] = &OrderBook{defaultAskMap, make(map[float64]float64), 1}

		}

	}

	// fmt.Println(OrderBooks)
	fmt.Println("Order Books Initialized")

}

func UpdateOrderPrice(priceChangeEvent PriceChange) {
	negRiskId := TokenToMarketMap[priceChangeEvent.AssetID]
	tokenId := priceChangeEvent.AssetID
	if priceChangeEvent.Side != "BUY" {
		// fmt.Println("Buy")
		old_min := OrderBooks[negRiskId].order_books[tokenId].min_ask
		price, _ := strconv.ParseFloat(priceChangeEvent.Price, 64)
		amount, _ := strconv.ParseFloat(priceChangeEvent.Size, 64)
		OrderBooks[negRiskId].order_books[tokenId].asks[price] = amount
		if price < OrderBooks[negRiskId].order_books[tokenId].min_ask && amount != 0 {
			orderBook := OrderBooks[negRiskId].order_books[tokenId]
			orderBook.min_ask = price
			OrderBooks[negRiskId].order_books[tokenId] = orderBook
		}

		if amount == 0 {
			delete(OrderBooks[negRiskId].order_books[tokenId].asks, price)
			if price == OrderBooks[negRiskId].order_books[tokenId].min_ask {

				OrderBooks[negRiskId].order_books[tokenId].min_ask = findMinKey(OrderBooks[negRiskId].order_books[tokenId].asks)

			}
		}

		difference := roundto2decimals(OrderBooks[negRiskId].order_books[tokenId].min_ask - old_min)
		OrderBooks[negRiskId].total_price = roundto2decimals(OrderBooks[negRiskId].total_price + difference)
		if OrderBooks[negRiskId].total_price < float64(OrderBooks[negRiskId].num_assets-1) {
			// Arbitrage opportunity
			// fmt.Println("Arbitrage opportunity")
			ExecuteArb(negRiskId)
			fmt.Println("Price Arbitrage opportunity", OrderBooks[negRiskId].total_price, float64(OrderBooks[negRiskId].num_assets-1))
			fmt.Println(OrderBooks[negRiskId].condition_id)

		}

	}
}

func UpdateOrderBook(bookDataEvent BookData) {
	negRiskId := TokenToMarketMap[bookDataEvent.AssetID]
	old_min := OrderBooks[negRiskId].order_books[bookDataEvent.AssetID].min_ask
	newBook := OrderBook{make(map[float64]float64), make(map[float64]float64), 1}
	for _, ask := range bookDataEvent.Asks {
		price, _ := strconv.ParseFloat(ask.Price, 64)
		amount, _ := strconv.ParseFloat(ask.Size, 64)
		newBook.asks[price] = amount

	}
	newBook.asks[1] = 10e10
	newBook.min_ask = findMinKey(newBook.asks)
	OrderBooks[negRiskId].order_books[bookDataEvent.AssetID] = &newBook

	difference := roundto2decimals(OrderBooks[negRiskId].order_books[bookDataEvent.AssetID].min_ask - old_min)
	OrderBooks[negRiskId].total_price = roundto2decimals(OrderBooks[negRiskId].total_price + difference)
	if OrderBooks[negRiskId].total_price < float64(OrderBooks[negRiskId].num_assets-1) {
		// Arbitrage opportunity
		fmt.Println("Order Book Arbitrage opportunity")
		fmt.Println(OrderBooks[negRiskId].condition_id)
		// ExecuteArb(negRiskId)

	}
}

func findMinKey(m map[float64]float64) float64 {
	min := 10e10
	for key := range m {
		if key < min {
			min = key
		}
	}
	return min
}

func roundto2decimals(n float64) float64 {
	return math.Round(n*1000) / 1000
}
