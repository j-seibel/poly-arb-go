package client

import (
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
	volume_to_trade := int64(10e10)
	const max_volume = 5 * PRICE_MULT
	// slice to keep track of indexes of the tokens to trade
	indexes := make([]int, len(markets_to_trade))
	for _, market := range markets_to_trade {
		no_token_id := market.No_token_id
		no_token_price := OrderBooks[neg_risk_id].order_books[no_token_id].min_ask
		no_token_volume := OrderBooks[neg_risk_id].order_books[no_token_id].asks[no_token_price]
		volume_to_trade = min(volume_to_trade, no_token_volume)

	}
	volume_to_trade = min(volume_to_trade, max_volume)
	if volume_to_trade < 5 {
		return
	}
	for _, market := range markets_to_trade {
		no_token_id := market.No_token_id
		no_token_index := TokenToIndex[no_token_id]
		no_token_price := OrderBooks[neg_risk_id].order_books[no_token_id].min_ask
		volume_to_trade = min(volume_to_trade, max_volume)
		indexes = append(indexes, no_token_index)

		if PRICE_MULT-no_token_price < 1 {
			continue
		}

		// Increment wait group counter

		// ExecuteOrder(no_token_price, volume_to_trade, no_token_id)
		// Execute order in a goroutine
		go func(no_token_price int64, volume_to_trade int64, no_token_id string, condition_id string) {
			defer ExecuteOrder(no_token_price, volume_to_trade, no_token_id)
		}(no_token_price, volume_to_trade, no_token_id, market.Condition_id)
		cooldown_queue[no_token_id] = true

		go func(no_token_id string) {
			time.Sleep(time.Duration(30) * time.Second)
			cooldown_queue[no_token_id] = false
		}(neg_risk_id)
	}

	time.Sleep(time.Duration(30) * time.Second)
	// CPU_Profiler.StopCPUProfile()

	// Wait for all orders to complete

	// Exit after all goroutines finish
	os.Exit(0)

}

func RoundToTickSize(price int64, token_id string) int64 {
	tick_size := TokenToTicksize[token_id]
	return price - (price % int64(tick_size*PRICE_MULT))
}
