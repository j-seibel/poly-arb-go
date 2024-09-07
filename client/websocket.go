package client

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

const host = "wss://ws-subscriptions-clob.polymarket.com/ws/market"

func ConnectSocket(asset_ids []string) {
	c, _, err := websocket.DefaultDialer.Dial(host, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	assets_ids, err := json.Marshal(asset_ids)
	subscriptionMessage := `{"type":"market","assets_ids":` + string(assets_ids) + `}`
	// send subscription message
	err = c.WriteMessage(websocket.TextMessage, []byte(subscriptionMessage))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
	for {

		_, message, err := c.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		var messageMap map[string]interface{}
		err = json.Unmarshal(message, &messageMap)

		if err != nil {
			fmt.Println("Error unmarshalling message:", err)
			continue // Skip to the next message
		}

		if messageMap["event_type"] == "price_change" {
			var priceChange PriceChange
			json.Unmarshal(message, &priceChange)
			// fmt.Println(priceChange)
			UpdateOrderPrice(priceChange)

		} else if messageMap["event_type"] == "book" {
			var bookData BookData
			json.Unmarshal(message, &bookData)
			// fmt.Println(bookData)
			UpdateOrderBook(bookData)
		}

	}
}
