package client

import (
	"fmt"
	"math"
	"strconv"
)

var OrderBooks = make(map[string]*NegRiskOrderBook)

const PRICE_MULT = 1000000

func InitOrderBooks() {
	for key, value := range NegRiskMarketMap {
		OrderBooks[key] = &NegRiskOrderBook{value[0].condition_id, value[0].yes_token_id, value[0].no_token_id, make(map[string]*OrderBook), int64(len(value)), int64(PRICE_MULT * len(value))}
		defaultAskMap := make(map[int64]int64)
		defaultAskMap[1] = 10e10
		for _, market := range value {
			OrderBooks[key].order_books[market.no_token_id] = &OrderBook{defaultAskMap, make(map[int64]int64), PRICE_MULT}

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
		fprice, _ := strconv.ParseFloat(priceChangeEvent.Price, 64)
		price := int64(fprice * PRICE_MULT)
		famount, _ := strconv.ParseFloat(priceChangeEvent.Size, 64)
		amount := int64(famount * PRICE_MULT)
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

		difference := OrderBooks[negRiskId].order_books[tokenId].min_ask - old_min
		OrderBooks[negRiskId].total_price = OrderBooks[negRiskId].total_price + difference
		// fmt.Printf("%v\n", OrderBooks[negRiskId])
		if OrderBooks[negRiskId].total_price < ((OrderBooks[negRiskId].num_assets - 1) * PRICE_MULT) {
			// Arbitrage opportunity
			// fmt.Println("Arbitrage opportunity")
			ExecuteArb(negRiskId)
			// fmt.Println("Price Arbitrage opportunity", OrderBooks[negRiskId].total_price, (OrderBooks[negRiskId].num_assets-1)*PRICE_MULT, price)
			// fmt.Println(OrderBooks[negRiskId].condition_id)

		}

	}
}

func UpdateOrderBook(bookDataEvent BookData) {
	negRiskId := TokenToMarketMap[bookDataEvent.AssetID]
	old_min := OrderBooks[negRiskId].order_books[bookDataEvent.AssetID].min_ask
	newBook := OrderBook{make(map[int64]int64), make(map[int64]int64), PRICE_MULT}
	for _, ask := range bookDataEvent.Asks {
		fprice, _ := strconv.ParseFloat(ask.Price, 64)
		price := int64(fprice * PRICE_MULT)
		famount, _ := strconv.ParseFloat(ask.Size, 64)
		amount := int64(famount * PRICE_MULT)
		newBook.asks[price] = amount

	}
	newBook.asks[PRICE_MULT] = 10e10
	newBook.min_ask = findMinKey(newBook.asks)
	OrderBooks[negRiskId].order_books[bookDataEvent.AssetID] = &newBook

	difference := OrderBooks[negRiskId].order_books[bookDataEvent.AssetID].min_ask - old_min
	OrderBooks[negRiskId].total_price = OrderBooks[negRiskId].total_price + difference
	// fmt.Println("OrderBook Updated", OrderBooks[negRiskId].total_price, difference)
	if OrderBooks[negRiskId].total_price < (OrderBooks[negRiskId].num_assets-1)*PRICE_MULT {
		// Arbitrage opportunity
		// fmt.Println("OrderBOOK opportunity", OrderBooks[negRiskId].total_price, (OrderBooks[negRiskId].num_assets-1)*PRICE_MULT)
		// fmt.Println(OrderBooks[negRiskId].condition_id)

		ExecuteArb(negRiskId)

	}
}

func findMinKey(m map[int64]int64) int64 {
	min := int64(10e10)
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
