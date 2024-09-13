package client

import (
	"fmt"
	"io"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/goccy/go-json"
)

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

func GetMidpoint(token_id string) string {
	endpoint := HOST + MID_POINT + "?token_id=" + token_id
	raw_creds, err := GetWithL0Headers(endpoint, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Midpoint %s\n", string(raw_creds[:]))
	return string(raw_creds[:])
}

func GetMidpoints(bookParams []BookParams) string {
	endpoint := HOST + MID_POINTS

	requests := []map[string]string{}

	for _, bp := range bookParams {
		requests = append(requests, map[string]string{
			"token_id": bp.token_id,
			"side":     "BUY",
		})
	}
	body, err := json.Marshal(requests)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	raw_creds, err := PostWithL0Headers(endpoint, body)
	if err != nil {
		fmt.Println(err)
	}
	return string(raw_creds[:])

}

func CreateAPIKey(nonce int64) APICreds {
	endpoint := HOST + CREATE_API_KEY
	headers := CreateLevel1Headers(nonce)
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
	headers := CreateLevel1Headers(nonce)
	raw_creds, err := GetWithL1Headers(endpoint, headers, nil)
	if err != nil {
		fmt.Println(err)
		return APICreds{}
	}
	fmt.Printf("Raw creds %s\n", string(raw_creds[:]))
	return APICreds{}
}

func GetNotifications(nonce int64) {
	endpoint := HOST + GET_NOTIFICATIONS + "?signature_type=2"
	requestArgs := RequestArgs{
		"GET",
		GET_NOTIFICATIONS,
		nil,
	}
	headers := CreateLevel2Headers(requestArgs)
	raw_creds, err := GetWithL2Headers(endpoint, headers, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Notifications %s\n", string(raw_creds[:]))
}

func GetMarkets(next_cursor string) string {
	endpoint := HOST + GET_MARKETS + "?next_cursor=" + next_cursor
	requestArgs := RequestArgs{
		"GET",
		GET_MARKETS,
		nil,
	}
	headers := CreateLevel2Headers(requestArgs)
	raw_creds, err := GetWithL2Headers(endpoint, headers, nil)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Printf("Markets %s\n", string(raw_creds[:]))
	return string(raw_creds[:])
}
