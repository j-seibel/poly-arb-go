package client

import (
	"encoding/json"
	"fmt"
	"sync"
	// "fmt"
)

var NegRiskMarketMap = make(map[string][]NegRiskMarketInfo)
var TokenToMarketMap = make(map[string]string)
var TokenToTicksize = make(map[string]float64)
var TokenToIndex = make(map[string]int)
var NumNegRiskMarketsMap = make(map[string]int)
var AssetsToWatch []string
var WG sync.WaitGroup

func FindNegRiskMarkets() {
	cursor := "MTAw"
	for cursor != "LTE=" {

		var marketMap map[string]interface{}

		markets := GetMarkets(cursor)
		json.Unmarshal([]byte(markets), &marketMap)

		for _, market := range marketMap["data"].([]interface{}) {
			marketDict := market.(map[string]interface{})
			yes_token := marketDict["tokens"].([]interface{})[0].(map[string]interface{})["token_id"].(string)
			no_token := marketDict["tokens"].([]interface{})[1].(map[string]interface{})["token_id"].(string)
			if marketDict["neg_risk"] == true {
				NumNegRiskMarketsMap[marketDict["neg_risk_market_id"].(string)] += 1
				TokenToIndex[yes_token] = NumNegRiskMarketsMap[marketDict["neg_risk_market_id"].(string)]
				TokenToIndex[no_token] = NumNegRiskMarketsMap[marketDict["neg_risk_market_id"].(string)]

			}

			if marketDict["neg_risk"] == true && marketDict["accepting_orders"] == true && marketDict["enable_order_book"] == true {
				marketInfo := NegRiskMarketInfo{marketDict["condition_id"].(string), yes_token, no_token}
				NegRiskMarketMap[marketDict["neg_risk_market_id"].(string)] = append(NegRiskMarketMap[marketDict["neg_risk_market_id"].(string)], marketInfo)
				TokenToMarketMap[yes_token] = marketDict["neg_risk_market_id"].(string)
				TokenToMarketMap[no_token] = marketDict["neg_risk_market_id"].(string)
				TokenToTicksize[yes_token] = marketDict["minimum_tick_size"].(float64)
				TokenToTicksize[no_token] = marketDict["minimum_tick_size"].(float64)
			}
		}
		cursor = marketMap["next_cursor"].(string)

	}

	for _, value := range NegRiskMarketMap {
		for _, asset := range value {
			AssetsToWatch = append(AssetsToWatch, asset.no_token_id)
		}
	}
	fmt.Println("Assets to watch: ", len(AssetsToWatch))
	InitOrderBooks()
	StartSubscription()
	WG.Wait()
}

func StartSubscription() {
	for i := 0; i < len(AssetsToWatch); i += 400 {
		WG.Add(1)
		go ConnectSocket(AssetsToWatch[i:min(i+400, len(AssetsToWatch))])
	}
}
