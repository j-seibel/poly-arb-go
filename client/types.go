package client

import (
	"github.com/ethereum/go-ethereum/common"
)

type RequestArgs struct {
	method      string
	requestPath string
	body        *string
}

type L2Headers struct {
	POLY_ADDRESS    string
	POLY_SIGNATURE  string
	POLY_TIMESTAMP  string
	POLY_API_KEY    string
	POLY_PASSPHRASE string
}

type L1Headers struct {
	POLY_ADDRESS   string
	POLY_SIGNATURE string
	POLY_TIMESTAMP string
	POLY_NONCE     string
}

type APICreds struct {
	api_key        string
	api_secret     string
	api_passphrase string
}

type Signer struct {
	private_key string
	account     common.Address
	chain_id    int
}

type ContractConfig struct {
	exchange           string
	collateral         string
	conditional_tokens string
}

type CLOBAuth struct {
	address   common.Address
	timestamp string
	nonce     int64
	message   string
}

type Domain struct {
	name     string
	version  string
	chain_id int
}

type BookParams struct {
	token_id string
	side     string
}

type Order struct {
	Price string `json:"price"`
	Size  string `json:"size"`
}

type BookData struct {
	Asks      []Order `json:"asks"`
	Bids      []Order `json:"bids"`
	AssetID   string  `json:"asset_id"`
	EventType string  `json:"event_type"`
	Hash      string  `json:"hash"`
	Market    string  `json:"market"`
	Timestamp string  `json:"timestamp"`
}

type PriceChange struct {
	AssetID   string `json:"asset_id"`
	EventType string `json:"event_type"`
	Hash      string `json:"hash"`
	Market    string `json:"market"`
	Price     string `json:"price"`
	Side      string `json:"side"`
	Size      string `json:"size"`
	Timestamp string `json:"timestamp"`
}

type NegRiskMarketInfo struct {
	Condition_id string
	Yes_token_id string
	No_token_id  string
}

type OrderBook struct {
	asks    map[int64]int64
	bids    map[int64]int64
	min_ask int64
}

type NegRiskOrderBook struct {
	condition_id string
	yes_token_id string
	no_token_id  string
	order_books  map[string]*OrderBook
	num_assets   int64
	total_price  int64
}
