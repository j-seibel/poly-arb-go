package client

import (
	"math"
	"os"
	"sync"
	"time"
)

var cooldown_queue = map[string]bool{}
var WG_Order sync.WaitGroup

func ExecuteArb(neg_risk_id string) {

	if cooldown_queue[neg_risk_id] {
		return
	}
	markets_to_trade := NegRiskMarketMap[neg_risk_id]
	volume_to_trade := 10e10
	const max_volume = 50
	// slice to keep track of indexes of the tokens to trade
	indexes := make([]int, len(markets_to_trade))
	for _, market := range markets_to_trade {
		no_token_id := market.no_token_id
		no_token_price := OrderBooks[neg_risk_id].order_books[no_token_id].min_ask
		no_token_volume := OrderBooks[neg_risk_id].order_books[no_token_id].asks[no_token_price]
		no_token_volume = roundto2decimals(no_token_volume)
		volume_to_trade = min(volume_to_trade, no_token_volume)

	}
	volume_to_trade = min(volume_to_trade, max_volume)
	if volume_to_trade < 5 {
		return
	}
	for _, market := range markets_to_trade {
		no_token_id := market.no_token_id
		no_token_index := TokenToIndex[no_token_id]
		no_token_price := OrderBooks[neg_risk_id].order_books[no_token_id].min_ask
		no_token_price = RoundToTickSize(no_token_price, no_token_id)
		volume_to_trade = min(volume_to_trade, max_volume)
		indexes = append(indexes, no_token_index)

		if 1-no_token_price < 0.001 {
			continue
		}

		// Increment wait group counter

		ExecuteOrder(no_token_price, volume_to_trade, no_token_id)
		// Execute order in a goroutine
		// go func(no_token_price float64, volume_to_trade float64, no_token_id string, condition_id string) {
		// 	defer WG_Order.Done() // Mark as done when goroutine finishes
		// 	ExecuteOrder(no_token_price, volume_to_trade, no_token_id)
		// }(no_token_price, volume_to_trade, no_token_id, market.condition_id)
		cooldown_queue[no_token_id] = true
		go func(no_token_id string) {
			time.Sleep(time.Duration(30) * time.Second)
			cooldown_queue[no_token_id] = false
		}(neg_risk_id)
	}

	// Wait for all orders to complete

	// Exit after all goroutines finish
	os.Exit(0)

}

func RoundToTickSize(price float64, token_id string) float64 {
	tick_size := TokenToTicksize[token_id]
	return math.Round(price/tick_size) / (1 / tick_size)
}
