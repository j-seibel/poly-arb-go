package client

import (
	"fmt"
	"os"
	"sync"

	"github.com/goccy/go-json"
	// "fmt"
)

var NegRiskMarketMap = make(map[string][]NegRiskMarketInfo)
var TokenToMarketMap = make(map[string]string)
var TokenToTicksize = make(map[string]float64)
var TokenToIndex = make(map[string]int)
var NumNegRiskMarketsMap = make(map[string]int)
var AssetsToWatch []string
var WG sync.WaitGroup
var CPU_Profiler = &Profiler{}

func FindNegRiskMarkets() {

	// CPU_Profiler.StartCPUProfile()
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
			AssetsToWatch = append(AssetsToWatch, asset.No_token_id)
		}
	}
	writeJsonAssetsToFile(AssetsToWatch, "assets.json")
	writeJsonNegRiskMarketsToFile("neg_risk_markets.json")
	writeJsonAssetsToFile(TokenToMarketMap, "token_to_market.json")
	writeJsonAssetsToFile(TokenToTicksize, "token_to_ticksize.json")
	writeJsonAssetsToFile(TokenToIndex, "token_to_index.json")
	writeJsonAssetsToFile(NumNegRiskMarketsMap, "num_neg_risk_markets.json")

	// StartSubscription()

}

func writeJsonNegRiskMarketsToFile(filename string) {
	file, _ := json.Marshal(NegRiskMarketMap)
	os.WriteFile(filename, file, 0644)
}

func writeJsonAssetsToFile(data interface{}, filename string) {
	file, _ := json.Marshal(data)
	os.WriteFile(filename, file, 0644)
}

func readJsonAssetsFromFile(filename string) []string {
	ast := make([]string, 0)
	file, _ := os.ReadFile(filename)
	json.Unmarshal(file, &ast)
	return ast
}

func readTokenToMarketFromFile(filename string) map[string]string {
	ast := make(map[string]string)
	file, _ := os.ReadFile(filename)
	json.Unmarshal(file, &ast)
	return ast
}

func readTokenToTicksizeFromFile(filename string) map[string]float64 {
	ast := make(map[string]float64)
	file, _ := os.ReadFile(filename)
	json.Unmarshal(file, &ast)
	return ast
}

func readTokenToIndexFromFile(filename string) map[string]int {
	ast := make(map[string]int)
	file, _ := os.ReadFile(filename)
	json.Unmarshal(file, &ast)
	return ast
}

func readNumNegRiskMarketsFromFile(filename string) map[string]int {
	ast := make(map[string]int)
	file, _ := os.ReadFile(filename)
	json.Unmarshal(file, &ast)
	return ast
}

func readNegRiskMarketsFromFile(filename string) map[string][]NegRiskMarketInfo {
	ast := make(map[string][]NegRiskMarketInfo)
	file, _ := os.ReadFile(filename)
	json.Unmarshal(file, &ast)
	return ast
}

func SetupDicts() {
	AssetsToWatch = readJsonAssetsFromFile("assets.json")
	TokenToMarketMap = readTokenToMarketFromFile("token_to_market.json")
	TokenToTicksize = readTokenToTicksizeFromFile("token_to_ticksize.json")
	TokenToIndex = readTokenToIndexFromFile("token_to_index.json")
	NumNegRiskMarketsMap = readNumNegRiskMarketsFromFile("num_neg_risk_markets.json")
	NegRiskMarketMap = readNegRiskMarketsFromFile("neg_risk_markets.json")
}

func StartSubscription() {
	SetupDicts()
	fmt.Println("%v\n", NegRiskMarketMap)
	InitOrderBooks()
	fmt.Printf("Assets to watch: %v\n", len(AssetsToWatch))
	println("Subscribing to sockets")
	for i := 0; i < len(AssetsToWatch); i += 400 {
		WG.Add(1)
		go ConnectSocket(AssetsToWatch[i:min(i+400, len(AssetsToWatch))])
	}
	WG.Wait()
}
