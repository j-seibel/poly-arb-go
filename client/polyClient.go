package client

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

const private_key = "0x3273d17583f924c65451317e658532dc62a9cc52505183eb4b65f121660e8ed1"

var singer = Signer{
	private_key,
	common.HexToAddress("0x7e5f4552091a69125d5dfcb7b8c2659029395bdf"),
	137,
}

func CreateLevel2Headers(signer Signer, creds APICreds, requestArgs RequestArgs) L2Headers {
	timestamp := time.Now().Unix()
	timestampString := fmt.Sprintf("%d", timestamp)
	return L2Headers{
		signer.account.String(),
		buildHMACSignature(creds.api_secret, timestampString, requestArgs.method, requestArgs.requestPath, requestArgs.body),
		timestampString,
		creds.api_key,
		creds.api_passphrase,
	}
}
func buildHMACSignature(secret, timestamp, method, requestPath string, body *string) string {
	// Decode the base64 secret
	base64Secret, err := base64.URLEncoding.DecodeString(secret)
	if err != nil {
		return ""
	}

	// Build the message to be signed
	message := timestamp + method + requestPath
	if body != nil {
		message += *body
	}

	// Create HMAC using SHA256
	h := hmac.New(sha256.New, base64Secret)
	h.Write([]byte(message))

	// Encode the signature in base64
	signature := base64.URLEncoding.EncodeToString(h.Sum(nil))

	return signature
}
