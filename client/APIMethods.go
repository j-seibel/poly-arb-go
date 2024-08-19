package client

import (
	"fmt"
	"io"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
)

var POLYMARKET_SIGNER = Signer{
	PK,
	common.HexToAddress("0xA3D381B8C135cEd27efbbd3f231a0E1B6B931ad0"),
	137,
}

func GetAddress() common.Address {
	return POLYMARKET_SIGNER.account
}

func GetCollateralAddress() common.Address {
	return common.HexToAddress(NEG_RISK_CONTRACT_CONFIG.collateral)
}

func GetConditionalAddress() common.Address {
	return common.HexToAddress(NEG_RISK_CONTRACT_CONFIG.conditional_tokens)
}

func GetExchangeAddress() common.Address {
	return common.HexToAddress(NEG_RISK_CONTRACT_CONFIG.exchange)
}

func GetOk() string {
	resp, err := http.Get(HOST)
	if err != nil {
		return "Request not OK"
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "Request not OK"
	}
	sb := string(body)
	return sb
}

func CreateAPIKey(nonce int64) APICreds {
	endpoint := HOST + CREATE_API_KEY
	headers := CreateLevel1Headers(POLYMARKET_SIGNER, nonce)
	raw_creds, err := PostWithL1Headers(endpoint, headers, nil)
	if err != nil {
		fmt.Println(err)
		return APICreds{}
	}
	fmt.Printf("Raw creds %s\n", string(raw_creds[:]))
	return APICreds{}
}

func DeriveAPIKey(nonce int64) APICreds {
	endpoint := HOST + DERIVE_API_KEY
	headers := CreateLevel1Headers(POLYMARKET_SIGNER, nonce)
	raw_creds, err := GetWithL1Headers(endpoint, headers, nil)
	if err != nil {
		fmt.Println(err)
		return APICreds{}
	}
	fmt.Printf("Raw creds %s\n", string(raw_creds[:]))
	return APICreds{}
}
