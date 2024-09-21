package client

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"poly/arb/builder"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"poly/arb/model"
)

var (
	private_key, _ = crypto.ToECDSA(common.Hex2Bytes(PK))
)

func ExecuteOrder(price int64, volume int64, tokenId string) {
	// Execute an order
	// fmt.Println("Price", RoundToTickSize(int64((price*(volume/PRICE_MULT))), tokenId), tokenId)
	// fmt.Println("Volume", volume)
	var builder = builder.NewExchangeOrderBuilderImpl(big.NewInt(137), nil)
	signed_order, err := builder.BuildSignedOrder(private_key, &model.OrderData{
		Maker:         "0x6cd02aAfEEb049150014D3D9356613897Ce54e6C",
		Signer:        "0xA3D381B8C135cEd27efbbd3f231a0E1B6B931ad0",
		Taker:         "0x0000000000000000000000000000000000000000",
		TokenId:       tokenId,
		MakerAmount:   fmt.Sprintf("%d", RoundToTickSize(int64((price*(volume/PRICE_MULT))), tokenId)),
		TakerAmount:   fmt.Sprintf("%d", RoundToTickSize(volume, tokenId)),
		Side:          model.BUY,
		FeeRateBps:    "0",
		Nonce:         "0",
		Expiration:    "0",
		SignatureType: model.POLY_GNOSIS_SAFE,
	},
		model.NegRiskCTFExchange)

	if err != nil {
		fmt.Println("Error building signed order", err)
	}
	body := GetOrderBody(signed_order)
	// print(RoundToTickSize(int64((price * (volume / PRICE_MULT))), tokenId))
	PostWithL2Headers(HOST+POST_ORDER, CreateLevel2Headers(RequestArgs{"POST", POST_ORDER, &body}), []byte(body))
}

func GetOrderBody(signed_order *model.SignedOrder) string {

	order_body := fmt.Sprintf(`{"order":{"salt": %d,"tokenId": "%s","makerAmount":"%s","takerAmount": "%s","side":"%s","expiration":"%s","nonce":"%s","feeRateBps":"%s","signatureType":%d,"maker":"%s","taker":"%s","signer":"%s","signature":"0x%s"}, "owner": "%s", "orderType": "%s"}`,
		signed_order.Salt, signed_order.TokenId, signed_order.MakerAmount, signed_order.TakerAmount, "BUY", "0", "0", signed_order.FeeRateBps, signed_order.SignatureType, signed_order.Maker, signed_order.Taker, signed_order.Signer, hex.EncodeToString(signed_order.Signature), API_CREDS.api_key, "GTC")

	return order_body
}
