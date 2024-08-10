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
